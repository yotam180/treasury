package server

import "github.com/gin-gonic/gin"

func repoInfo(c *gin.Context) {
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

	metadata := repo.GetMetadata()

	c.JSON(200, gin.H{
		"repository": repo,
		"metadata":   metadata,
		"releases":   releases,
	})
}

func init() {
	Main.GET("/api/repos/:repo", repoInfo)
}
