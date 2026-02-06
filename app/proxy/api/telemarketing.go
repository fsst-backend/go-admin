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

type TelemarketingProxy struct {
	api.Api
	TargetURL string // 目标服务地址，可覆盖配置
}

// Proxy 反向代理：将 /Telemarketing/open/v1 转发到配置的目标（默认端口 9999）
func (tp TelemarketingProxy) Proxy(c *gin.Context) {
	targetURL := tp.TargetURL
	if targetURL == "" {
		targetURL = config.ExtConfig.Telemarketing.TargetURL
	}
	if targetURL == "" {
		targetURL = "http://localhost:9999"
	}

	remote, err := url.Parse(targetURL)
	if err != nil {
		tp.Error(500, err, "目标服务地址解析失败")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		path := c.Request.URL.Path
		req.URL.Path = path
		req.URL.RawQuery = c.Request.URL.RawQuery
		req.Header.Del("Origin")
		if clientIP := c.ClientIP(); clientIP != "" {
			req.Header.Set("X-Real-IP", clientIP)
			req.Header.Set("X-Forwarded-For", clientIP)
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
		tp.Logger.Errorf("Telemarketing proxy error: %v", err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("X-Proxy-By", "go-admin")
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
