package models

import "go-admin/common/models"

// SysPermissionApi 权限与API关系表
// 不使用数据库外键，所有关联在代码中通过 permission_id 和 api_id 手动处理

type SysPermissionApi struct {
	Id           int `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	PermissionId int `json:"permissionId" gorm:"column:permission_id;type:int;index;comment:权限ID"`
	ApiId        int `json:"apiId" gorm:"column:api_id;type:int;index;comment:API ID"`
	models.ControlBy
	models.ModelTime
}

func (SysPermissionApi) TableName() string {
	return "sys_permission_api"
}

func (e *SysPermissionApi) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysPermissionApi) GetId() interface{} {
	return e.Id
}

// SysPermissionApi字段常量定义 - 用于GORM查询和函数调用
const (
	SysPermissionApiId           = "id"
	SysPermissionApiPermissionId = "permission_id"
	SysPermissionApiApiId        = "api_id"
)
