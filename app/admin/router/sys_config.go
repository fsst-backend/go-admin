package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysConfigRouter)
}

// 需认证的路由代码
func registerSysConfigRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysConfig{}
	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)
	r := v1.Group("/config").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/get", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("/configKey").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware)
	{
		r1.GET("/get", api.GetSysConfigByKEYForService)
	}

	r2 := v1.Group("/app-config")
	{
		r2.GET("", api.Get2SysApp)
	}

	r3 := v1.Group("/set-config").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware)
	{
		r3.PUT("", api.Update2Set)
		r3.GET("", api.Get2Set)
	}

}
