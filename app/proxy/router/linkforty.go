package router

import (
	"go-admin/app/proxy/api"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerLinkFortyProxyRouter)
}

// registerLinkFortyProxyRouter 注册 LinkForty 反向代理（/linkforty/api/v1/admin）
func registerLinkFortyProxyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	proxyAPI := api.LinkFortyProxy{}
	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	r := v1.Group("/linkforty/api/v1/admin").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.Any("/*path", proxyAPI.Proxy)
	}
}
