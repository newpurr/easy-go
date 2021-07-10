package middleware

import (
	"bytes"
	"github.com/newpurr/easy-go/application"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newpurr/easy-go/pkg/logger"
)

type EpetAccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w EpetAccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func EpetAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().UnixNano() / 1e6
		c.Next()
		endTime := time.Now().UnixNano() / 1e6

		fields := logger.Fields{
			"context": map[string]interface{}{
				"request_uri":          c.Request.URL.Path,
				"request_method":       c.Request.Method,
				"refer_service_name":   "",
				"refer_request_host":   "",
				"gateway_trace":        "",
				"x_consumer_custom_id": "",
				"token":                "",
				"request_body":         c.Request.PostForm.Encode(),
				"response_body":        bodyWriter.body.String(),
				"request_time":         float64(beginTime) / 1e3,
				"response_time":        float64(endTime) / 1e3,
				"time_used":            float64(endTime-beginTime) / 1e3,
				"status_code":          "200",
				"header":               []map[string]string{},
				"service_name":         "",
				"service_addr":         "",
				"service_port":         "",
				"service_user":         "",
			},
			"datetime":   time.Now().Format("2006-01-02 15:04:05.000"),
			"level":      200,
			"level_name": "INFO",
		}
		s := "request log"
		application.Logger.WithFields(fields).Infof(c, s)
	}
}
