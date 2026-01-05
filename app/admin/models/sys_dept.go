package models

import "go-admin/common/models"

type SysDept struct {
	DeptId   int    `json:"deptId" gorm:"column:dept_id;type:int;primaryKey;autoIncrement;"` //部门编码
	ParentId int    `json:"parentId" gorm:"column:parent_id;type:int;"`                      //上级部门
	DeptPath string `json:"deptPath" gorm:"column:dept_path;type:varchar(255);size:255;"`    //
	DeptName string `json:"deptName"  gorm:"column:dept_name;type:varchar(128);size:128;"`   //部门名称
	Sort     int    `json:"sort" gorm:"column:sort;type:int;size:4;"`                        //排序
	Leader   string `json:"leader" gorm:"column:leader;type:varchar(128);size:128;"`         //负责人
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(11);size:11;"`             //手机
	Email    string `json:"email" gorm:"column:email;type:varchar(64);size:64;"`             //邮箱
	Status   int    `json:"status" gorm:"column:status;type:tinyint;size:4;"`                //状态
	models.ControlBy
	models.ModelTime
	DataScope string    `json:"dataScope" gorm:"-"`
	Params    string    `json:"params" gorm:"-"`
	Children  []SysDept `json:"children" gorm:"-"`
}

func (*SysDept) TableName() string {
	return "sys_dept"
}

func (e *SysDept) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDept) GetId() interface{} {
	return e.DeptId
}

// SysDept字段常量定义 - 用于GORM查询和函数调用
const (
	SysDeptDeptId   = "dept_id"
	SysDeptParentId = "parent_id"
	SysDeptDeptPath = "dept_path"
	SysDeptDeptName = "dept_name"
	SysDeptSort     = "sort"
	SysDeptLeader   = "leader"
	SysDeptPhone    = "phone"
	SysDeptEmail    = "email"
	SysDeptStatus   = "status"
)
