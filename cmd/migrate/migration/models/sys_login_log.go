package models

import (
	"time"

	"go-admin/common/models"
)

type SysLoginLog struct {
	models.Model
	Username      string    `json:"username" gorm:"column:username;type:varchar(128);comment:用户名"`
	Status        string    `json:"status" gorm:"column:status;type:tinyint;comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"column:ipaddr;type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"column:login_location;type:varchar(255);comment:归属地"`
	Browser       string    `json:"browser" gorm:"column:browser;type:varchar(255);comment:浏览器"`
	Os            string    `json:"os" gorm:"column:os;type:varchar(255);comment:系统"`
	Platform      string    `json:"platform" gorm:"column:platform;type:varchar(255);comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"column:login_time;type:datetime;comment:登录时间"`
	Remark       string    `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注"`
	Msg          string    `json:"msg" gorm:"column:msg;type:varchar(255);comment:信息"`
	TokenVersion int       `json:"tokenVersion" gorm:"column:token_version;type:int;default:0;comment:登录令牌版本"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:最后更新时间"`
	models.ControlBy
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}
