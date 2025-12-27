package apis

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"

	serviceauth "go-admin/common/upload"
	"go-admin/config"
)

type UploadToken struct {
	api.Api
}

// GetToken 生成上传服务的认证 token
// @Summary 生成上传服务token
// @Description 根据配置文件中的 appKey 和 secret 生成上传服务的认证token
// @Tags 上传服务
// @Accept application/json
// @Product application/json
// @Success 200 {object} response.Response "{"code": 0, "message": {\"appKey\": \"admin\", \"token\": \"xxx\", \"expire\": 1234567890}}"
// @Router /lotus/api/v1/upload/token [get]
// @Security Bearer
func (e UploadToken) GetToken(c *gin.Context) {
	e.MakeContext(c)

	// 从配置文件中获取 appKey 和 secret
	uploadConfig := config.ExtConfig.Upload.GetServiceConfig()

	if uploadConfig.AppKey == "" || uploadConfig.Secret == "" {
		e.Error(500, nil, "上传服务配置未完整(缺少 appKey 或 secret)")
		return
	}

	// 生成 token,有效期 1 小时
	token, expire := serviceauth.GenerateUploadToken(uploadConfig, time.Hour)

	e.OK(gin.H{
		"appKey": uploadConfig.AppKey,
		"token":  token,
		"expire": expire,
	}, "生成成功")
}
