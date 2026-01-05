package models

import "go-admin/common/models"

type SysRole struct {
	RoleId      int             `json:"roleId" gorm:"column:role_id;type:int;primaryKey;autoIncrement"` // 角色编码
	RoleName    string          `json:"roleName" gorm:"column:role_name;type:varchar(128);size:128;"`   // 角色名称
	Status      string          `json:"status" gorm:"column:status;type:tinyint;size:4;"`               // 状态 1禁用 2正常
	RoleKey     string          `json:"roleKey" gorm:"column:role_key;type:varchar(128);size:128;"`     //角色代码
	RoleSort    int             `json:"roleSort" gorm:"column:role_sort;type:int;"`                     //角色排序
	Flag        string          `json:"flag" gorm:"column:flag;type:varchar(128);size:128;"`            //
	Remark      string          `json:"remark" gorm:"column:remark;type:varchar(255);size:255;"`        //备注
	Admin       bool            `json:"admin" gorm:"column:admin;type:tinyint;size:4;"`
	DataScope   string          `json:"dataScope" gorm:"column:data_scope;type:varchar(128);size:128;"` // 数据范围
	Permissions []SysPermission `json:"permissions,omitempty" gorm:"-"`
	Menus       []SysMenu       `json:"Menus,omitempty" gorm:"-"`
	models.ControlBy
	models.ModelTime
}

func (*SysRole) TableName() string {
	return "sys_role"
}

func (e *SysRole) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysRole) GetId() interface{} {
	return e.RoleId
}

// SysRole字段常量定义 - 用于GORM查询和函数调用
const (
	SysRoleRoleId    = "role_id"
	SysRoleRoleName  = "role_name"
	SysRoleStatus    = "status"
	SysRoleRoleKey   = "role_key"
	SysRoleRoleSort  = "role_sort"
	SysRoleFlag      = "flag"
	SysRoleRemark    = "remark"
	SysRoleAdmin     = "admin"
	SysRoleDataScope = "data_scope"
)
