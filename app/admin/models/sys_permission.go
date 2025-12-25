package models

import "go-admin/common/models"

// SysPermission 权限定义表
// 不使用数据库外键，所有关联在代码中处理

type SysPermission struct {
	Id       int      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Code     string   `json:"code" gorm:"size:128;not null;uniqueIndex:uk_permission_code;comment:权限唯一编码"`
	Name     string   `json:"name" gorm:"size:128;not null;comment:权限名称"`
	Type     string   `json:"type" gorm:"size:16;not null;comment:权限类型(menu/button/api/page)"`
	ParentId int      `json:"parentId" gorm:"size:20;default:0;comment:父权限ID（权限树）"`
	Sort     int      `json:"sort" gorm:"default:0;comment:排序"`
	Status   int      `json:"status" gorm:"size:4;default:1;comment:状态 1启用 0禁用"`
	Remark   string   `json:"remark" gorm:"size:255;comment:备注说明"`
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
