package router

import (
	"go-admin/app/proxy/api"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSmsProxyRouter)
}

// registerSmsProxyRouter 注册 SMS 反向代理路由（/poplar/sms/v1）
func registerSmsProxyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	proxyAPI := api.SmsProxy{}

	r := v1.Group("/poplar/sms/v1").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.Any("/*path", proxyAPI.Proxy)
	}

	export := v1.Group("/poplar/sms/noauth/v1")
	{
		export.Any("/*path", proxyAPI.Proxy)
	}
}
