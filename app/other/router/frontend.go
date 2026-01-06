package router

import (
	"go-admin/app/other/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerFrontendRouter)
}

// 需认证的路由代码
func registerFrontendRouter(v1 *gin.RouterGroup) {
	//前端配置路由
	var configApi = apis.Config{}
	v1.GET("config/frontend", configApi.GetFrontendConfig)

}
