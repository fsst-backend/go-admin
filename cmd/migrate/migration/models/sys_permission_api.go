package models

// SysPermissionApi 权限与API关系表
// 不使用数据库外键，所有关联在代码中通过 permission_id 和 api_id 手动处理

type SysPermissionApi struct {
	Id           int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	PermissionId int `json:"permissionId" gorm:"size:20;index;comment:权限ID"`
	ApiId        int `json:"apiId" gorm:"size:20;index;comment:API ID"`
	ControlBy
	ModelTime
}

func (SysPermissionApi) TableName() string {
	return "sys_permission_api"
}
