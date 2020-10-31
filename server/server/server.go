package server

import (
	"github.com/gin-gonic/gin"
	"github.com/user/treasury/repository"
)

/*
Main is the main gin server instance
*/
var Main = gin.Default()

/*
Bucket is the bucket to serve repositories and versions from.
*/
var Bucket *repository.Bucket = nil

func listRepositories(c *gin.Context) {
	repos, err := Bucket.ListRepositories()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": repos,
	})
}

func init() {
	Main.GET("/api/repos", listRepositories)
}
