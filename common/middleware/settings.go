package middleware

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	{Url: "/lotus/api/v1/dict/type-option-select", Method: "GET"},
	{Url: "/lotus/api/v1/dict-data/option-select", Method: "GET"},
	{Url: "/lotus/api/v1/deptTree", Method: "GET"},
	{Url: "/lotus/api/v1/db/tables/page", Method: "GET"},
	{Url: "/lotus/api/v1/db/columns/page", Method: "GET"},
	{Url: "/lotus/api/v1/gen/toproject/:tableId", Method: "GET"},
	{Url: "/lotus/api/v1/gen/todb/:tableId", Method: "GET"},
	{Url: "/lotus/api/v1/gen/tabletree", Method: "GET"},
	{Url: "/lotus/api/v1/gen/preview/:tableId", Method: "GET"},
	{Url: "/lotus/api/v1/gen/apitofile/:tableId", Method: "GET"},
	{Url: "/lotus/api/v1/getCaptcha", Method: "GET"},
	{Url: "/lotus/api/v1/getinfo", Method: "GET"},
	{Url: "/lotus/api/v1/menuTreeselect", Method: "GET"},
	{Url: "/lotus/api/v1/menurole", Method: "GET"},
	{Url: "/lotus/api/v1/menuids", Method: "GET"},
	{Url: "/lotus/api/v1/roleMenuTreeselect/:roleId", Method: "GET"},
	{Url: "/lotus/api/v1/roleDeptTreeselect/:roleId", Method: "GET"},
	{Url: "/lotus/api/v1/refresh_token", Method: "GET"},
	{Url: "/lotus/api/v1/configKey/:configKey", Method: "GET"},
	{Url: "/lotus/api/v1/app-config", Method: "GET"},
	{Url: "/lotus/api/v1/user/profile", Method: "GET"},
	{Url: "/lotus/info", Method: "GET"},
	{Url: "/lotus/api/v1/login", Method: "POST"},
	{Url: "/lotus/api/v1/logout", Method: "POST"},
	{Url: "/lotus/api/v1/user/avatar", Method: "POST"},
	{Url: "/lotus/api/v1/user/pwd", Method: "PUT"},
	{Url: "/lotus/api/v1/metrics", Method: "GET"},
	{Url: "/lotus/api/v1/health", Method: "GET"},
	{Url: "/lotus", Method: "GET"},
	{Url: "/lotus/api/v1/server-monitor", Method: "GET"},
	{Url: "/lotus/api/v1/public/uploadFile", Method: "POST"},
	{Url: "/lotus/api/v1/user/pwd/set", Method: "PUT"},
	{Url: "/lotus/api/v1/sys-user", Method: "PUT"},
}
