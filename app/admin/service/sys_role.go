package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/config"
	"gorm.io/gorm/clause"

	"github.com/casbin/casbin/v2"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
	"go-admin/common/mycasbin"
)

type SysRole struct {
	service.Service
}

// GetPage 获取SysRole列表
func (e *SysRole) GetPage(c *dto.SysRoleGetPageReq, list *[]models.SysRole, count *int64) error {
	var err error
	var data models.SysRole

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysRole对象
func (e *SysRole) Insert(c *dto.SysRoleInsertReq, cb *casbin.SyncedEnforcer) error {
	var err error
	var data models.SysRole
	c.Generate(&data)
	tx := e.Orm
	if config.DatabaseConfig.Driver != "sqlite3" {
		tx = e.Orm.Begin()
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
	}
	var count int64
	err = tx.Model(&data).Where("role_key = ?", c.RoleKey).Count(&count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	if count > 0 {
		err = errors.New("roleKey已存在，需更换在提交！")
		e.Log.Errorf("db error:%s", err)
		return err
	}

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 获取角色关联的菜单并设置到Casbin策略
	// 首先根据传入的菜单ID创建角色菜单关系
	if len(c.MenuIds) > 0 {
		// 创建角色菜单关系
		roleMenus := make([]models.SysRoleMenu, 0)
		for _, menuId := range c.MenuIds {
			rm := models.SysRoleMenu{
				RoleId: data.RoleId,
				MenuId: menuId,
			}
			roleMenus = append(roleMenus, rm)
		}

		err = tx.Create(&roleMenus).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 查询角色菜单关系
	var roleMenus []models.SysRoleMenu
	err = tx.Where("role_id = ?", data.RoleId).Find(&roleMenus).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 从菜单获取权限code
	menuIds := make([]int, 0, len(roleMenus))
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}

	var menus []models.SysMenu
	if len(menuIds) > 0 {
		err = tx.Where("menu_id in ?", menuIds).Find(&menus).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 从菜单中的权限code获取权限ID
	permCodes := make([]string, 0, len(menus))
	for _, menu := range menus {
		if menu.PermissionCode != "" {
			permCodes = append(permCodes, menu.PermissionCode)
		}
	}

	var perms []models.SysPermission
	if len(permCodes) > 0 {
		err = tx.Where("code in ?", permCodes).Find(&perms).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 获取权限ID列表
	permIds := make([]int, 0, len(perms))
	for _, perm := range perms {
		permIds = append(permIds, perm.Id)
	}

	// 创建角色权限关系
	rolePerms := make([]models.SysRolePermission, 0)
	for _, permId := range permIds {
		rp := models.SysRolePermission{
			RoleId:       data.RoleId,
			PermissionId: permId,
		}
		rolePerms = append(rolePerms, rp)
	}

	if len(rolePerms) > 0 {
		err = tx.Create(&rolePerms).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 从权限API关系表获取API权限
	var permApis []models.SysPermissionApi
	if len(permIds) > 0 {
		err = tx.Where("permission_id in ?", permIds).Find(&permApis).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 构建Casbin策略
	policies := [][]string{}

	// 添加权限API策略
	// 获取所有相关的API ID
	apiIds := make([]int, 0, len(permApis))
	for _, permApi := range permApis {
		apiIds = append(apiIds, permApi.ApiId)
	}

	// 一次性获取所有API信息
	var apis []models.SysApi
	if len(apiIds) > 0 {
		err = tx.Where("id in ?", apiIds).Find(&apis).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 将API信息放入map中便于快速查找
	apiMap := make(map[int]models.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.Id] = api
	}

	// 构建策略
	for _, permApi := range permApis {
		api, exists := apiMap[permApi.ApiId]
		if !exists {
			continue // 如果API不存在，跳过
		}
		policy := []string{data.RoleKey, api.Path, api.Action}
		policies = append(policies, policy)
	}

	// 写入Casbin策略
	if len(policies) > 0 {
		_, err = cb.AddNamedPolicies("p", policies)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update 修改SysRole对象
func (e *SysRole) Update(c *dto.SysRoleUpdateReq, cb *casbin.SyncedEnforcer) error {
	var err error
	tx := e.Orm
	if config.DatabaseConfig.Driver != "sqlite3" {
		tx = e.Orm.Begin()
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
	}

	// 检查是否尝试修改SuperAdmin角色的RoleKey
	var oldRole models.SysRole
	err = tx.Where("role_id = ?", c.RoleId).First(&oldRole).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 如果原角色是admin，禁止修改RoleKey
	if oldRole.RoleKey == mycasbin.SuperAdmin && c.RoleKey != mycasbin.SuperAdmin {
		err = errors.New("SuperAdmin角色的RoleKey不能被修改")
		e.Log.Errorf("Cannot modify SuperAdmin RoleKey")
		return err
	}

	var model = models.SysRole{}
	c.Generate(&model)
	// 更新关联的数据，使用 FullSaveAssociations 模式
	db := tx.Debug().Save(&model)

	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	// 删除旧的角色菜单关系
	err = tx.Where("role_id = ?", model.RoleId).Delete(&models.SysRoleMenu{}).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 重新创建角色菜单关系
	if len(c.MenuIds) > 0 {
		roleMenus := make([]models.SysRoleMenu, 0)
		for _, menuId := range c.MenuIds {
			rm := models.SysRoleMenu{
				RoleId: model.RoleId,
				MenuId: menuId,
			}
			roleMenus = append(roleMenus, rm)
		}

		err = tx.Create(&roleMenus).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 删除旧的角色权限关系
	err = tx.Where("role_id = ?", model.RoleId).Delete(&models.SysRolePermission{}).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 重新创建角色权限关系
	// 查询当前角色的菜单
	var roleMenus []models.SysRoleMenu
	err = tx.Where("role_id = ?", model.RoleId).Find(&roleMenus).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 从菜单获取权限code
	menuIds := make([]int, 0, len(roleMenus))
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}

	var menus []models.SysMenu
	if len(menuIds) > 0 {
		err = tx.Where("menu_id in ?", menuIds).Find(&menus).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 从菜单中的权限code获取权限ID
	permCodes := make([]string, 0, len(menus))
	for _, menu := range menus {
		if menu.PermissionCode != "" {
			permCodes = append(permCodes, menu.PermissionCode)
		}
	}

	var perms []models.SysPermission
	if len(permCodes) > 0 {
		err = tx.Where("code in ?", permCodes).Find(&perms).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 获取权限ID列表
	permIds := make([]int, 0, len(perms))
	for _, perm := range perms {
		permIds = append(permIds, perm.Id)
	}

	// 创建角色权限关系
	rolePerms := make([]models.SysRolePermission, 0)
	for _, permId := range permIds {
		rp := models.SysRolePermission{
			RoleId:       model.RoleId,
			PermissionId: permId,
		}
		rolePerms = append(rolePerms, rp)
	}

	if len(rolePerms) > 0 {
		err = tx.Create(&rolePerms).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 删除旧的Casbin策略
	_, err = cb.RemoveFilteredPolicy(0, model.RoleKey)
	if err != nil {
		return err
	}

	// 从权限API关系表获取API权限
	var permApis []models.SysPermissionApi
	if len(permIds) > 0 {
		err = tx.Where("permission_id in ?", permIds).Find(&permApis).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 获取所有相关的API ID
	apiIds := make([]int, 0, len(permApis))
	for _, permApi := range permApis {
		apiIds = append(apiIds, permApi.ApiId)
	}

	// 一次性获取所有API信息
	var apis []models.SysApi
	if len(apiIds) > 0 {
		err = tx.Where("id in ?", apiIds).Find(&apis).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			return err
		}
	}

	// 将API信息放入map中便于快速查找
	apiMap := make(map[int]models.SysApi, len(apis))
	for _, api := range apis {
		apiMap[api.Id] = api
	}

	// 构建新的Casbin策略
	policies := [][]string{}
	for _, permApi := range permApis {
		api, exists := apiMap[permApi.ApiId]
		if !exists {
			continue // 如果API不存在，跳过
		}
		policy := []string{model.RoleKey, api.Path, api.Action}
		policies = append(policies, policy)
	}

	// 写入新的Casbin策略
	if len(policies) > 0 {
		_, err = cb.AddNamedPolicies("p", policies)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove 删除SysRole
func (e *SysRole) Remove(c *dto.SysRoleDeleteReq, cb *casbin.SyncedEnforcer) error {
	var err error
	tx := e.Orm
	if config.DatabaseConfig.Driver != "sqlite3" {
		tx = e.Orm.Begin()
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
	}

	// 首先获取角色信息，以便获取RoleKey用于删除Casbin策略
	var role models.SysRole
	err = tx.Where("role_id = ?", c.GetId()).First(&role).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 检查是否为SuperAdmin角色，禁止删除
	if role.RoleKey == mycasbin.SuperAdmin {
		err = errors.New("SuperAdmin角色不能被删除")
		e.Log.Errorf("Cannot delete SuperAdmin role")
		return err
	}

	// 删除角色菜单关系
	err = tx.Where("role_id = ?", c.GetId()).Delete(&models.SysRoleMenu{}).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 删除角色权限关系
	err = tx.Where("role_id = ?", c.GetId()).Delete(&models.SysRolePermission{}).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	// 删除角色本身
	var model = models.SysRole{}
	db := tx.Select(clause.Associations).Delete(&model)

	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	// 清除 sys_casbin_rule 权限表里 当前角色的所有记录
	_, _ = cb.RemoveFilteredPolicy(0, role.RoleKey)

	return nil
}

func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) *SysRole {
	var err error
	tx := e.Orm
	if config.DatabaseConfig.Driver != "sqlite3" {
		tx = e.Orm.Begin()
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
	}
	var model = models.SysRole{}
	c.Generate(&model)
	// 更新关联的数据，使用 FullSaveAssociations 模式
	db := tx.Model(&model).Debug().Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}

	// 删除旧的角色部门关系
	err = tx.Where("role_id = ?", c.RoleId).Delete(&models.SysRoleDept{}).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}

	// 重新创建角色部门关系
	if len(c.DeptIds) > 0 {
		roleDepts := make([]models.SysRoleDept, 0)
		for _, deptId := range c.DeptIds {
			rd := models.SysRoleDept{
				RoleId: c.RoleId,
				DeptId: deptId,
			}
			roleDepts = append(roleDepts, rd)
		}

		err = tx.Create(&roleDepts).Error
		if err != nil {
			e.Log.Errorf("db error:%s", err)
			_ = e.AddError(err)
			return e
		}
	}

	return e
}

// UpdateStatus 修改SysRole对象status
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) error {
	var err error
	tx := e.Orm
	if config.DatabaseConfig.Driver != "sqlite3" {
		tx = e.Orm.Begin()
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
	}
	var model = models.SysRole{}
	tx.First(&model, c.GetId())

	// 检查是否为SuperAdmin角色，禁止禁用
	if model.RoleKey == mycasbin.SuperAdmin && c.Status == "2" {
		err = errors.New("SuperAdmin角色不能被禁用")
		e.Log.Errorf("Cannot disable SuperAdmin role")
		return err
	}

	c.Generate(&model)
	// 更新关联的数据，使用 FullSaveAssociations 模式
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// GetWithName 获取SysRole对象
func (e *SysRole) GetWithName(d *dto.SysRoleByName, model *models.SysRole) *SysRole {
	var err error
	db := e.Orm.Where("role_name = ?", d.RoleName).First(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	return e
}

// GetPremissonByRoleId 根据角色ID获取权限代码列表
func (e *SysRole) GetPremissonByRoleId(roleId int, host string) ([]string, error) {
	// 1. 查询角色的权限关联
	var rolePerms []models.SysRolePermission
	if err := e.Orm.Where("role_id = ?", roleId).Find(&rolePerms).Error; err != nil {
		return nil, err
	}

	if len(rolePerms) == 0 {
		return []string{}, nil
	}

	// 2. 获取权限ID列表
	permIds := make([]int, 0, len(rolePerms))
	for _, rp := range rolePerms {
		permIds = append(permIds, rp.PermissionId)
	}

	// 3. 查询权限详情,只获取已启用的权限
	var permissions []models.SysPermission
	if err := e.Orm.Where("id in ? AND status = ?", permIds, 1).Find(&permissions).Error; err != nil {
		return nil, err
	}

	// 4. 收集权限代码
	permissionCodes := make([]string, 0, len(permissions))
	for _, perm := range permissions {
		if perm.Code != "" {
			permissionCodes = append(permissionCodes, perm.Code)
		}
	}

	return permissionCodes, nil
}

// GetById 获取SysRole对象
func (e *SysRole) Get(GetRoleReq *dto.SysRoleGetReq, object *models.SysRole) error {
	if err := e.Orm.Model(&models.SysRole{}).Where("role_id = ?", GetRoleReq.Id).First(object).Error; err != nil {
		return err
	}

	// 追加加载角色关联的权限列表
	var rolePerms []models.SysRolePermission
	if err := e.Orm.Where("role_id = ?", object.RoleId).Find(&rolePerms).Error; err != nil {
		return err
	}
	if len(rolePerms) == 0 {
		return nil
	}

	permIdSet := make(map[int]struct{}, len(rolePerms))
	for _, rp := range rolePerms {
		permIdSet[rp.PermissionId] = struct{}{}
	}
	permIds := make([]int, 0, len(permIdSet))
	for id := range permIdSet {
		permIds = append(permIds, id)
	}

	var perms []models.SysPermission
	if err := e.Orm.Where("id in ?", permIds).Find(&perms).Error; err != nil {
		return err
	}
	object.Permissions = perms

	// 追加加载角色关联的菜单列表
	var roleMenus []models.SysRoleMenu
	if err := e.Orm.Where("role_id = ?", object.RoleId).Find(&roleMenus).Error; err != nil {
		return err
	}
	if len(roleMenus) == 0 {
		return nil
	}

	menuIdSet := make(map[int]struct{}, len(roleMenus))
	for _, rm := range roleMenus {
		menuIdSet[rm.MenuId] = struct{}{}
	}
	menuIds := make([]int, 0, len(menuIdSet))
	for id := range menuIdSet {
		menuIds = append(menuIds, id)
	}

	var allMenus []models.SysMenu
	if err := e.Orm.Where("menu_id in ?", menuIds).Find(&allMenus).Error; err != nil {
		return err
	}

	var menusTree []models.SysMenu

	for i := 0; i < len(allMenus); i++ {
		if allMenus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&allMenus, allMenus[i])
		menusTree = append(menusTree, menusInfo)
	}

	object.Menus = menusTree
	return nil
}

// buildMenuTreeForRole 构建角色的菜单树
func buildMenuTreeForRole(menu models.SysMenu, allMenus []models.SysMenu) models.SysMenu {
	min := make([]models.SysMenu, 0)
	for j := 0; j < len(allMenus); j++ {
		if menu.MenuId != allMenus[j].ParentId {
			continue
		}
		mi := models.SysMenu{}
		mi.MenuId = allMenus[j].MenuId
		mi.MenuName = allMenus[j].MenuName
		mi.Icon = allMenus[j].Icon
		mi.Path = allMenus[j].Path
		mi.MenuType = allMenus[j].MenuType
		mi.Perm = allMenus[j].Perm
		mi.ParentId = allMenus[j].ParentId
		mi.KeepAlive = allMenus[j].KeepAlive
		mi.Component = allMenus[j].Component
		mi.SortValue = allMenus[j].SortValue
		mi.IsHide = allMenus[j].IsHide
		mi.CreatedAt = allMenus[j].CreatedAt
		mi.Children = []models.SysMenu{}
		mi.PermissionCode = allMenus[j].PermissionCode
		mi.Permission = allMenus[j].Permission

		if mi.MenuType != "F" {
			ms := buildMenuTreeForRole(mi, allMenus)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}

// GetRoleMenuId 根据roleId获取角色关联的菜单ID列表
func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	var err error
	var role models.SysRole

	// 1. 检查角色是否存在
	err = e.Orm.First(&role, roleId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		e.Log.Errorf("Query role error: %s", err)
		return nil, err
	}

	// 2. 查询角色菜单关系
	var roleMenus []models.SysRoleMenu
	err = e.Orm.Where("role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		e.Log.Errorf("Query role menus error: %s", err)
		return nil, err
	}

	// 3. 提取菜单ID列表
	menuIds := make([]int, 0, len(roleMenus))
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}

	return menuIds, nil
}
