package service

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	cDto "go-admin/common/dto"
	cModels "go-admin/common/models"
	"go-admin/common/mycasbin"
)

type SysMenu struct {
	service.Service
}

// GetPage 获取SysMenu列表
func (e *SysMenu) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) *SysMenu {
	var menu = make([]models.SysMenu, 0)
	err := e.getPage(c, &menu).Error
	if err != nil {
		_ = e.AddError(err)
		return e
	}

	// 加载权限信息
	e.loadPermissionsForMenus(&menu)

	for i := 0; i < len(menu); i++ {
		if menu[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menu, menu[i])
		*menus = append(*menus, menusInfo)
	}
	return e
}

// getPage 菜单分页列表
func (e *SysMenu) getPage(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) *SysMenu {
	var err error
	var data models.SysMenu

	err = e.Orm.Model(&data).
		Scopes(
			cDto.OrderDest(models.SysMenuSortValue, false),
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("getSysMenuPage error:%s", err)
		_ = e.AddError(err)
		return e
	}

	// 批量加载关联的权限信息
	e.loadPermissionsForMenus(list)

	return e
}

// Get 获取SysMenu对象
func (e *SysMenu) Get(d *dto.SysMenuGetReq, model *models.SysMenu) *SysMenu {
	var err error
	var data models.SysMenu

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("GetSysMenu error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}

	// 加载关联的权限信息
	if model.PermissionCode != "" {
		var permission models.SysPermission
		err = e.Orm.Where("code = ?", model.PermissionCode).First(&permission).Error
		if err != nil {
			e.Log.Warnf("加载菜单权限失败: %v", err)
		} else {
			model.Permission = permission
		}
	}

	return e
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(c *dto.SysMenuInsertReq) *SysMenu {
	var err error
	var data models.SysMenu
	c.Generate(&data)

	// 检查PermissionCode是否为空
	if data.PermissionCode != "" {
		// 检查PermissionCode是否已存在
		var count int64
		err = e.Orm.Model(&models.SysMenu{}).Where("permission_code = ?", data.PermissionCode).Count(&count).Error
		if err != nil {
			e.Log.Errorf("检查PermissionCode唯一性失败: %s", err)
			_ = e.AddError(err)
			return e
		}
		if count > 0 {
			err = errors.New("PermissionCode已存在，请使用唯一的PermissionCode")
			e.Log.Errorf("PermissionCode重复: %s", data.PermissionCode)
			_ = e.AddError(err)
			return e
		}
	} else {
		err = errors.New("PermissionCode不能为空")
		e.Log.Errorf("PermissionCode 参数为空")
		_ = e.AddError(err)
		return e
	}

	tx := e.Orm.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // 重新抛出panic
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
	}
	c.MenuId = data.MenuId
	err = e.initPaths(tx, &data)
	if err != nil {
		tx.Rollback()
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
	}
	tx.Commit()
	return e
}

func (e *SysMenu) initPaths(tx *gorm.DB, menu *models.SysMenu) error {
	var err error
	var data models.SysMenu
	parentMenu := new(models.SysMenu)
	if menu.ParentId != 0 {
		err = tx.Model(&data).First(parentMenu, menu.ParentId).Error
		if err != nil {
			return err
		}
		if parentMenu.MenuPath == "" {
			// 父级MenuPath异常，重建整条链路路径
			return e.RebuildAllMenuPath(tx)
		} else {
			menu.MenuPath = parentMenu.MenuPath + "/" + pkg.IntToString(menu.MenuId)
		}
	} else {
		menu.MenuPath = "/0/" + pkg.IntToString(menu.MenuId)
	}
	err = tx.Model(&data).Where("menu_id = ?", menu.MenuId).Update(models.SysMenuMenuPath, menu.MenuPath).Error
	return err
}

// RebuildAllMenuPath 全量重建所有菜单的 menu_path
func (e *SysMenu) RebuildAllMenuPath(tx *gorm.DB) error {
	// 1. 查询所有菜单
	var menus []models.SysMenu
	if err := tx.Find(&menus).Error; err != nil {
		return err
	}

	if len(menus) == 0 {
		return nil
	}

	// 2. 构建 parent -> children 映射
	childrenMap := make(map[int64][]*models.SysMenu)
	menuMap := make(map[int64]*models.SysMenu)

	for i := range menus {
		m := &menus[i]
		menuMap[int64(m.MenuId)] = m
		childrenMap[int64(m.ParentId)] = append(childrenMap[int64(m.ParentId)], m)
	}

	// 3. DFS 构建路径
	pathMap := make(map[int64]string)
	visited := make(map[int64]bool)

	var dfs func(m *models.SysMenu, parentPath string) error
	dfs = func(m *models.SysMenu, parentPath string) error {
		// 环检测（防止 parent 指向自己 / 死循环）
		if visited[int64(m.MenuId)] {
			return fmt.Errorf("menu cycle detected, menu_id=%d", m.MenuId)
		}
		visited[int64(m.MenuId)] = true

		currentPath := parentPath + "/" + pkg.IntToString(m.MenuId)
		pathMap[int64(m.MenuId)] = currentPath

		for _, child := range childrenMap[int64(m.MenuId)] {
			if err := dfs(child, currentPath); err != nil {
				return err
			}
		}

		return nil
	}

	// 4. 从虚拟根开始
	for _, root := range childrenMap[0] {
		rootPath := "/0/" + pkg.IntToString(root.MenuId)
		pathMap[int64(root.MenuId)] = rootPath
		visited[int64(root.MenuId)] = true

		for _, child := range childrenMap[int64(root.MenuId)] {
			if err := dfs(child, rootPath); err != nil {
				return err
			}
		}
	}

	// 5. 更新数据库（逐条，安全优先）
	for id, path := range pathMap {
		if err := tx.Model(&models.SysMenu{}).
			Where("menu_id = ?", id).
			Update(models.SysMenuMenuPath, path).Error; err != nil {
			return err
		}
	}

	return nil
}

// Update 修改SysMenu对象
func (e *SysMenu) Update(c *dto.SysMenuUpdateReq) *SysMenu {
	var err error
	tx := e.Orm.Debug().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // 重新抛出panic
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysMenu{}
	// 先查询原记录以获取oldPath
	if err = tx.First(&model, c.MenuId).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	oldPath := model.MenuPath

	// 使用map进行更新，支持零值
	updateData := map[string]interface{}{
		"menu_name":     c.MenuName,
		"title":         c.Title,
		"menu_type":     c.MenuType,
		"menu_path":     c.MenuPath,
		"path":          c.Path,
		"perm":          c.Perm,
		"component":     c.Component,
		"icon":          c.Icon,
		"sort_value":    c.SortValue,
		"is_external":   c.IsExternal,
		"external_link": c.ExternalLink,
		"text_badge":    c.TextBadge,
		"active_path":   c.ActivePath,
		"status":        c.Status,
		"keep_alive":    c.KeepAlive,
		"is_hide":       c.IsHide,
		"is_iframe":     c.IsIframe,
		"show_badge":    c.ShowBadge,
		"fixed_tab":     c.FixedTab,
		"is_hide_tab":   c.IsHideTab,
		"is_full_page":  c.IsFullPage,
		"parent_id":     c.ParentId,
	}
	db := tx.Model(&models.SysMenu{}).Where("menu_id = ?", c.MenuId).Updates(updateData)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}
	// 更新子路径：如果父路径程改变，最通配罦衔路径也要修改
	var menuList []models.SysMenu
	tx.Where("menu_path like ?", oldPath+"%").Find(&menuList)
	for _, v := range menuList {
		v.MenuPath = strings.Replace(v.MenuPath, oldPath, model.MenuPath, 1)
		tx.Model(&v).Where("menu_id = ?", v.MenuId).Update(models.SysMenuMenuPath, v.MenuPath)
	}
	return e
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(d *dto.SysMenuDeleteReq) *SysMenu {
	var err error
	var data models.SysMenu

	// 使用事务确保数据一致性
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // 重新抛出panic
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 先删除角色与菜单的关联关系
	err = tx.Where("menu_id in ?", d.Ids).Delete(&models.SysRoleMenu{}).Error
	if err != nil {
		e.Log.Errorf("Delete role-menu relations error: %s", err)
		_ = e.AddError(err)
		return e
	}

	// 删除菜单
	db := tx.Model(&data).Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		e.Log.Errorf("Delete error: %s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		_ = e.AddError(err)
		return e
	}
	return e
}

// GetList 获取菜单数据
func (e *SysMenu) GetList(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) error {
	var err error
	var data models.SysMenu

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 加载权限信息
	e.loadPermissionsForMenus(list)

	return nil
}

// SetLabel 修改角色中 设置菜单基础数据
func (e *SysMenu) SetLabel() (m []dto.MenuLabel, err error) {
	var list []models.SysMenu
	err = e.GetList(&dto.SysMenuGetPageReq{}, &list)

	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].MenuId
		e.MenuName = list[i].MenuName
		e.Title = list[i].Title
		deptsInfo := menuLabelCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	list := *menuList

	min := make([]models.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Perm = list[j].Perm
		mi.ParentId = list[j].ParentId
		mi.KeepAlive = list[j].KeepAlive
		mi.Component = list[j].Component
		mi.SortValue = list[j].SortValue
		mi.IsHide = list[j].IsHide
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysMenu{}
		mi.PermissionCode = list[j].PermissionCode
		mi.Permission = list[j].Permission

		if mi.MenuType != cModels.Button {
			ms := menuCall(menuList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

func menuDistinct(menuList []models.SysMenu) (result []models.SysMenu) {
	distinctMap := make(map[int]struct{}, len(menuList))
	for _, menu := range menuList {
		if _, ok := distinctMap[menu.MenuId]; !ok {
			distinctMap[menu.MenuId] = struct{}{}
			result = append(result, menu)
		}
	}
	return result
}

func recursiveSetMenu(orm *gorm.DB, mIds []int, menus *[]models.SysMenu) error {
	if len(mIds) == 0 || menus == nil {
		return nil
	}
	var subMenus []models.SysMenu
	err := orm.Where(fmt.Sprintf(" menu_type in ('%s', '%s', '%s') and menu_id in ?",
		cModels.Directory, cModels.Menu, cModels.Button), mIds).Order("sort_value").Find(&subMenus).Error
	if err != nil {
		return err
	}

	subIds := make([]int, 0)
	for _, menu := range subMenus {
		if menu.ParentId != 0 {
			subIds = append(subIds, menu.ParentId)
		}
		if menu.MenuType != cModels.Button {
			*menus = append(*menus, menu)
		}
	}
	return recursiveSetMenu(orm, subIds, menus)
}

// SetMenuRole 获取左侧菜单树使用
func (e *SysMenu) SetMenuRole(userId int) (m []models.SysMenu, err error) {
	menus, err := e.getByUserId(userId)
	m = make([]models.SysMenu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}

func (e *SysMenu) getByUserId(userId int) ([]models.SysMenu, error) {
	var err error
	data := make([]models.SysMenu, 0)

	// 1. 查询用户的角色
	var userRoles []models.SysUserRole
	err = e.Orm.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}

	if len(userRoles) == 0 {
		// 用户没有角色，返回空菜单
		return data, nil
	}

	// 2. 获取角色ID列表
	roleIds := make([]int, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}

	// 3. 查询角色信息，检查是否有SuperAdmin
	var roles []models.SysRole
	err = e.Orm.Where("role_id in ?", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 检查是否有admin角色
	isSuperAdmin := false
	for _, role := range roles {
		if role.RoleKey == mycasbin.SuperAdmin {
			isSuperAdmin = true
			break
		}
	}

	// 4. SuperAdmin：直接加载全部菜单
	if isSuperAdmin {
		err = e.Orm.
			Where("menu_type IN ('M','C') AND deleted_at IS NULL").
			Order("sort_value").
			Find(&data).
			Error
		if err != nil {
			return nil, err
		}
	} else {
		// 5. 非 SuperAdmin：通过角色菜单关系查询菜单
		var roleMenus []models.SysRoleMenu
		err = e.Orm.Where("role_id in ?", roleIds).Find(&roleMenus).Error
		if err != nil {
			return nil, err
		}

		if len(roleMenus) == 0 {
			// 角色没有菜单权限
			return data, nil
		}

		// 6. 提取菜单ID
		menuIDs := make([]int, 0, len(roleMenus))
		for _, rm := range roleMenus {
			menuIDs = append(menuIDs, rm.MenuId)
		}

		// 7. 递归补全父级菜单
		if err := recursiveSetMenu(e.Orm, menuIDs, &data); err != nil {
			return nil, err
		}

		// 8. 菜单去重
		data = menuDistinct(data)

		sort.Sort(models.SysMenuSlice(data))
	}

	// 加载权限信息
	e.loadPermissionsForMenus(&data)

	return data, err
}

func (e *SysMenu) loadPermissionsForMenus(menus *[]models.SysMenu) {
	if len(*menus) == 0 {
		return
	}

	// 收集所有权限Code
	permissionCodes := make([]string, 0)
	permissionMap := make(map[string]models.SysPermission)

	for _, menu := range *menus {
		if menu.PermissionCode != "" {
			permissionCodes = append(permissionCodes, menu.PermissionCode)
		}
	}

	if len(permissionCodes) == 0 {
		return
	}

	// 批量查询权限
	var permissions []models.SysPermission
	err := e.Orm.Where("code IN ?", permissionCodes).Find(&permissions).Error
	if err != nil {
		e.Log.Warnf("批量加载菜单权限失败: %v", err)
		return
	}

	// 构建权限映射
	for _, perm := range permissions {
		permissionMap[perm.Code] = perm
	}

	// 关联权限到菜单
	for i := range *menus {
		if (*menus)[i].PermissionCode != "" {
			if perm, exists := permissionMap[(*menus)[i].PermissionCode]; exists {
				(*menus)[i].Permission = perm
			}
		}
	}
}

// ExportMenuPermission 导出菜单和权限数据
func (e *SysMenu) ExportMenuPermission() (*dto.MenuPermissionIO, error) {
	// 获取所有菜单
	var allMenus []models.SysMenu
	var menuReq dto.SysMenuGetPageReq
	err := e.GetList(&menuReq, &allMenus)
	if err != nil {
		e.Log.Errorf("获取菜单列表失败: %v", err)
		return nil, err
	}

	// 构建菜单树
	menuTree := make([]models.SysMenu, 0)
	for i := 0; i < len(allMenus); i++ {
		if allMenus[i].ParentId != 0 {
			continue
		}
		menuInfo := menuCall(&allMenus, allMenus[i])
		menuTree = append(menuTree, menuInfo)
	}

	// 转换为导出格式
	menuIOList := convertMenusToIO(menuTree)

	// 获取所有权限
	var permissions []models.SysPermission
	err = e.Orm.Find(&permissions).Error
	if err != nil {
		e.Log.Errorf("获取权限列表失败: %v", err)
		return nil, err
	}

	// 转换为导出格式
	permissionIOList := make([]dto.PermissionIO, 0)
	for _, perm := range permissions {
		// 获取权限关联的API
		var permApis []models.SysPermissionApi
		err = e.Orm.Where("permission_id = ?", perm.Id).Find(&permApis).Error
		if err != nil {
			e.Log.Warnf("获取权限API关联失败: %v", err)
			continue
		}

		// 获取API详情
		apiIds := make([]int, 0)
		for _, pa := range permApis {
			apiIds = append(apiIds, pa.ApiId)
		}

		var apis []models.SysApi
		if len(apiIds) > 0 {
			err = e.Orm.Where("id in ?", apiIds).Find(&apis).Error
			if err != nil {
				e.Log.Warnf("获取API详情失败: %v", err)
			}
		}

		apiIOList := make([]dto.ApiIO, 0)
		for _, api := range apis {
			apiIOList = append(apiIOList, dto.ApiIO{
				Method: api.Action,
				URL:    api.Path,
			})
		}

		permissionIOList = append(permissionIOList, dto.PermissionIO{
			Code: perm.Code,
			Name: perm.Name,
			Type: perm.Type,
			Apis: apiIOList,
		})
	}

	result := &dto.MenuPermissionIO{
		Menus:       menuIOList,
		Permissions: permissionIOList,
	}

	return result, nil
}

// convertMenusToIO 递归转换菜单模型为导出格式
func convertMenusToIO(menus []models.SysMenu) []dto.MenuIO {
	result := make([]dto.MenuIO, 0)
	for _, menu := range menus {
		menuIO := dto.MenuIO{
			MenuType:       menu.MenuType,
			Path:           menu.Path,
			Component:      menu.Component,
			Perm:           menu.Perm,
			MenuName:       menu.MenuName,
			Title:          menu.Title,
			PermissionCode: menu.PermissionCode,
			Icon:           menu.Icon,
			SortValue:      menu.SortValue,
			IsExternal:     menu.IsExternal,
			ExternalLink:   menu.ExternalLink,
			TextBadge:      menu.TextBadge,
			ActivePath:     menu.ActivePath,
			Status:         menu.Status,
			KeepAlive:      menu.KeepAlive,
			IsHide:         menu.IsHide,
			IsIframe:       menu.IsIframe,
			ShowBadge:      menu.ShowBadge,
			FixedTab:       menu.FixedTab,
			IsHideTab:      menu.IsHideTab,
			IsFullPage:     menu.IsFullPage,
			Children:       convertMenusToIO(menu.Children),
		}
		result = append(result, menuIO)
	}
	return result
}

// ImportMenuPermission 导入菜单和权限数据
func (e *SysMenu) ImportMenuPermission(data *dto.MenuPermissionIO) error {
	// 开始事务
	tx := e.Orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// 导入权限数据
	err := e.importPermissions(tx, data.Permissions)
	if err != nil {
		e.Log.Errorf("导入权限数据失败: %v", err)
		tx.Rollback()
		return err
	}

	// 导入菜单数据
	err = e.importMenus(tx, data.Menus)
	if err != nil {
		e.Log.Errorf("导入菜单数据失败: %v", err)
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		e.Log.Errorf("提交事务失败: %v", err)
		return err
	}

	return nil
}

// importPermissions 导入权限数据
func (e *SysMenu) importPermissions(tx *gorm.DB, permissions []dto.PermissionIO) error {
	for _, perm := range permissions {
		// 检查权限是否已存在
		var existingPermission models.SysPermission
		if err := tx.Where("code = ?", perm.Code).First(&existingPermission).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 权限不存在，创建新权限
				permissionModel := models.SysPermission{
					Code:     perm.Code,
					Name:     perm.Name,
					Type:     perm.Type,
					ParentId: 0, // 暂时设置为0，如果需要层级关系可以扩展
					Sort:     0, // 暂时设置为0
					Status:   1, // 默认启用
				}

				if err := tx.Create(&permissionModel).Error; err != nil {
					e.Log.Errorf("创建权限失败: %v, Code: %s", err, perm.Code)
					return err
				}

				// 创建API关联
				for _, api := range perm.Apis {
					// 先查找API是否存在
					var apiModel models.SysApi
					if err := tx.Where("path = ? AND action = ?", api.URL, api.Method).First(&apiModel).Error; err != nil {
						if errors.Is(err, gorm.ErrRecordNotFound) {
							// 如果API不存在，跳过关联
							e.Log.Warnf("API不存在，跳过关联: %s %s", api.Method, api.URL)
							continue
						} else {
							// 其他错误，返回
							e.Log.Errorf("查询API失败: %v", err)
							return err
						}
					}

					// 创建权限API关联
					permApi := models.SysPermissionApi{
						PermissionId: permissionModel.Id,
						ApiId:        apiModel.Id,
					}
					if err := tx.Create(&permApi).Error; err != nil {
						e.Log.Errorf("创建权限API关联失败: %v", err)
						return err
					}
				}
			} else {
				// 其他错误，返回
				e.Log.Errorf("查询权限失败: %v", err)
				return err
			}
		} else {
			// 权限已存在，更新权限信息
			updateData := map[string]interface{}{
				"name":   perm.Name,
				"type":   perm.Type,
				"status": 1, // 默认启用
			}
			if err := tx.Model(&existingPermission).Updates(updateData).Error; err != nil {
				e.Log.Errorf("更新权限失败: %v, Code: %s", err, perm.Code)
				return err
			}

			// 删除旧的API关联
			if err := tx.Where("permission_id = ?", existingPermission.Id).Delete(&models.SysPermissionApi{}).Error; err != nil {
				e.Log.Errorf("删除权限API关联失败: %v", err)
				return err
			}

			// 重新创建API关联
			for _, api := range perm.Apis {
				// 先查找API是否存在
				var apiModel models.SysApi
				if err := tx.Where("path = ? AND action = ?", api.URL, api.Method).First(&apiModel).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						// 如果API不存在，跳过关联
						e.Log.Warnf("API不存在，跳过关联: %s %s", api.Method, api.URL)
						continue
					} else {
						// 其他错误，返回
						e.Log.Errorf("查询API失败: %v", err)
						return err
					}
				}

				// 创建权限API关联
				permApi := models.SysPermissionApi{
					PermissionId: existingPermission.Id,
					ApiId:        apiModel.Id,
				}
				if err := tx.Create(&permApi).Error; err != nil {
					e.Log.Errorf("创建权限API关联失败: %v", err)
					return err
				}
			}
		}
	}

	return nil
}

// importMenus 导入菜单数据
func (e *SysMenu) importMenus(tx *gorm.DB, menus []dto.MenuIO) error {
	// 使用map存储已创建的菜单，以便处理父子关系
	menuIdMap := make(map[string]int) // key: 菜单PermissionCode，value: 菜单ID

	// 递归导入菜单及其子菜单，同时设置排序顺序
	for i, menu := range menus {
		err := e.importSingleMenu(tx, menu, 0, &menuIdMap, i+1)
		if err != nil {
			return err
		}
	}

	return nil
}

// importSingleMenu 导入单个菜单
func (e *SysMenu) importSingleMenu(tx *gorm.DB, menu dto.MenuIO, parentId int, menuIdMap *map[string]int, sortOrder int) error {
	// 检查菜单是否已存在（通过permission_code）
	var existingMenu models.SysMenu
	if err := tx.Where("permission_code = ?", menu.PermissionCode).First(&existingMenu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 菜单不存在，创建新菜单
			menuModel := models.SysMenu{
				MenuType:       menu.MenuType,
				Path:           menu.Path,
				Component:      menu.Component,
				Perm:           menu.Perm,
				MenuName:       menu.MenuName,
				Title:          menu.Title,
				PermissionCode: menu.PermissionCode,
				Icon:           menu.Icon,
				SortValue:      menu.SortValue,
				IsExternal:     menu.IsExternal,
				ExternalLink:   menu.ExternalLink,
				TextBadge:      menu.TextBadge,
				ActivePath:     menu.ActivePath,
				Status:         menu.Status,
				KeepAlive:      menu.KeepAlive,
				IsHide:         menu.IsHide,
				IsIframe:       menu.IsIframe,
				ShowBadge:      menu.ShowBadge,
				FixedTab:       menu.FixedTab,
				IsHideTab:      menu.IsHideTab,
				ParentId:       parentId,
			}

			// 检查PermissionCode是否为空
			if menuModel.PermissionCode != "" {
				// 检查PermissionCode是否已存在
				var count int64
				err = tx.Model(&models.SysMenu{}).Where("permission_code = ?", menuModel.PermissionCode).Count(&count).Error
				if err != nil {
					e.Log.Errorf("检查PermissionCode唯一性失败: %s", err)
					return err
				}
				if count > 0 {
					err = errors.New("PermissionCode已存在，请使用唯一的PermissionCode")
					e.Log.Errorf("PermissionCode重复: %s", menuModel.PermissionCode)
					return err
				}
			}

			// 创建菜单
			if err := tx.Create(&menuModel).Error; err != nil {
				e.Log.Errorf("创建菜单失败: %v, Name: %s", err, menu.MenuName)
				return err
			}

			// 存储菜单ID映射
			(*menuIdMap)[menu.PermissionCode] = menuModel.MenuId

			// 递归导入子菜单
			for i, child := range menu.Children {
				err := e.importSingleMenu(tx, child, menuModel.MenuId, menuIdMap, i+1)
				if err != nil {
					return err
				}
			}

		} else {
			// 其他错误，返回
			e.Log.Errorf("查询菜单失败: %v", err)
			return err
		}
	} else {
		// 菜单已存在，更新菜单信息
		updateData := map[string]interface{}{
			"menu_type":       menu.MenuType,
			"path":            menu.Path,
			"component":       menu.Component,
			"perm":            menu.Perm,
			"menu_name":       menu.MenuName,
			"title":           menu.Title,
			"permission_code": menu.PermissionCode,
			"icon":            menu.Icon,
			"sort_value":      menu.SortValue,
			"is_external":     menu.IsExternal,
			"external_link":   menu.ExternalLink,
			"text_badge":      menu.TextBadge,
			"active_path":     menu.ActivePath,
			"status":          menu.Status,
			"keep_alive":      menu.KeepAlive,
			"is_hide":         menu.IsHide,
			"is_iframe":       menu.IsIframe,
			"show_badge":      menu.ShowBadge,
			"fixed_tab":       menu.FixedTab,
			"is_hide_tab":     menu.IsHideTab,
			"parent_id":       parentId,
		}
		if err := tx.Model(&existingMenu).Updates(updateData).Error; err != nil {
			e.Log.Errorf("更新菜单失败: %v, Name: %s", err, menu.MenuName)
			return err
		}

		// 存储菜单ID映射
		(*menuIdMap)[menu.PermissionCode] = existingMenu.MenuId

		// 递归导入子菜单
		for i, child := range menu.Children {
			err := e.importSingleMenu(tx, child, existingMenu.MenuId, menuIdMap, i+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
