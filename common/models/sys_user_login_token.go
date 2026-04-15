package models

// SysUserLoginToken 按 (用户, 登录渠道) 记录单设备登录版本，加渠道只需常量 + 客户端传 loginChannel
type SysUserLoginToken struct {
	UserID       int    `gorm:"column:user_id;primaryKey;comment:用户ID"`
	LoginChannel string `gorm:"column:login_channel;type:varchar(32);primaryKey;comment:登录渠道"`
	TokenVersion int    `gorm:"column:token_version;default:0;comment:该渠道会话版本，新登录递增"`
}

func (SysUserLoginToken) TableName() string {
	return "sys_user_login_token"
}
