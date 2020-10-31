package server

import (
	"github.com/gin-gonic/gin"
)

func listFiles(c *gin.Context) {
	repoID := c.Param("repo")
	releaseID := c.Param("release")

	repo, err := Bucket.OpenRepo(repoID)
	if err != nil {
		showError(c, 404, err)
		return
	}

	release, err := repo.OpenRelease(releaseID)
	if err != nil {
		showError(c, 404, err)
		return
	}

	files := release.ListFiles()
	metadata := release.GetMetadata()

	c.JSON(200, gin.H{
		"data":     release,
		"files":    files,
		"metadata": metadata,
	})
}

func init() {
	Main.GET("/api/repos/:repo/releases/:release", listFiles)
}
