package models

type SysUserRole struct {
	Id     int `gorm:"primaryKey;autoIncrement;comment:主键编码" json:"id"`
	UserId int `gorm:"size:11;index;comment:用户ID" json:"userId"`
	RoleId int `gorm:"size:11;index;comment:角色ID" json:"roleId"`
	ControlBy
	ModelTime
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
