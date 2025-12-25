package models

type SysRoleMenu struct {
	Id     int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RoleId int `json:"roleId" gorm:"column:role_id;size:20;comment:角色编码"`
	MenuId int `json:"menuId" gorm:"column:menu_id;size:20;comment:菜单编码"`
	ControlBy
	ModelTime
}

func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}