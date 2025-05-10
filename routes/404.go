package routes

import (
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.HTML(404, "404.html", gin.H{"title": "404 Not Found", "content": "404 Not Found !"})
}
