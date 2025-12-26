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
//	    domainID: default-domain
//	  upload:
//	    appKey: admin
//	    secret: your-secret-key
//
// 使用方法: config.ExtConfig......即可!!
type Extend struct {
	AMap   AMap   // 这里配置对应配置文件的结构即可
	Violet Violet // Violet 反向代理配置
	Upload Upload // Upload 上传服务配置
}

type AMap struct {
	Key string `yaml:"key" json:"key"` // 高德地图API密钥
}

// Violet 反向代理配置
type Violet struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 目标服务地址
	DomainID  int64  `yaml:"domainID" json:"domainID"`   // 域ID
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
