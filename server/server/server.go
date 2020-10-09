package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yotam180/treasury/repository"
)

/*
Main is the main gin server instance
*/
var Main = gin.Default()

/*
Bucket is the bucket to serve repositories and versions from.
*/
var Bucket *repository.Bucket = nil
