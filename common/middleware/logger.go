package middleware

import (
	"bytes"
	"encoding/json"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"

	"go-admin/common/global"
)

// responseBodyWriter 包装 ResponseWriter 以捕获响应体（用于代理路由的业务错误判断）
// 仅缓冲前 8KB 用于解析 code 字段，避免大响应占用过多内存
const maxBodyCaptureSize = 8192

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	if w.body.Len() < maxBodyCaptureSize {
		remain := maxBodyCaptureSize - w.body.Len()
		if len(b) <= remain {
			w.body.Write(b)
		} else {
			w.body.Write(b[:remain])
		}
	}
	return w.ResponseWriter.Write(b)
}

// parseProxyResponseCode 解析代理响应 JSON 中的 code 字段，code==0 为成功
func parseProxyResponseCode(body []byte) (code int, ok bool) {
	var m struct {
		Code float64 `json:"code"`
	}
	if err := json.Unmarshal(body, &m); err != nil {
		return 0, false
	}
	return int(m.Code), true
}

// LoggerToFile 日志记录到文件
func SaveOperaLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			rb, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.Warnf("copy body error, %s", err.Error())
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		// Violet 代理路由返回 HTTP 200 但 body 中 code!=0 表示业务错误，需捕获响应体判断
		url := c.Request.RequestURI
		var bodyWriter *responseBodyWriter
		if strings.Contains(url, "/poplar/violet/") {
			bodyWriter = &responseBodyWriter{ResponseWriter: c.Writer, body: bytes.NewBuffer(nil)}
			c.Writer = bodyWriter
		}

		c.Next()
		if strings.Contains(url, "logout") ||
			strings.Contains(url, "login") {
			return
		}


		


		// 结束时间
		endTime := time.Now()
		if c.Request.Method == http.MethodOptions {
			return
		}

		rt, bl := c.Get("result")
		var result = ""
		if bl {
			rb, err := json.Marshal(rt)
			if err != nil {
				log.Warnf("json Marshal result error, %s", err.Error())
			} else {
				result = string(rb)
			}
		}

		st, bl := c.Get("status")
		var statusBus = 0
		if bl {
			statusBus = st.(int)
		}
		// 当 handler 未设置 status 时（如代理路由），使用 HTTP 状态码判断成功/失败
		if statusBus == 0 {
			statusBus = c.Writer.Status()
		}
		// 代理路由：body 中 code!=0 表示业务错误（如 Violet 返回 {"code":10007,"message":"配置ID无效"}）
		if bodyWriter != nil && statusBus == http.StatusOK {
			if respCode, ok := parseProxyResponseCode(bodyWriter.body.Bytes()); ok && respCode != 0 {
				statusBus = http.StatusInternalServerError
			}
		}

		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := common.GetClientIP(c)
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		log.WithFields(logData).Info()
		defer func() {
			log.Fields(map[string]interface{}{})
		}()
		if c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB && statusCode != 404 {
			SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime, body, result, statusBus)
		}
	}
}

// SetDBOperLog 写入操作日志表 fixme 该方法后续即将弃用
func SetDBOperLog(c *gin.Context, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration, body string, result string, status int) {

	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})
	l["_fullPath"] = c.FullPath()
	l["operUrl"] = reqUri
	l["operIp"] = clientIP
	l["operLocation"] = "" // pkg.GetLocation(clientIP, gaConfig.ExtConfig.AMap.Key)
	l["operName"] = user.GetUserName(c)
	l["requestMethod"] = reqMethod
	l["operParam"] = body
	l["operTime"] = time.Now()
	l["jsonResult"] = result
	l["latencyTime"] = latencyTime.String()
	l["statusCode"] = statusCode
	l["userAgent"] = c.Request.UserAgent()
	l["createBy"] = user.GetUserId(c)
	l["updateBy"] = user.GetUserId(c)
	if status == 0 || status == http.StatusOK {
		l["status"] = dto.OperaStatusEnabel
	} else {
		l["status"] = dto.OperaStatusDisable
	}
	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", global.OperateLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		// 日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}
