package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2/util"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		var res, casbinExclude bool
		var err error

		// 获取userId和rolekey
		userIdInterface := v[jwtauth.IdentityKey]

		// 检查是否在排除列表中
		for _, i := range CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}

		// 使用user_{userId}进行权限检查
		userSubject := fmt.Sprintf("user_%v", userIdInterface)
		res, err = e.Enforce(userSubject, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Errorf("AuthCheckRole error:%s method:%s path:%s user:%s", err, c.Request.Method, c.Request.URL.Path, userSubject)
			response.Error(c, 500, err, "")
			return
		}

		if res {
			log.Infof("isTrue: %v user: %s method: %s path: %s", res, userSubject, c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v user: %s method: %s path: %s message: %s", res, userSubject, c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}

	}
}
