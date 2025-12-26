package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
)

type System struct {
	api.Api
}

// GenerateCaptchaHandler 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登殆
// @Success 200 {object} response.Response{data=string,id=string,msg=string} "{\"code\": 0, \"data\": [...]}"
// @Router /lotus/api/v1/captcha [get]
func (e System) GenerateCaptchaHandler(c *gin.Context) {
	if err := e.MakeContext(c).Errors; err != nil {
		e.Error(500, err, "服务初始化失败！")
		return
	}
	id, b64s, answer, err := captcha.DriverDigitFunc()
	if err != nil {
		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(500, err, "验证码获取失败")
		return
	}
	e.Logger.Infof("Captcha generated - id: %s, answer: %s", id, answer)
	e.OK(gin.H{
		"id":   id,
		"data": b64s,
	}, "success")
}
