package config

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
//
// 使用方法: config.ExtConfig......即可!!
type Extend struct {
	AMap   AMap   // 这里配置对应配置文件的结构即可
	Violet Violet // Violet 反向代理配置
}

type AMap struct {
	Key string
}

// Violet 反向代理配置
type Violet struct {
	TargetURL string `yaml:"targetURL" json:"targetURL"` // 目标服务地址
	DomainID  string `yaml:"domainID" json:"domainID"`   // 域ID
}
