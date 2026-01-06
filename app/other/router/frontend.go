package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerFrontendRouter)
}

// 需认证的路由代码
func registerFrontendRouter(v1 *gin.RouterGroup) {
	//前端路由
	v1.GET("config/frontend", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
