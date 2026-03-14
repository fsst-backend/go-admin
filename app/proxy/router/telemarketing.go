package router

import (
	"go-admin/app/proxy/api"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTelemarketingProxyRouter)
}

// registerTelemarketingProxyRouter 注册 Telemarketing 反向代理路由（目标端口 9999）
func registerTelemarketingProxyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	proxyAPI := api.TelemarketingProxy{}
	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	// 需要认证和权限验证的代理路由：/Telemarketing/open/v1 -> 目标服务（默认 http://localhost:9999）
	r := v1.Group("/Telemarketing/open/v1").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.Any("/*path", proxyAPI.Proxy)
	}
}
