package models

import "go-admin/common/models"

type SysUserRole struct {
	Id     int `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码" json:"id"`
	UserId int `gorm:"column:user_id;type:int;index;comment:用户ID" json:"userId"`
	RoleId int `gorm:"column:role_id;type:int;index;comment:角色ID" json:"roleId"`
	models.ControlBy
	models.ModelTime
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

func (e *SysUserRole) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUserRole) GetId() interface{} {
	return e.Id
}

// SysUserRole字段常量定义 - 用于GORM查询和函数调用
const (
	SysUserRoleId = "id"
	SysUserUserId = "user_id"
	SysUserRoleID = "role_id"
)
