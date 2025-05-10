package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Method + " | " + c.Request.RequestURI + " | " + c.RemoteIP())
		c.Next()
	}
}
