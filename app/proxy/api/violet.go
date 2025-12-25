package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type VioletProxy struct {
	api.Api
}

// ServerInfo 获取系统信息
func (vp VioletProxy) Proxy(c *gin.Context) {
	// fullPath := path.Join(destPrefix, c.Param("path"))
	// c.Request.URL.Path = fullPath

	// c.Request.Header.Del("origin") // 去掉origin头
	// if c.GetBool(PoplarProxyKey) {
	// 	c.Request.Header.Set(HeaderPoplarUserID, c.GetString(uuidKeyName))
	// 	c.Request.Header.Set(HeaderPoplarDomainID, c.GetString(currentDomainKeyName))
	// }
	// proxy.ServeHTTP(c.Writer, c.Request)
}
