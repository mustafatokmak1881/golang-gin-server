package routes

import (
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.HTML(200, "user.html", gin.H{"title": "User Page", "content": "User Content"})
}
