package router

import (
	"go-admin/app/proxy/api"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerVioletProxyRouter)
}

// registerVioletProxyRouter 注册反向代理路由
// 需要认证的路由,将请求转发到目标服务
func registerVioletProxyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	proxyAPI := api.VioletProxy{
		// 可以在这里设置目标服务地址
		// TargetURL: "http://target-service:8080",
	}

	// 需要认证和权限验证的代理路由
	r := v1.Group("/poplar/violet").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 通配符路由,捕获所有路径
		// 注意: /*path 只捕获 /poplar/violet 之后的部分
		// 如果需要保留完整路径,应在 Proxy 中使用 c.Request.URL.Path
		r.Any("/*path", proxyAPI.Proxy)
	}
}
