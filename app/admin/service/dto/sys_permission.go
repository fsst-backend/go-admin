package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// SysPermissionGetPageReq 权限定义表分页查询请求
// 仅查询字段参与 search，保持与表结构列名一致
// 表：sys_permission

type SysPermissionGetPageReq struct {
	dto.Pagination `search:"-"`
	Code           string `form:"code"  search:"type:contains;column:code;table:sys_permission" comment:"权限唯一编码"`
	Name           string `form:"name"  search:"type:contains;column:name;table:sys_permission" comment:"权限名称"`
	Type           string `form:"type"  search:"type:exact;column:type;table:sys_permission" comment:"权限类型(menu/button/api/page)"`
	Status         int    `form:"status" search:"type:exact;column:status;table:sys_permission" comment:"状态 1启用 0禁用"`
	SysPermissionOrder
}

type SysPermissionOrder struct {
	IdOrder        string `form:"idOrder"        search:"type:order;column:id;table:sys_permission"`
	CodeOrder      string `form:"codeOrder"      search:"type:order;column:code;table:sys_permission"`
	SortOrder      string `form:"sortOrder"      search:"type:order;column:sort;table:sys_permission"`
	CreatedAtOrder string `form:"createdAtOrder" search:"type:order;column:created_at;table:sys_permission"`
}

func (m *SysPermissionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// SysPermissionInsertReq 新增权限

type SysPermissionInsertReq struct {
	Id       int    `json:"id" comment:"主键编码"`
	Code     string `json:"code"     comment:"权限唯一编码" vd:"len($)>0"`
	Name     string `json:"name"     comment:"权限名称" vd:"len($)>0"`
	Type     string `json:"type"     comment:"权限类型(menu/button/api/page)" vd:"len($)>0"`
	ParentId int    `json:"parentId" comment:"父权限ID"`
	Sort     int    `json:"sort"     comment:"排序"`
	Status   int    `json:"status"   comment:"状态 1启用 0禁用"`
	Remark   string `json:"remark"   comment:"备注说明"`
	ApiIds   []int  `json:"apiIds"   comment:"关联的API ID列表"`
	common.ControlBy
}

func (s *SysPermissionInsertReq) Generate(model *models.SysPermission) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Code = s.Code
	model.Name = s.Name
	model.Type = s.Type
	model.ParentId = s.ParentId
	model.Sort = s.Sort
	model.Status = s.Status
	model.Remark = s.Remark
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysPermissionInsertReq) GetId() interface{} {
	return s.Id
}

// SysPermissionUpdateReq 修改权限

type SysPermissionUpdateReq struct {
	Id       int    `json:"id" comment:"主键编码"`
	Code     string `json:"code"     comment:"权限唯一编码" vd:"len($)>0"`
	Name     string `json:"name"     comment:"权限名称" vd:"len($)>0"`
	Type     string `json:"type"     comment:"权限类型(menu/button/api/page)" vd:"len($)>0"`
	ParentId int    `json:"parentId" comment:"父权限ID"`
	Sort     int    `json:"sort"     comment:"排序"`
	Status   int    `json:"status"   comment:"状态 1启用 0禁用"`
	Remark   string `json:"remark"   comment:"备注说明"`
	ApiIds   []int  `json:"apiIds"   comment:"关联的API ID列表"`
	common.ControlBy
}

func (s *SysPermissionUpdateReq) Generate(model *models.SysPermission) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Code = s.Code
	model.Name = s.Name
	model.Type = s.Type
	model.ParentId = s.ParentId
	model.Sort = s.Sort
	model.Status = s.Status
	model.Remark = s.Remark
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysPermissionUpdateReq) GetId() interface{} {
	return s.Id
}

// SysPermissionGetReq 获取单个权限

type SysPermissionGetReq struct {
	Id int `form:"id"`
}

func (s *SysPermissionGetReq) GetId() interface{} {
	return s.Id
}

// SysPermissionDeleteReq 删除权限

type SysPermissionDeleteReq struct {
	Ids []int `json:"ids"`
	common.ControlBy
}

func (s *SysPermissionDeleteReq) GetId() interface{} {
	return s.Ids
}
