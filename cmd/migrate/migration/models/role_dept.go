package models

import "go-admin/common/models"

type SysRoleDept struct {
	Id     int `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	RoleId int `json:"roleId" gorm:"column:role_id;type:int;comment:角色编码"`
	DeptId int `json:"deptId" gorm:"column:dept_id;type:int;comment:部门编码"`
	models.ControlBy
	models.ModelTime
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
