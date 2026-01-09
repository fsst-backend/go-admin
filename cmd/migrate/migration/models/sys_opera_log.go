package models

import (
	"time"

	"go-admin/common/models"
)

type SysOperaLog struct {
	models.Model
	Title         string    `json:"title" gorm:"column:title;type:varchar(255);comment:操作模块"`
	BusinessType  string    `json:"businessType" gorm:"column:business_type;type:varchar(128);comment:操作类型"`
	BusinessTypes string    `json:"businessTypes" gorm:"column:business_types;type:varchar(128);comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"column:method;type:varchar(128);comment:函数"`
	RequestMethod string    `json:"requestMethod" gorm:"column:request_method;type:varchar(128);comment:请求方式 GET POST PUT DELETE"`
	OperatorType  string    `json:"operatorType" gorm:"column:operator_type;type:varchar(128);comment:操作类型"`
	OperName      string    `json:"operName" gorm:"column:oper_name;type:varchar(128);comment:操作者"`
	DeptName      string    `json:"deptName" gorm:"column:dept_name;type:varchar(128);comment:部门名称"`
	OperUrl       string    `json:"operUrl" gorm:"column:oper_url;type:varchar(255);comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"column:oper_ip;type:varchar(128);comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"column:oper_location;type:varchar(128);comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"column:oper_param;type:text;comment:请求参数"`
	Status        string    `json:"status" gorm:"column:status;type:tinyint;comment:操作状态 1:正常 2:关闭"`
	OperTime      time.Time `json:"operTime" gorm:"column:oper_time;type:datetime;comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"column:json_result;type:json;comment:返回数据"`
	Remark        string    `json:"remark" gorm:"column:remark;type:varchar(255);comment:备注"`
	LatencyTime   string    `json:"latencyTime" gorm:"column:latency_time;type:varchar(128);comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"column:user_agent;type:varchar(255);comment:ua"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:最后更新时间"`
	models.ControlBy
}

func (SysOperaLog) TableName() string {
	return "sys_opera_log"
}
