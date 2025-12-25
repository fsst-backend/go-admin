package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDictRouter)
}

func registerDictRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	dictApi := apis.SysDictType{}
	dataApi := apis.SysDictData{}

	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	dicts := v1.Group("/dict").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{

		dicts.GET("/data", dataApi.GetPage)
		dicts.GET("/data/get", dataApi.Get)
		dicts.POST("/data", dataApi.Insert)
		dicts.PUT("/data", dataApi.Update)
		dicts.DELETE("/data", dataApi.Delete)

		dicts.GET("/type-option-select", dictApi.GetAll)
		dicts.GET("/type", dictApi.GetPage)
		dicts.GET("/type/get", dictApi.Get)
		dicts.POST("/type", dictApi.Insert)
		dicts.PUT("/type", dictApi.Update)
		dicts.DELETE("/type", dictApi.Delete)
	}
	opSelect := v1.Group("/dict-data").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		opSelect.GET("/option-select", dataApi.GetAll)
	}
}
