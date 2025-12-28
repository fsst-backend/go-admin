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
	"go-admin/common/actions"
)

type SysApi struct {
	api.Api
}

// GetPage 获取接口管理列表
// @Summary 获取接口管理列表
// @Description 获取接口管理列表
// @Tags 接口管理
// @Param name query string false "名称"
// @Param title query string false "标题"
// @Param path query string false "地址"
// @Param action query string false "类型"
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {object} response.Response{message=response.Page{list=[]models.SysApi}} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-api [get]
// @Security Bearer
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiGetPageReq{}
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
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysApi, 0)
	var count int64
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetOffset(), req.GetLimit(), "查询成功")
}

// Get 获取接口管理
// @Summary 获取接口管理
// @Description 获取接口管理
// @Tags 接口管理
// @Success 200 {object} response.Response{message=models.SysApi} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-api/get [get]
// @Security Bearer
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
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
	p := actions.GetPermissionFromContext(c)
	var object models.SysApi
	err = s.Get(&req, p, &object).Error
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Update 修改接口管理
// @Summary 修改接口管理
// @Description 修改接口管理
// @Tags 接口管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiUpdateReq true "body"
// @Success 200 {object} response.Response{message=string}	"{"code": 0, "message": "修改成功"}"
// @Router /lotus/api/v1/sys-api [put]
// @Security Bearer
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)
	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// DeleteSysApi 删除接口管理
// @Summary 删除接口管理
// @Description 删除接口管理
// @Tags 接口管理
// @Param data body dto.SysApiDeleteReq true "body"
// @Success 200 {object} response.Response{message=string}	"{"code": 0, "message": "删除成功"}"
// @Router /lotus/api/v1/sys-api [delete]
// @Security Bearer
func (e SysApi) DeleteSysApi(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}
