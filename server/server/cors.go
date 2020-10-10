package server

import (
	"github.com/gin-gonic/gin"
)

// TODO: Remove in production?
func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func init() {
	Main.Use(corsMiddleware)
}
