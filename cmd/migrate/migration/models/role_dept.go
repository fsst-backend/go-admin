package models

type SysRoleDept struct {
	Id     int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	RoleId int `json:"roleId" gorm:"column:role_id;size:20;comment:角色编码"`
	DeptId int `json:"deptId" gorm:"column:dept_id;size:20;comment:部门编码"`
	ControlBy
	ModelTime
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
