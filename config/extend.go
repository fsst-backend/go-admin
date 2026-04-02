package config

import serviceauth "go-admin/common/upload"

var ExtConfig Extend

const (
	HeaderPoplarUserUUID = "X-Poplar-User-Id"
	HeaderPoplarDomainID = "X-Poplar-Domain-Id"
	HeaderPoplarUserID   = "X-Poplar-UserId"
)

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//	  violet:
//	    targetURL: http://localhost:8080
//	    domainID: 1  # Violet / LinkForty 代理共用，透传 X-Poplar-Domain-Id
//	  linkforty:
//	    targetURL: http://localhost:8081
//	  telemarketing:
//	    targetURL: http://localhost:9999
//	  sms:
//	    targetURL: http://localhost:xxxx
//	  chatAdminAPI:
//	    targetURL: http://chat-admin-api:9999
//	  upload:
//	    appKey: admin
//	    secret: your-secret-key
//
// 使用方法: config.ExtConfig......即可!!
type Extend struct {
	AMap          AMap           // 这里配置对应配置文件的结构即可
	Violet        Violet         // Violet 反向代理配置
	LinkForty     LinkForty      // LinkForty 反向代理配置（/linkforty/api/v1/admin）
	Telemarketing Telemarketing  // Telemarketing 反向代理配置
	SMS           SMS            // SMS 反向代理配置（/poplar/sms/v1）
	ChatAdminAPI  ChatAdminAPI   // 聊天管理后台 API 代理（/poplar/chat_admin_api/v1，路径原样转发）
	Upload        Upload         // Upload 上传服务配置
	Frontend      FrontendConfig // 前端配置
}

type AMap struct {
	Key string `yaml:"key" json:"key"` // 高德地图API密钥
}

type FrontendConfig struct {
	BaseSiteURL   string `yaml:"baseSiteURL" json:"baseSiteURL"`     // 基础站点URL
	BaseAPIURL    string `yaml:"baseAPIURL" json:"baseAPIURL"`       // 基础API URL
	BaseH5URL     string `yaml:"baseH5URL" json:"baseH5URL"`         // H5基础URL
	BaseUploadURL string `yaml:"baseUploadURL" json:"baseUploadURL"` // 上传基础URL
}

// Violet 反向代理配置（domainID 与 LinkForty 代理共用）
type Violet struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 目标服务地址
	DomainID  int64  `yaml:"domainID" json:"domainID"`   // 域ID，LinkForty 代理同样透传此值
}

// LinkForty 反向代理配置（/linkforty/api/v1/admin）
type LinkForty struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // LinkForty 远程服务根地址
}

// Telemarketing 反向代理配置（端口 9999）
type Telemarketing struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 目标服务地址，默认 http://localhost:9999
}

// SMS 反向代理配置（/poplar/sms/v1）
type SMS struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 目标服务地址
}

// ChatAdminAPI 反向代理（路径原样转发至 targetURL，镜像 nwachat_im_admin_go）
type ChatAdminAPI struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 如 http://chat-admin-api:9999
}

// Upload 上传服务配置
type Upload struct {
	AppKey string `yaml:"appKey" json:"appKey"` // 应用标识
	Secret string `yaml:"secret" json:"secret"` // 密钥
}

// GetServiceConfig 获取上传服务配置
func (u *Upload) GetServiceConfig() serviceauth.UploadServiceConfig {
	return serviceauth.UploadServiceConfig{
		AppKey: u.AppKey,
		Secret: u.Secret,
	}
}
