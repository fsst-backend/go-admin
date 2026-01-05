package router

import (
	"go-admin/app/other/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	// routerCheckRole = append(routerCheckRole, registerFileRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerDocRouter)
}

// 需认证的路由代码
func registerFileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	var api = apis.File{}
	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	r := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware)
	{
		r.POST("/public/uploadFile", api.UploadFile)
	}
}

func registerDocRouter(v1 *gin.RouterGroup) {
	r := v1.Group("")
	{
		r.Static("/docs", "/app/docs")
	}
}
