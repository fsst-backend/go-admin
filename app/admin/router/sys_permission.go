package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysPermissionRouter)
}

// 需认证的路由代码
func registerSysPermissionRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysPermission{}

	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)
	r := v1.Group("/sys-permission").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/get", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}
}
