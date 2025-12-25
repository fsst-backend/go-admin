package models

import "go-admin/common/models"

type SysRoleDept struct {
	Id     int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	RoleId int `json:"roleId" gorm:"column:role_id;size:20;comment:角色编码"`
	DeptId int `json:"deptId" gorm:"column:dept_id;size:20;comment:部门编码"`
	models.ControlBy
	models.ModelTime
}

// SysRoleDept字段常量定义 - 用于GORM查询和函数调用
const (
	SysRoleDeptId     = "id"
	SysRoleDeptRoleId = "role_id"
	SysRoleDeptDeptId = "dept_id"
)

func (*SysRoleDept) TableName() string {
	return "sys_role_dept"
}

func (e *SysRoleDept) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRoleDept) GetId() interface{} {
	return e.Id
}
