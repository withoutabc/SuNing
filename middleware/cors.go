package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	// 允许所有header
	var headerKeys []string
	for k, _ := range c.Request.Header {
		headerKeys = append(headerKeys, k)
	}
	headerStr := strings.Join(headerKeys, ", ")
	if headerStr != "" {
		headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	} else {
		headerStr = "access-control-allow-origin, access-control-allow-headers"
	}

	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有的域,也可以指定某几个特定的域
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		// header的类型
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma") //允许跨域设置可以返回其他子段
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")                                                                                                           // 跨域关键设置 让浏览器可以解析
		c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                                                                                                                     // 缓存请求信息 单位为秒
		c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                                                                                                                            // 跨域请求是否需要带cookie信息 默认设置为true
	}

	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
