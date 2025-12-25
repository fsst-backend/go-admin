package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

// 需认证的路由代码
func registerSysRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysRole{}

	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)
	r := v1.Group("/role").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/get", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}
	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.PUT("/role-status", api.Update2Status)
		r1.PUT("/roledatascope", api.Update2DataScope)
	}
}
