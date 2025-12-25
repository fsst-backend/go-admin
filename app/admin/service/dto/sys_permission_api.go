package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysPermissionApiGetPageReq 权限与API关系分页查询
// 表：sys_permission_api

type SysPermissionApiGetPageReq struct {
	dto.Pagination `search:"-"`
	PermissionId   int `form:"permissionId" search:"type:exact;column:permission_id;table:sys_permission_api" comment:"权限ID"`
	ApiId          int `form:"apiId"        search:"type:exact;column:api_id;table:sys_permission_api" comment:"API ID"`
	SysPermissionApiOrder
}

type SysPermissionApiOrder struct {
	IdOrder          string `form:"idOrder"          search:"type:order;column:id;table:sys_permission_api"`
	PermissionIdOrder string `form:"permissionIdOrder" search:"type:order;column:permission_id;table:sys_permission_api"`
	ApiIdOrder       string `form:"apiIdOrder"       search:"type:order;column:api_id;table:sys_permission_api"`
}

func (m *SysPermissionApiGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysPermissionApiInsertReq 新增权限与API关系

type SysPermissionApiInsertReq struct {
	Id           int `json:"id" comment:"主键编码"`
	PermissionId int `json:"permissionId" comment:"权限ID" vd:"$>0"`
	ApiId        int `json:"apiId"        comment:"API ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysPermissionApiInsertReq) Generate(model *models.SysPermissionApi) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.PermissionId = s.PermissionId
	model.ApiId = s.ApiId
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysPermissionApiInsertReq) GetId() interface{} {
	return s.Id
}

// SysPermissionApiUpdateReq 修改权限与API关系

type SysPermissionApiUpdateReq struct {
	Id           int `json:"id" comment:"主键编码"`
	PermissionId int `json:"permissionId" comment:"权限ID" vd:"$>0"`
	ApiId        int `json:"apiId"        comment:"API ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysPermissionApiUpdateReq) Generate(model *models.SysPermissionApi) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.PermissionId = s.PermissionId
	model.ApiId = s.ApiId
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysPermissionApiUpdateReq) GetId() interface{} {
	return s.Id
}

// SysPermissionApiGetReq 获取单个权限与API关系

type SysPermissionApiGetReq struct {
	Id int `uri:"id"`
}

func (s *SysPermissionApiGetReq) GetId() interface{} {
	return s.Id
}

// SysPermissionApiDeleteReq 删除权限与API关系

type SysPermissionApiDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysPermissionApiDeleteReq) GetId() interface{} {
	return s.Ids
}
