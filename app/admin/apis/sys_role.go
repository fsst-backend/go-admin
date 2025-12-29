package apis

import (
	"errors"
	"fmt"
	"go-admin/common/global"
	"go-admin/common/mycasbin"
	"net/http"

	"go-admin/app/admin/models"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
)

type SysRole struct {
	api.Api
}

// GetPage 获取SysRole列表
// @Summary 获取SysRole列表
// @Description 获取SysRole列表
// @Tags 角色管理
// @Param roleName query string false "角色名称"
// @Param status query string false "状态"
// @Param roleKey query string false "角色代码"
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {object} response.Response{message=[]models.SysRole} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role [get]
// @Security Bearer
func (e SysRole) GetPage(c *gin.Context) {
	req := dto.SysRoleGetPageReq{}
	s := service.SysRole{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.Name = ctx.GetUsername(c)
	list := make([]models.SysRole, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "获取角色列表失败,"+err.Error())
		return
	}

	e.PageOK(list, int(count), req.GetOffset(), req.GetLimit(), "获取成功")
}

// Insert 创建SysRole
// @Summary 创建SysRole
// @Description 创建SysRole
// @Tags 角色管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRoleInsertReq true "data"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role [post]
// @Security Bearer
func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
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
	req.CreateBy = user.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}

	// 检查是否尝试创建SuperAdmin角色
	if req.RoleKey == mycasbin.SuperAdmin {
		e.Logger.Error("不能创建SuperAdmin角色")
		e.Error(500, errors.New("不能创建SuperAdmin角色"), "创建失败, 不能创建SuperAdmin角色")
		return
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Insert(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "创建失败,"+err.Error())
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "创建失败,"+err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Get
// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色管理
// @Success 200 {object} response.Response{message=models.SysRole} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role/get [get]
// @Security Bearer
func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(" %s ", err.Error()))
		return
	}

	object := models.SysRole{}
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(object, "查询成功")
}

// Update 修改用户角色
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色管理
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleUpdateReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role [put]
// @Security Bearer
func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	req.SetUpdateBy(user.GetUserId(c))

	err = s.Update(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		return
	}

	_, err = global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "更新失败,"+err.Error())
		return
	}

	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色管理
// @Param data body dto.SysRoleDeleteReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role [delete]
// @Security Bearer
func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("删除角色 %v 失败，\r\n失败信息 %s", req.Ids, err.Error()))
		return
	}

	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Remove(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "")
		return
	}

	e.OK(req.GetId(), fmt.Sprintf("删除角色角色 %v 状态成功！", req.GetId()))
}

// Update2Status 修改用户角色状态
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色管理
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateStatusReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/role-status [put]
// @Security Bearer
func (e SysRole) Update2Status(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.UpdateStatus(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

// Update2DataScope 更新角色数据权限
// @Summary 更新角色数据权限
// @Description 获取JSON
// @Tags 角色管理
// @Accept  application/json
// @Product application/json
// @Param data body dto.RoleDataScopeReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/roledatascope [put]
// @Security Bearer
func (e SysRole) Update2DataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
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
	err = s.UpdateDataScope(&req).Error
	if err != nil {
		e.Error(500, err, fmt.Sprintf("更新角色数据权限失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(nil, "操作成功")
}

// SetRoleMenus 设置角色与菜单的绑定关系
// @Summary 设置角色与菜单的绑定关系
// @Description 设置角色与菜单的绑定关系
// @Tags 角色管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SetRoleMenusReq true "body"
// @Success 200 {object} response.Response{message=string} "{\"code\": 0, \"message\": [...]}"
// @Router /lotus/api/v1/role-menu [put]
// @Security Bearer
func (e SysRole) SetRoleMenus(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SetRoleMenusReq{}
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

	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.SetRoleMenus(req.RoleId, req.MenuIds, cb)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("设置角色菜单关系失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(nil, "设置角色菜单关系成功")
}
