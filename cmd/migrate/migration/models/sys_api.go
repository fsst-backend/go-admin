package models

import "go-admin/common/models"

type SysApi struct {
	Id     int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"column:handle;type:varchar(128);comment:handle"` // 启用了，需要使用单独的权限报承载
	Title  string `json:"title" gorm:"column:title;type:varchar(128);comment:标题"`
	Path   string `json:"path" gorm:"column:path;type:varchar(128);comment:地址"`
	Action string `json:"action" gorm:"column:action;type:varchar(16);comment:请求类型"`
	Type   string `json:"type" gorm:"column:type;type:varchar(16);comment:接口类型"`
	Tag    string `json:"tag" gorm:"column:tag;type:varchar(128);comment:标签"`
	models.ModelTime
	models.ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}
