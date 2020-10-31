package main

import (
	"github.com/user/treasury/altfs"
	"github.com/user/treasury/repository"
	"github.com/user/treasury/server"
)

func main() {
	config := getConfig()
	fs := altfs.NewFS(config.Reads, config.Writes)
	bucket := repository.NewBucket(fs)

	server.Bucket = bucket

	server.Main.NoRoute(serveStatic)
	server.Main.Run()
}
