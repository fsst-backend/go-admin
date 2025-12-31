package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}

	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	r := v1.Group("/sys-user").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetPage)
		r.GET("/get", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/role", api.SetUserRole) // 新增：设置用户角色
	}

	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		user.PUT("/pwd/set", api.UpdatePwd)
		user.PUT("/pwd/reset", api.ResetPwd)
		user.PUT("/status", api.UpdateStatus)
	}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware)
	{
		v1auth.GET("/getinfo", api.GetInfo)
		user.GET("/profile", api.GetProfile)
		user.POST("/avatar", api.InsetAvatar)
	}
}
