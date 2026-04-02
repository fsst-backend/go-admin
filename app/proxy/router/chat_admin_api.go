package router

import (
	"go-admin/app/proxy/api"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerChatAdminAPIProxyRouter)
}

// registerChatAdminAPIProxyRouter /lotus/api/v1/poplar/chat_admin_api/v1 → nwachat_im_admin_go（chat-admin-api）
func registerChatAdminAPIProxyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	proxyAPI := api.ChatAdminProxy{}
	optLogMiddleware := sdk.Runtime.GetMiddlewareKey(middleware.OperaLogToDB).(gin.HandlerFunc)

	r := v1.Group("/poplar/chat_admin_api/v1").Use(authMiddleware.MiddlewareFunc()).Use(optLogMiddleware).Use(middleware.AuthCheckRole())
	{
		r.Any("", proxyAPI.Proxy)
		r.Any("/*path", proxyAPI.Proxy)
	}
}
