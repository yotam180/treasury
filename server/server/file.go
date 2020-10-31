package server

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/user/treasury/repository"
)

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

func uploadFile(c *gin.Context) {
	repoID := c.Param("repo")
	releaseID := c.Param("release")

	var err error

	var repo *repository.Repo
	if Bucket.Exists(repoID) {
		repo, err = Bucket.OpenRepo(repoID)
		if err != nil {
			showError(c, 404, err)
			return
		}
	} else {
		repo, err = Bucket.CreateRepo(repoID)
		if err != nil {
			showError(c, 404, err)
			return
		}
	}

	var release *repository.Release
	if Bucket.Exists(path.Join(repoID, releaseID)) {
		release, err = repo.OpenRelease(releaseID)
		if err != nil {
			showError(c, 404, err)
			return
		}
	} else {
		release, err = repo.CreateRelease(releaseID)
		if err != nil {
			showError(c, 404, err)
			return
		}
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		showError(c, 404, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		showError(c, 404, err)
		return
	}

	err = release.AddFile(fileHeader.Filename, file)
	if err != nil {
		showError(c, 404, err)
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func init() {
	Main.GET("/api/repos/:repo/releases/:release/files/*file", downloadFile)
	Main.POST("/api/repos/:repo/releases/:release/upload", uploadFile)
}
