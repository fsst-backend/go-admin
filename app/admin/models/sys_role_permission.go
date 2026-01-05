package models

import "go-admin/common/models"

// SysRolePermission 角色与权限关系表
// 不使用数据库外键，所有关联在代码中通过 role_id 和 permission_id 手动处理

type SysRolePermission struct {
	Id           int `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	RoleId       int `json:"roleId" gorm:"column:role_id;type:int;size:11;index;comment:角色ID"`
	PermissionId int `json:"permissionId" gorm:"column:permission_id;type:int;size:20;index;comment:权限ID"`
	models.ControlBy
	models.ModelTime
}

func (SysRolePermission) TableName() string {
	return "sys_role_permission"
}

func (e *SysRolePermission) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRolePermission) GetId() interface{} {
	return e.Id
}

// SysRolePermission字段常量定义 - 用于GORM查询和函数调用
const (
	SysRolePermissionId           = "id"
	SysRolePermissionRoleId       = "role_id"
	SysRolePermissionPermissionId = "permission_id"
)
