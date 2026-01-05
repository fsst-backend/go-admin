package models

import (
	"go-admin/common/models"
)

type SysConfig struct {
	models.Model
	ConfigName  string `json:"configName" gorm:"column:config_name;type:varchar(128);size:128;comment:ConfigName"`    //
	ConfigKey   string `json:"configKey" gorm:"column:config_key;type:varchar(128);size:128;comment:ConfigKey"`       //
	ConfigValue string `json:"configValue" gorm:"column:config_value;type:varchar(255);size:255;comment:ConfigValue"` //
	ConfigType  string `json:"configType" gorm:"column:config_type;type:varchar(64);size:64;comment:ConfigType"`
	IsFrontend  string `json:"isFrontend" gorm:"column:is_frontend;type:varchar(64);size:64;comment:是否前台"` //
	Remark      string `json:"remark" gorm:"column:remark;type:varchar(128);size:128;comment:Remark"`      //
	models.ControlBy
	models.ModelTime
}

func (*SysConfig) TableName() string {
	return "sys_config"
}

func (e *SysConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysConfig) GetId() interface{} {
	return e.Id
}

// SysConfig字段常量定义 - 用于GORM查询和函数调用
const (
	SysConfigId          = "id"
	SysConfigConfigName  = "config_name"
	SysConfigConfigKey   = "config_key"
	SysConfigConfigValue = "config_value"
	SysConfigConfigType  = "config_type"
	SysConfigIsFrontend  = "is_frontend"
	SysConfigRemark      = "remark"
)
