package server

import "github.com/gin-gonic/gin"

func showError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
