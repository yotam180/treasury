package main

import (
	"path"

	"github.com/gin-gonic/gin"
)

func getMimetype(file string) string {
	switch path.Ext(file) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	default:
		return "text/plain"
	}
}

func serveStatic(c *gin.Context) {
	asset, err := Asset(c.Request.URL.Path[1:])
	if err == nil {
		c.Data(200, getMimetype(c.Request.URL.Path[1:]), asset) // TODO: Mime type?
		return
	}

	asset, err = Asset("index.html")
	if err == nil {
		c.Data(200, "text/html", asset)
		return
	}

	c.Data(404, "text/plain", []byte("Whoops, there was an error on our side"))
}
