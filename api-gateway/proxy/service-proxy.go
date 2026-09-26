package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func UserServiceProxy() gin.HandlerFunc {

	//targetURL, err := url.Parse("http://localhost:8081")
	targetURL := os.Getenv("USER_SERVICE_URL")
	target, err := url.Parse(targetURL)

	if err != nil {
		log.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
func OrderServiceProxy() gin.HandlerFunc {

	//targetURL, err := url.Parse("http://localhost:8082")
	targetURL := os.Getenv("ORDER_SERVICE_URL")
	target, err := url.Parse(targetURL)

	if err != nil {
		log.Println(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
