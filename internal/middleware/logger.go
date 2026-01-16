package middleware
import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func(c *gin.Context){
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		status := c.Writer.Status()
		clientIp := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf(
			"🕙 [%d] %s | %13v | %s %s\n",
			status,
			clientIp,
			latency,
			method,
			path,
		)
	}
}