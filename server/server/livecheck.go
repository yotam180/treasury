package server

import "github.com/gin-gonic/gin"

func liveCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "alive",
	})
}

func init() {
	Main.GET("/api/live", liveCheck)
}
