package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

// GraphQL 请求结构
type GraphQLRequest struct {
	OperationName string `json:"operationName"`
}

func GqlLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" && c.Request.URL.Path == "/graphql" {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				var gqlRequest GraphQLRequest
				if json.Unmarshal(bodyBytes, &gqlRequest) == nil && gqlRequest.OperationName != "" {
					c.Set("operationName", gqlRequest.OperationName) // 存入 Gin 上下文
				}
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}
		c.Next()
	}
}

func GqlLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 设置日志路径
		path := fmt.Sprintf("\"%s\"", param.Path)
		if param.Path == "/graphql" {
			// 从 Gin 上下文获取 operationName
			operationName := "unknown"
			if val, exists := param.Keys["operationName"]; exists {
				operationName = val.(string)
			}
			path = fmt.Sprintf("\"/graphql\" | %s", operationName)
		}

		// 格式化日志
		return fmt.Sprintf("[GIN] %s | \033[32m%3d\033[0m | %15s | %15s | %-7s %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"), // 时间
			param.StatusCode,    // 状态码 (对齐3位)
			param.Latency.Abs(), // 响应时间
			param.ClientIP,      // 客户端IP
			param.Method,        // 请求方法
			path,                // 请求路径
		)
	})
}
