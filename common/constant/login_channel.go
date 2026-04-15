package constant

// 登录渠道常量（与 sys_user_login_token.login_channel 对应）。
// 新端可在 loginChannel 传自定义小写串（≤32），无需加列；建议在常量中登记便于文档与校验。
const (
	LoginChannelAdmin = "admin"   // 后台 Web / 默认
	LoginChannelCSApp = "cs_app" // 客服 App
)
