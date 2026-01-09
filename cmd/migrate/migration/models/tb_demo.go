package models

import "go-admin/common/models"

type TbDemo struct {
	models.Model
	Name string `json:"name" gorm:"column:name;type:varchar(128);comment:名称"`
	models.ModelTime
	models.ControlBy
}

func (TbDemo) TableName() string {
	return "tb_demo"
}
