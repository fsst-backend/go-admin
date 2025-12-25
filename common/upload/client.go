package serviceauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type UploadServiceConfig struct {
	AppKey string // admin / oms / cms
	Secret string // 与上传服务约定的密钥
}

func GenerateUploadToken(
	conf UploadServiceConfig,
	ttl time.Duration,
) (token string, expire int64) {

	expire = time.Now().Add(ttl).Unix()

	// ⚠️ 注意：不包含 method / path / body
	raw := conf.AppKey + ":" + strconv.FormatInt(expire, 10)

	mac := hmac.New(sha256.New, []byte(conf.Secret))
	mac.Write([]byte(raw))
	token = hex.EncodeToString(mac.Sum(nil))

	return
}
