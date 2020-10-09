package server

import "github.com/gin-gonic/gin"

func listRepositories(c *gin.Context) {
	repos, err := Bucket.ListRepositories()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	results := make([]string, len(repos))
	for i, repo := range repos {
		results[i] = repo.Name
	}

	c.JSON(200, gin.H{
		"data": results,
	})
}

func init() {
	Main.GET("/api/repos", listRepositories)
}
