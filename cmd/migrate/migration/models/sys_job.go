package models

import "go-admin/common/models"

type SysJob struct {
	JobId          int    `json:"jobId" gorm:"column:job_id;type:int;primaryKey;autoIncrement"`    // 编码
	JobName        string `json:"jobName" gorm:"column:job_name;type:varchar(255);"`               // 名称
	JobGroup       string `json:"jobGroup" gorm:"column:job_group;type:varchar(255);"`             // 任务分组
	JobType        int    `json:"jobType" gorm:"column:job_type;type:tinyint;"`                    // 任务类型
	CronExpression string `json:"cronExpression" gorm:"column:cron_expression;type:varchar(255);"` // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"column:invoke_target;type:varchar(255);"`     // 调用目标
	Args           string `json:"args" gorm:"column:args;type:varchar(255);"`                      // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"column:misfire_policy;type:int;"`            // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"column:concurrent;type:tinyint;"`               // 是否并发
	Status         int    `json:"status" gorm:"column:status;type:tinyint;"`                       // 状态
	EntryId        int    `json:"entry_id" gorm:"column:entry_id;type:int;"`                       // job启动时返回的id
	models.ModelTime
	models.ControlBy
}

func (SysJob) TableName() string {
	return "sys_job"
}
