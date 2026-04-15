package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/config"
)

type ChatAdminProxy struct {
	api.Api
	TargetURL string
}

// Proxy 反向代理：路径与原样 query 转发至 chat-admin-api（nwachat_im_admin_go），不改写前缀
func (p ChatAdminProxy) Proxy(c *gin.Context) {
	targetURL := p.TargetURL
	if targetURL == "" {
		targetURL = config.ExtConfig.ChatAdminAPI.TargetURL
	}
	if targetURL == "" {
		p.Error(500, nil, "未配置目标服务地址,请在配置文件中设置 extend.chatAdminAPI.targetURL")
		return
	}

	remote, err := url.Parse(targetURL)
	if err != nil {
		p.Error(500, err, "目标服务地址解析失败")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = c.Request.URL.Path
		req.URL.RawQuery = c.Request.URL.RawQuery
		req.Header.Del("Origin")
		if clientIP := c.ClientIP(); clientIP != "" {
			req.Header.Set("X-Real-IP", clientIP)
			req.Header.Set("X-Forwarded-For", clientIP)
		}
		domainID := config.ExtConfig.Violet.DomainID
		if domainID != 0 {
			req.Header.Set(config.HeaderPoplarDomainID, fmt.Sprintf("%d", domainID))
		}
		if data, exists := c.Get(jwtauth.JwtPayloadKey); exists {
			if claims, ok := data.(jwtauth.MapClaims); ok {
				if uuid, ok := claims["uuid"].(string); ok && uuid != "" {
					req.Header.Set(config.HeaderPoplarUserUUID, uuid)
				}
				if userID := claims[jwtauth.IdentityKey]; userID != nil {
					req.Header.Set(config.HeaderPoplarUserID, fmt.Sprintf("%v", userID))
				}
				if userName, ok := claims[jwtauth.NiceKey].(string); ok && userName != "" {
					req.Header.Set("X-User-Name", userName)
				}
			}
		} else {
			if userUUID, exists := c.Get("uuid"); exists {
				if uuid, ok := userUUID.(string); ok && uuid != "" {
					req.Header.Set(config.HeaderPoplarUserUUID, uuid)
				}
			}
			if userID, exists := c.Get("userId"); exists {
				if uid, ok := userID.(int); ok && uid > 0 {
					req.Header.Set(config.HeaderPoplarUserID, fmt.Sprintf("%d", uid))
				}
			}
			if userName, exists := c.Get("userName"); exists {
				if name, ok := userName.(string); ok && name != "" {
					req.Header.Set("X-User-Name", name)
				}
			}
		}
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		p.Logger.Errorf("ChatAdmin API proxy error: %v", err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("X-Proxy-By", "go-admin")
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
