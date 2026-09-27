package proxy

import (
	"log"
	"net/http"
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

	return func(c *gin.Context) {
		targetURL, err := url.Parse("http://localhost:8082")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid order service URL"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			c.JSON(http.StatusBadGateway, gin.H{"error": "order service unavailable"})
		}
		// Forward the request to Order Service.
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
