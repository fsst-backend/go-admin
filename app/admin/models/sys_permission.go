package models

import "go-admin/common/models"

// SysPermission 权限定义表
// 不使用数据库外键，所有关联在代码中处理

type SysPermission struct {
	Id       int      `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	Code     string   `json:"code" gorm:"column:code;type:varchar(128);size:128;not null;uniqueIndex:uk_permission_code;comment:权限唯一编码"`
	Name     string   `json:"name" gorm:"column:name;type:varchar(128);size:128;not null;comment:权限名称"`
	Type     string   `json:"type" gorm:"column:type;type:varchar(16);size:16;not null;comment:权限类型(menu/button/api/page)"`
	ParentId int      `json:"parentId" gorm:"column:parent_id;type:int;size:20;default:0;comment:父权限ID（权限树）"`
	Sort     int      `json:"sort" gorm:"column:sort;type:int;default:0;comment:排序"`
	Status   int      `json:"status" gorm:"column:status;type:tinyint;size:4;default:1;comment:状态 1启用 0禁用"`
	Remark   string   `json:"remark" gorm:"column:remark;type:varchar(255);size:255;comment:备注说明"`
	Apis     []SysApi `json:"apis,omitempty" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (SysPermission) TableName() string {
	return "sys_permission"
}

func (e *SysPermission) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysPermission) GetId() interface{} {
	return e.Id
}

// SysPermission字段常量定义 - 用于GORM查询和函数调用
const (
	SysPermissionId       = "id"
	SysPermissionCode     = "code"
	SysPermissionName     = "name"
	SysPermissionType     = "type"
	SysPermissionParentId = "parent_id"
	SysPermissionSort     = "sort"
	SysPermissionStatus   = "status"
	SysPermissionRemark   = "remark"
)
