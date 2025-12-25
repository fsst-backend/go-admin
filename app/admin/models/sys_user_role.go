package models

import "go-admin/common/models"

type SysUserRole struct {
	Id     int `gorm:"primaryKey;autoIncrement;comment:主键编码" json:"id"`
	UserId int `gorm:"size:11;index;comment:用户ID" json:"userId"`
	RoleId int `gorm:"size:11;index;comment:角色ID" json:"roleId"`
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
