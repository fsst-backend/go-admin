package models

import (
	"go-admin/common/models"
)

type SysDictData struct {
	DictCode  int    `json:"dictCode" gorm:"column:dict_code;type:int;primaryKey;autoIncrement;comment:主键编码"`
	DictSort  int    `json:"dictSort" gorm:"column:dict_sort;type:int;comment:DictSort"`
	DictLabel string `json:"dictLabel" gorm:"column:dict_label;type:varchar(128);comment:DictLabel"`
	DictValue string `json:"dictValue" gorm:"column:dict_value;type:varchar(255);comment:DictValue"`
	DictType  string `json:"dictType" gorm:"column:dict_type;type:varchar(64);comment:DictType"`
	CssClass  string `json:"cssClass" gorm:"column:css_class;type:varchar(128);comment:CssClass"`
	ListClass string `json:"listClass" gorm:"column:list_class;type:varchar(128);comment:ListClass"`
	IsDefault string `json:"isDefault" gorm:"column:is_default;type:varchar(8);comment:IsDefault"`
	Status    int    `json:"status" gorm:"column:status;type:tinyint;comment:Status"`
	Default   string `json:"default" gorm:"column:default_value;type:varchar(8);comment:Default"`
	Remark    string `json:"remark" gorm:"column:remark;type:varchar(255);comment:Remark"`
	models.ControlBy
	models.ModelTime
}

func (*SysDictData) TableName() string {
	return "sys_dict_data"
}

func (e *SysDictData) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDictData) GetId() interface{} {
	return e.DictCode
}

// SysDictData字段常量定义 - 用于GORM查询和函数调用
const (
	SysDictDataDictCode  = "dict_code"
	SysDictDataDictSort  = "dict_sort"
	SysDictDataDictLabel = "dict_label"
	SysDictDataDictValue = "dict_value"
	SysDictDataDictType  = "dict_type"
	SysDictDataCssClass  = "css_class"
	SysDictDataListClass = "list_class"
	SysDictDataIsDefault = "is_default"
	SysDictDataStatus    = "status"
	SysDictDataDefault   = "default_value"
	SysDictDataRemark    = "remark"
)
