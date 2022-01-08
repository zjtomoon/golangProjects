package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func request(c echo.Context) error {
	req := c.Request()
	format := "<pre><strong>Request Information</strong>\n\n<code>Protocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n</code></pre>"
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}

func stream(c echo.Context) error {
	res := c.Response()
	gone := res.CloseNotify()
	res.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	res.WriteHeader(http.StatusOK)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Fprint(res, "<pre><strong>Clock Stream</strong>\n\n<code>")
	for {
		fmt.Fprintf(res, "%v\n", time.Now())
		res.Flush()
		select {
		case <-ticker.C:
		case <-gone:
			break
		}
	}
}

func main() {
	e := echo.New()
	e.GET("/request", request)
	e.GET("/stream", stream)
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}

//HTTP/2 (原本的名字是 HTTP/2.0) 是万维网使用的 HTTP 网络协议的第二个主要版本。HTTP/2 提供了更快的速度和更好的用户体验。
//
//特性
//使用二进制格式传输数据，而不是文本。使得在解析和优化扩展上更为方便。
//多路复用，所有的请求都是通过一个 TCP 连接并发完成。
//对消息头采用 HPACK 进行压缩传输，能够节省消息头占用的网络的流量。
//Server Push：服务端能够更快的把资源推送给客户端。

//生成一个自签名的 X.509 TLS 证书(HTTP/2 需要 TLS 才能运行)
//go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
