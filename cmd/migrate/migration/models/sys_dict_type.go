package models

import "go-admin/common/models"

type DictType struct {
	ID       int    `json:"id" gorm:"column:dict_id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	DictName string `json:"dictName" gorm:"column:dict_name;type:varchar(128);comment:DictName"`
	DictType string `json:"dictType" gorm:"column:dict_type;type:varchar(128);comment:DictType"`
	Status   int    `json:"status" gorm:"column:status;type:tinyint;comment:Status"`
	Remark   string `json:"remark" gorm:"column:remark;type:varchar(255);comment:Remark"`
	models.ControlBy
	models.ModelTime
}

func (DictType) TableName() string {
	return "sys_dict_type"
}
