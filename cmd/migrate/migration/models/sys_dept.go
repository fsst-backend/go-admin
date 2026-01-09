package models

import "go-admin/common/models"

type SysDept struct {
	DeptId      int    `json:"deptId" gorm:"column:dept_id;type:int;primaryKey;autoIncrement;"` //部门编码
	ParentId    int    `json:"parentId" gorm:"column:parent_id;type:int;"`                      //上级部门
	DeptPath    string `json:"deptPath" gorm:"column:dept_path;type:varchar(255);"`             //
	DeptName    string `json:"deptName"  gorm:"column:dept_name;type:varchar(128);"`            //部门名称
	DeptCatalog string `json:"deptCatalog" gorm:"column:dept_catalog;type:varchar(128);"`       //部门类型 finance 财务  hr 人力资源 customer_service 客服
	Sort        int    `json:"sort" gorm:"column:sort;type:int;"`                               //排序
	Leader      string `json:"leader" gorm:"column:leader;type:varchar(128);"`                  //负责人 对应的uuid
	Phone       string `json:"phone" gorm:"column:phone;type:varchar(11);"`                     //手机
	Email       string `json:"email" gorm:"column:email;type:varchar(64);"`                     //邮箱
	Status      int    `json:"status" gorm:"column:status;type:tinyint;"`                       //状态  1 启用 0 未启用
	models.ControlBy
	models.ModelTime
}

func (SysDept) TableName() string {
	return "sys_dept"
}
