package middleware

import (
	"MarvelousBlog-Backend/config"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求处理前时间
		start := time.Now()
		//处理请求
		c.Next()
		//处理请求后时间
		end := time.Now()

		//计算实际处理的时间
		latencyTime := end.Sub(start)
		//获取请求相关信息，如请求方法，IP地址、路由、状态码等信息
		requestMethod := c.Request.Method
		requestUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		config.Log.Infof("| %s |Status : %3d |ClientIP : %13v |Latency : %15s |URI %s |",
			requestMethod, statusCode, clientIP, latencyTime, requestUri)
	}
}
