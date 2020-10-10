package server

import (
	"github.com/gin-gonic/gin"
)

func listReleases(c *gin.Context) {
	repoID := c.Param("repo")
	repo, err := Bucket.OpenRepo(repoID)
	if err != nil {
		showError(c, 404, err)
		return
	}

	releases, err := repo.ListReleases()
	if err != nil {
		showError(c, 500, err)
		return
	}

	for i, j := 0, len(releases)-1; i < j; i, j = i+1, j-1 {
		releases[i], releases[j] = releases[j], releases[i]
	}

	c.JSON(200, gin.H{
		"repository": repo,
		"releases":   releases,
	})
}

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
	Main.GET("/api/repos/:repo", listReleases)
	Main.GET("/api/repos/:repo/releases/:release", listFiles)
}
