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

type SysOperaLog struct {
	models.Model
	Title         string    `json:"title" gorm:"column:title;type:varchar(255);size:255;comment:操作模块"`
	BusinessType  string    `json:"businessType" gorm:"column:business_type;type:varchar(128);size:128;comment:操作类型"`
	BusinessTypes string    `json:"businessTypes" gorm:"column:business_types;type:varchar(128);size:128;comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"column:method;type:varchar(128);size:128;comment:函数"`
	RequestMethod string    `json:"requestMethod" gorm:"column:request_method;type:varchar(128);size:128;comment:请求方式 GET POST PUT DELETE"`
	OperatorType  string    `json:"operatorType" gorm:"column:operator_type;type:varchar(128);size:128;comment:操作类型"`
	OperName      string    `json:"operName" gorm:"column:oper_name;type:varchar(128);size:128;comment:操作者"`
	DeptName      string    `json:"deptName" gorm:"column:dept_name;type:varchar(128);size:128;comment:部门名称"`
	OperUrl       string    `json:"operUrl" gorm:"column:oper_url;type:varchar(255);size:255;comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"column:oper_ip;type:varchar(128);size:128;comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"column:oper_location;type:varchar(128);size:128;comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"column:oper_param;type:text;comment:请求参数"`
	Status        string    `json:"status" gorm:"column:status;type:tinyint;size:4;comment:操作状态 1:正常 2:关闭"`
	OperTime      time.Time `json:"operTime" gorm:"column:oper_time;type:datetime;comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"column:json_result;type:json;comment:返回数据"`
	Remark        string    `json:"remark" gorm:"column:remark;type:varchar(255);size:255;comment:备注"`
	LatencyTime   string    `json:"latencyTime" gorm:"column:latency_time;type:varchar(128);size:128;comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"column:user_agent;type:varchar(255);size:255;comment:ua"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:最后更新时间"`
	models.ControlBy
}

func (*SysOperaLog) TableName() string {
	return "sys_opera_log"
}

func (e *SysOperaLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysOperaLog) GetId() interface{} {
	return e.Id
}

// SysOperaLog字段常量定义 - 用于GORM查询和函数调用
const (
	SysOperaLogId            = "id"
	SysOperaLogTitle         = "title"
	SysOperaLogBusinessType  = "business_type"
	SysOperaLogBusinessTypes = "business_types"
	SysOperaLogMethod        = "method"
	SysOperaLogRequestMethod = "request_method"
	SysOperaLogOperatorType  = "operator_type"
	SysOperaLogOperName      = "oper_name"
	SysOperaLogDeptName      = "dept_name"
	SysOperaLogOperUrl       = "oper_url"
	SysOperaLogOperIp        = "oper_ip"
	SysOperaLogOperLocation  = "oper_location"
	SysOperaLogOperParam     = "oper_param"
	SysOperaLogStatus        = "status"
	SysOperaLogOperTime      = "oper_time"
	SysOperaLogJsonResult    = "json_result"
	SysOperaLogRemark        = "remark"
	SysOperaLogLatencyTime   = "latency_time"
	SysOperaLogUserAgent     = "user_agent"
	SysOperaLogCreatedAt     = "created_at"
	SysOperaLogUpdatedAt     = "updated_at"
)

// SaveOperaLog 从队列中获取操作日志
func SaveOperaLog(message storage.Messager) (err error) {
	//准备db
	db := sdk.Runtime.GetDbByKey(message.GetPrefix())
	if db == nil {
		err = errors.New("db not exist")
		log.Errorf("host[%s]'s %s", message.GetPrefix(), err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		log.Errorf("json Marshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	var l SysOperaLog
	err = json.Unmarshal(rb, &l)
	if err != nil {
		log.Errorf("json Unmarshal error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	// JsonResult 已改为 JSON 类型，如果为空则设置为 null
	if l.JsonResult == "" {
		l.JsonResult = "{}"
	}
	err = db.Create(&l).Error
	if err != nil {
		log.Errorf("db create error, %s", err.Error())
		// Log writing to the database ignores error
		return nil
	}
	return nil
}
