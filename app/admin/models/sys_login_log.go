package models

import (
	"encoding/json"
	"errors"
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/storage"

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
	Remark        string    `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注"`
	Msg           string    `json:"msg" gorm:"column:msg;type:varchar(255);comment:信息"`
	TokenVersion  int       `json:"tokenVersion" gorm:"column:token_version;type:int;default:0;comment:登录令牌版本"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:最后更新时间"`
	models.ControlBy
}

func (*SysLoginLog) TableName() string {
	return "sys_login_log"
}

func (e *SysLoginLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysLoginLog) GetId() interface{} {
	return e.Id
}

// SysLoginLog字段常量定义 - 用于GORM查询和函数调用
const (
	SysLoginLogId            = "id"
	SysLoginLogUsername      = "username"
	SysLoginLogStatus        = "status"
	SysLoginLogIpaddr        = "ipaddr"
	SysLoginLogLoginLocation = "login_location"
	SysLoginLogBrowser       = "browser"
	SysLoginLogOs            = "os"
	SysLoginLogPlatform      = "platform"
	SysLoginLogLoginTime     = "login_time"
	SysLoginLogRemark        = "remark"
	SysLoginLogMsg           = "msg"
	SysLoginLogCreatedAt     = "created_at"
	SysLoginLogUpdatedAt     = "updated_at"
)

// SaveLoginLog 从队列中获取登录日志
func SaveLoginLog(message storage.Messager) (err error) {
	//准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		return err
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		return err
	}
	var l SysLoginLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		return err
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		return err
	}
	return nil
}
