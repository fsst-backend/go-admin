package handler

import (
	"go-admin/common"
	"net/http"

	"go-admin/common/global"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/mssola/user_agent"
	"gorm.io/gorm"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		return jwt.MapClaims{
			"uuid":          u.UUID,
			jwt.IdentityKey: u.UserId,
			jwt.NiceKey:     u.Username,
			"token_version": u.TokenVersion,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UUID":        claims["uuid"],
		"UserName":    claims["nice"],
	}
}

// Authenticator 获取token
// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Tags 登陆
// @Accept  application/json
// @Product application/json
// @Param account body Login  true "account"
// @Success 200 {string} string "{"code": 0, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /lotus/api/v1/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	var loginVals Login
	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	if err = c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		status = "1"

		return nil, jwt.ErrMissingLoginValues
	}
	if config.ApplicationConfig.Mode != "dev" {
		if !captcha.Verify(loginVals.UUID, loginVals.Code, true) {
			username = loginVals.Username
			msg = "验证码错误"
			status = "1"

			return nil, jwt.ErrInvalidVerificationode
		}
	}
	sysUser, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username
		// 单设备登录：递增 token_version，使其他设备上的旧 token 失效
		if err := db.Table("sys_user").Where("user_id = ?", sysUser.UserId).Update("token_version", gorm.Expr("COALESCE(token_version,0) + 1")).Error; err != nil {
			log.Warnf("increment token_version error: %s", err.Error())
		} else if err := db.Table("sys_user").Where("user_id = ?", sysUser.UserId).Select("token_version").Scan(&sysUser.TokenVersion).Error; err == nil {
			// 刷新 sysUser.TokenVersion 供 PayloadFunc 写入 JWT
		}
		return map[string]interface{}{"user": sysUser}, nil
	} else {
		msg = "登录失败"
		status = "1"
		log.Warnf("%s login failed!", loginVals.Username)
	}
	return nil, jwt.ErrFailedAuthentication
}

// LoginLogToDB Write log to database
func LoginLogToDB(c *gin.Context, status string, msg string, username string) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = common.GetClientIP(c)
	l["loginLocation"] = "" // pkg.GetLocation(common.GetClientIP(c),gaConfig.ExtConfig.AMap.Key)
	l["loginTime"] = pkg.GetCurrentTime()
	l["status"] = status
	l["remark"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["username"] = username
	l["msg"] = msg

	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", global.LoginLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}

// LogOut
// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 0, "msg": "成功退出系统" }"
// @Router /lotus/logout [post]
// @Security Bearer
func LogOut(c *gin.Context) {
	LoginLogToDB(c, "2", "退出成功", user.GetUserName(c))
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	userId := getIntFromClaims(claims, jwt.IdentityKey)
	if userId == 0 {
		return false
	}

	// 单设备登录：校验 token_version 精确匹配，若用户在其他设备新登录则当前 token 失效
	// 使用 == 判断而非 <，可正确处理 INT 有符号溢出回绕（2.1e9 → -2.1e9）的情况
	tokenVer := getIntFromClaims(claims, "token_version")
	db, err := pkg.GetOrm(c)
	if err == nil {
		var dbVer int
		if db.Table("sys_user").Where("user_id = ?", userId).Select("token_version").Scan(&dbVer).Error == nil {
			if tokenVer != dbVer {
				return false
			}
		}
	}

	// 设置用户信息到上下文
	c.Set("userId", userId)
	if uuid, ok := claims["uuid"].(string); ok {
		c.Set("uuid", uuid)
	}
	if nice, ok := claims["nice"].(string); ok {
		c.Set("userName", nice)
	}
	return true
}

func getIntFromClaims(claims jwt.MapClaims, key string) int {
	if v, ok := claims[key]; !ok || v == nil {
		return 0
	} else if n, ok := toInt(v); ok {
		return n
	}
	return 0
}

func toInt(v interface{}) (int, bool) {
	switch x := v.(type) {
	case float64:
		return int(x), true
	case int:
		return x, true
	case int64:
		return int(x), true
	}
	return 0, false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
