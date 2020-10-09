package server

import "github.com/gin-gonic/gin"

func downloadFile(c *gin.Context) {
	repoID := c.Param("repo")
	releaseID := c.Param("release")
	fileID := c.Param("file")

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

	file, err := release.GetFile(fileID)
	if err != nil {
		showError(c, 404, err)
		return
	}
	fileName := file.Name()
	file.Close()

	c.FileAttachment(fileName, fileID)
}

func init() {
	Main.GET("/api/repos/:repo/releases/:release/files/:file", downloadFile)
}
