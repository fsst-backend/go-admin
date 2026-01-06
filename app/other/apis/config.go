package apis

import (
	"go-admin/app/other/service/dto"
	"go-admin/common/models"
	"go-admin/config"

	"github.com/gin-gonic/gin"
)

type Config struct{}

// GetFrontendConfig 获取前端配置
// @Summary 获取前端配置
// @Description 获取前端所需的基础配置信息
// @Tags 配置管理
// @Accept  application/json
// @Product application/json
// @Success 200 {object} models.Response{data=dto.FrontendConfig} "前端配置信息"
// @Router /lotus/api/v1/config/frontend [get]
func (e Config) GetFrontendConfig(c *gin.Context) {
	var configData = dto.FrontendConfig{
		BaseSiteURL:   config.ExtConfig.Frontend.BaseSiteURL,
		BaseAPIURL:    config.ExtConfig.Frontend.BaseAPIURL,
		BaseH5URL:     config.ExtConfig.Frontend.BaseH5URL,
		BaseUploadURL: config.ExtConfig.Frontend.BaseUploadURL,
	}

	c.JSON(200, models.Response{Code: 200, Data: configData, Msg: "获取成功"})
}
