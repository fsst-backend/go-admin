package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/global"
)

type SysPermission struct {
	api.Api
}

// GetPage 获取权限定义列表
// @Summary 获取权限定义列表
// @Description 获取权限定义列表
// @Tags 权限定义
// @Param code query string false "权限编码"
// @Param name query string false "权限名称"
// @Param type query string false "权限类型"
// @Param status query int false "状态"
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {object} response.Response{message=response.Page{list=[]models.SysPermission}} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-permission [get]
// @Security Bearer
func (e SysPermission) GetPage(c *gin.Context) {
	s := service.SysPermission{}
	req := dto.SysPermissionGetPageReq{}
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

	list := make([]models.SysPermission, 0)
	var count int64
	if err = s.GetPage(&req, &list, &count); err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetOffset(), req.GetLimit(), "查询成功")
}

// Get 获取单个权限定义
// @Summary 获取单个权限定义
// @Description 获取单个权限定义
// @Tags 权限定义
// @Success 200 {object} response.Response{message=models.SysPermission} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-permission/get [get]
// @Security Bearer
func (e SysPermission) Get(c *gin.Context) {
	req := dto.SysPermissionGetReq{}
	s := service.SysPermission{}
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
	var object models.SysPermission
	if err = s.Get(&req, &object); err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert 创建权限定义
// @Summary 创建权限定义
// @Description 创建权限定义
// @Tags 权限定义
// @Accept application/json
// @Product application/json
// @Param data body dto.SysPermissionInsertReq true "data"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-permission [post]
// @Security Bearer
func (e SysPermission) Insert(c *gin.Context) {
	s := service.SysPermission{}
	req := dto.SysPermissionInsertReq{}
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
	req.SetCreateBy(user.GetUserId(c))
	if err = s.Insert(&req); err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改权限定义
// @Summary 修改权限定义
// @Description 修改权限定义
// @Tags 权限定义
// @Accept application/json
// @Product application/json
// @Param data body dto.SysPermissionUpdateReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-permission [put]
// @Security Bearer
func (e SysPermission) Update(c *gin.Context) {
	s := service.SysPermission{}
	req := dto.SysPermissionUpdateReq{}
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

	// 获取 Casbin Enforcer
	cb, err := global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "加载权限策略失败")
		return
	}

	if err = s.Update(&req, cb); err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete 删除权限定义
// @Summary 删除权限定义
// @Description 删除权限定义
// @Tags 权限定义
// @Param data body dto.SysPermissionDeleteReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-permission [delete]
// @Security Bearer
func (e SysPermission) Delete(c *gin.Context) {
	s := service.SysPermission{}
	req := dto.SysPermissionDeleteReq{}
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
	req.SetUpdateBy(user.GetUserId(c))

	// 获取 Casbin Enforcer
	cb, err := global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "加载权限策略失败")
		return
	}

	if err = s.Remove(&req, cb); err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
