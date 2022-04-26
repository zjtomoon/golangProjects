package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	// 参考: https://www.cnblogs.com/wang_yb/p/14593606.html
	GinHttps(false) // 这里false 表示 http 服务，非 https
}

func GinHttps(isHttps bool) error {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "test for 【%s】", "https")
	})

	if isHttps {
		r.Use(TlsHandler(8000))

		return r.RunTLS(":"+strconv.Itoa(8000), "/path/to/test.pem", "/path/to/test.key")
	}

	return r.Run(":" + strconv.Itoa(8000))
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
