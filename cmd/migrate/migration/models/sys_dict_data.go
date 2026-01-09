package models

import "go-admin/common/models"

type DictData struct {
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

func (DictData) TableName() string {
	return "sys_dict_data"
}
