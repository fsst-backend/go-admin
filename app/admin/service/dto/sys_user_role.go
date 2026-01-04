package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysUserRoleGetPageReq 用户与角色关系分页查询
// 表：sys_user_role

type SysUserRoleGetPageReq struct {
	dto.OffsetLimitPagination `search:"-"`
	UserId                    *int `form:"userId" search:"type:exact;column:user_id;table:sys_user_role" comment:"用户ID"`
	RoleId                    *int `form:"roleId" search:"type:exact;column:role_id;table:sys_user_role" comment:"角色ID"`
	SysUserRoleOrder
}

type SysUserRoleOrder struct {
	IdOrder     string `form:"idOrder"     search:"type:order;column:id;table:sys_user_role"`
	UserIdOrder string `form:"userIdOrder" search:"type:order;column:user_id;table:sys_user_role"`
	RoleIdOrder string `form:"roleIdOrder" search:"type:order;column:role_id;table:sys_user_role"`
}

func (m *SysUserRoleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysUserRoleInsertReq 新增用户与角色关系

type SysUserRoleInsertReq struct {
	Id     int `json:"id" comment:"主键编码"`
	UserId int `json:"userId" comment:"用户ID" vd:"$>0"`
	RoleId int `json:"roleId" comment:"角色ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysUserRoleInsertReq) Generate(model *models.SysUserRole) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.UserId = s.UserId
	model.RoleId = s.RoleId
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysUserRoleInsertReq) GetId() interface{} {
	return s.Id
}

// SysUserRoleUpdateReq 修改用户与角色关系

type SysUserRoleUpdateReq struct {
	Id     int `json:"id" comment:"主键编码"`
	UserId int `json:"userId" comment:"用户ID" vd:"$>0"`
	RoleId int `json:"roleId" comment:"角色ID" vd:"$>0"`
	common.ControlBy
}

func (s *SysUserRoleUpdateReq) Generate(model *models.SysUserRole) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.UserId = s.UserId
	model.RoleId = s.RoleId
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysUserRoleUpdateReq) GetId() interface{} {
	return s.Id
}

// SysUserRoleGetReq 获取单个用户与角色关系

type SysUserRoleGetReq struct {
	Id int `uri:"id"`
}

func (s *SysUserRoleGetReq) GetId() interface{} {
	return s.Id
}

// SysUserRoleDeleteReq 删除用户与角色关系

type SysUserRoleDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysUserRoleDeleteReq) GetId() interface{} {
	return s.Ids
}
