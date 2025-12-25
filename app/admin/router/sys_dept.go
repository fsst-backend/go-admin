package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysDept{}

	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)
	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/get", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware)
	{
		r1.GET("/deptTree", api.Get2Tree)
	}
}
