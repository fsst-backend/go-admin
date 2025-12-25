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

type VioletProxy struct {
	api.Api
	TargetURL string // 目标服务地址
}

// Proxy 反向代理处理
// 将请求转发到目标服务
func (vp VioletProxy) Proxy(c *gin.Context) {
	// 从配置文件读取目标服务地址
	targetURL := vp.TargetURL
	if targetURL == "" {
		// 如果实例没有设置,尝试从全局配置读取
		targetURL = config.ExtConfig.Violet.TargetURL
	}
	if targetURL == "" {
		vp.Error(500, nil, "未配置目标服务地址,请在配置文件中设置 extend.violet.targetURL")
		return
	}

	// 解析目标URL
	remote, err := url.Parse(targetURL)
	if err != nil {
		vp.Error(500, err, "目标服务地址解析失败")
		return
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(remote)

	// 自定义Director函数,修改请求
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// 调用原始Director设置基本信息
		originalDirector(req)

		// 设置目标路径
		// 从通配符路径中获取实际路径
		path := c.Param("path")
		if path == "" {
			path = c.Request.URL.Path
		}
		req.URL.Path = path
		req.URL.RawQuery = c.Request.URL.RawQuery

		// 移除origin头,避免CORS问题
		req.Header.Del("Origin")

		// 设置真实客户端IP
		if clientIP := c.ClientIP(); clientIP != "" {
			req.Header.Set("X-Real-IP", clientIP)
			req.Header.Set("X-Forwarded-For", clientIP)
		}

		// 添加自定义头:域ID
		domainID := config.ExtConfig.Violet.DomainID
		if domainID != "" {
			req.Header.Set(config.HeaderPoplarDomainID, domainID)
		}

		// 从JWT Claims中获取用户信息
		if data, exists := c.Get(jwtauth.JwtPayloadKey); exists {
			if claims, ok := data.(jwtauth.MapClaims); ok {
				// 获取用户UUID
				if uuid, ok := claims["uuid"].(string); ok && uuid != "" {
					req.Header.Set(config.HeaderPoplarUserUUID, uuid)
				}

				// 获取用户ID (identity key)
				if userID := claims[jwtauth.IdentityKey]; userID != nil {
					req.Header.Set(config.HeaderPoplarUserID, fmt.Sprintf("%v", userID))
				}

				// 获取用户名 (nice key)
				if userName, ok := claims[jwtauth.NiceKey].(string); ok && userName != "" {
					req.Header.Set("X-User-Name", userName)
				}
			}
		} else {
			// 备用方案:从context中获取(Authorizator设置的)
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

	// 自定义错误处理
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		vp.Logger.Errorf("Proxy error: %v", err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}

	// 修改响应
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 可以在这里修改响应头
		resp.Header.Set("X-Proxy-By", "go-admin")
		return nil
	}

	// 执行代理请求
	proxy.ServeHTTP(c.Writer, c.Request)
}
