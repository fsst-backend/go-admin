package dto

type GenTokenResponse struct {
	AppKey string `json:"appKey"`
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}
