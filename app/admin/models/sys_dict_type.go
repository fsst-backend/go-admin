package models

import (
	"go-admin/common/models"
)

type SysDictType struct {
	ID       int    `json:"id" gorm:"column:dict_id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	DictName string `json:"dictName" gorm:"column:dict_name;type:varchar(128);comment:DictName"`
	DictType string `json:"dictType" gorm:"column:dict_type;type:varchar(128);comment:DictType"`
	Status   int    `json:"status" gorm:"column:status;type:tinyint;comment:Status"`
	Remark   string `json:"remark" gorm:"column:remark;type:varchar(255);comment:Remark"`
	models.ControlBy
	models.ModelTime
}

func (*SysDictType) TableName() string {
	return "sys_dict_type"
}

func (e *SysDictType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDictType) GetId() interface{} {
	return e.ID
}

// SysDictType字段常量定义 - 用于GORM查询和函数调用
const (
	SysDictTypeId       = "dict_id"
	SysDictTypeDictName = "dict_name"
	SysDictTypeDictType = "dict_type"
	SysDictTypeStatus   = "status"
	SysDictTypeRemark   = "remark"
)
