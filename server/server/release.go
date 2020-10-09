package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func listReleases(c *gin.Context) {
	repoID := c.Param("name")
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "Malformed 'name' parameter: " + err.Error(),
	// 	})
	// 	return
	// }

	fmt.Println("repoID: ", repoID)

	repo, err := Bucket.OpenRepo(repoID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("repo: ", repo)

	releases, err := repo.ListReleases()
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": releases,
	})
}

func init() {
	Main.GET("/api/repos/:name/releases", listReleases)
}
