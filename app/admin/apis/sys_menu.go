package apis

import (
	"go-admin/app/admin/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysMenu struct {
	api.Api
}

// GetPage Menu列表数据
// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /lotus/api/v1/menu [get]
// @Security Bearer
func (e SysMenu) GetPage(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list).Error
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "查询成功")
}

// Get 获取菜单详情
// @Summary Menu详情数据
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}
// @Router /lotus/api/v1/menu/get [get]
// @Security Bearer
func (e SysMenu) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object = models.SysMenu{}

	err = s.Get(&req, &object).Error
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert 创建菜单
// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMenuInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /lotus/api/v1/menu [post]
// @Security Bearer
func (e SysMenu) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改菜单
// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMenuUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}
// @Router /lotus/api/v1/menu [put]
// @Security Bearer
func (e SysMenu) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req).Error
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param data body dto.SysMenuDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /lotus/api/v1/menu [delete]
// @Security Bearer
func (e SysMenu) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(control, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Remove(control).Error
	if err != nil {
		e.Logger.Errorf("RemoveSysMenu error, %s", err)
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(control.GetId(), "删除成功")
}

// GetMenuRole 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Summary 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /lotus/api/v1/menurole [get]
// @Security Bearer
func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.SetMenuRole(user.GetUserId(c))

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(result, "")
}

// roleMenuTreeselect 根据角色ID查询菜单下拉树结构
// @Summary 获取角色的菜单树形结构
// @Description 根据角色ID查询所有菜单，然后根据该角色的已经关联的菜单标记对它们不可选控
// @Tags 菜单
// @Accept  application/json
// @Produce application/json
// @Param roleId query int false "角色ID"
// @Success 200 {object} response.Response "{\"code\": 200, \"data\": {...}}"
// @Router /lotus/api/v1/roleMenuTreeselect [get]
// @Security Bearer
func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 1. 获取所有菜单树形结构
	result, err := m.SetLabel()
	if err != nil {
		e.Logger.Errorf("查询菜单树失败: %s", err)
		e.Error(500, err, "查询菜单失败")
		return
	}

	// 2. 如果指定了角色ID，获取该角色已关联的菜单ID列表
	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = r.GetRoleMenuId(req.RoleId)
		if err != nil {
			e.Logger.Errorf("获取角色菜单关系失败: roleId=%d, error=%s", req.RoleId, err)
			e.Error(500, err, "获取角色菜单关系失败")
			return
		}
	}

	// 3. 返回结果
	e.OK(gin.H{
		"menus":       result,  // 所有菜单树形结构
		"checkedKeys": menuIds, // 当前角色已勾选的菜单ID列表
	}, "获取成功")
}
