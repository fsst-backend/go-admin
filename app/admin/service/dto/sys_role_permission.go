package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysRolePermissionGetPageReq 角色与权限关系分页查询
// 表：sys_role_permission

type SysRolePermissionGetPageReq struct {
	dto.Pagination `search:"-"`
	RoleId         int `form:"roleId"         search:"type:exact;column:role_id;table:sys_role_permission" comment:"角色ID"`
	PermissionId   int `form:"permissionId"   search:"type:exact;column:permission_id;table:sys_role_permission" comment:"权限ID"`
	SysRolePermissionOrder
}

type SysRolePermissionOrder struct {
	IdOrder           string `form:"idOrder"           search:"type:order;column:id;table:sys_role_permission"`
	RoleIdOrder       string `form:"roleIdOrder"       search:"type:order;column:role_id;table:sys_role_permission"`
	PermissionIdOrder string `form:"permissionIdOrder" search:"type:order;column:permission_id;table:sys_role_permission"`
}

func (m *SysRolePermissionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysRolePermissionInsertReq 新增角色与权限关系

type SysRolePermissionInsertReq struct {
	Id           int `json:"id" comment:"主键编码"`
	RoleId       int `json:"roleId" comment:"角色ID" vd:"$>0"`
	PermissionId int `json:"permissionId" comment:"权限ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysRolePermissionInsertReq) Generate(model *models.SysRolePermission) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.RoleId = s.RoleId
	model.PermissionId = s.PermissionId
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysRolePermissionInsertReq) GetId() interface{} {
	return s.Id
}

// SysRolePermissionUpdateReq 修改角色与权限关系

type SysRolePermissionUpdateReq struct {
	Id           int `json:"id" comment:"主键编码"`
	RoleId       int `json:"roleId" comment:"角色ID" vd:"$>0"`
	PermissionId int `json:"permissionId" comment:"权限ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysRolePermissionUpdateReq) Generate(model *models.SysRolePermission) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.RoleId = s.RoleId
	model.PermissionId = s.PermissionId
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysRolePermissionUpdateReq) GetId() interface{} {
	return s.Id
}

// SysRolePermissionGetReq 获取单个角色与权限关系

type SysRolePermissionGetReq struct {
	Id int `uri:"id"`
}

func (s *SysRolePermissionGetReq) GetId() interface{} {
	return s.Id
}

// SysRolePermissionDeleteReq 删除角色与权限关系

type SysRolePermissionDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysRolePermissionDeleteReq) GetId() interface{} {
	return s.Ids
}
