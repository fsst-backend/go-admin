package models

import "go-admin/common/models"

type SysRoleMenu struct {
	Id     int `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	RoleId int `json:"roleId" gorm:"column:role_id;type:int;comment:角色编码"`
	MenuId int `json:"menuId" gorm:"column:menu_id;type:int;comment:菜单编码"`
	models.ControlBy
	models.ModelTime
}

func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
