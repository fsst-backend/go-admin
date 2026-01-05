package models

import "go-admin/common/models"

type SysRoleMenu struct {
	Id     int `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	RoleId int `json:"roleId" gorm:"column:role_id;type:int;size:20;comment:角色编码"`
	MenuId int `json:"menuId" gorm:"column:menu_id;type:int;size:20;comment:菜单编码"`
	models.ControlBy
	models.ModelTime
}

// SysRoleMenu字段常量定义 - 用于GORM查询和函数调用
const (
	SysRoleMenuId     = "id"
	SysRoleMenuRoleId = "role_id"
	SysRoleMenuMenuId = "menu_id"
)

func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (e *SysRoleMenu) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRoleMenu) GetId() interface{} {
	return e.Id
}
