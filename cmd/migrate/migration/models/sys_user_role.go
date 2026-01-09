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
