package main

import (
	"github.com/yotam180/treasury/altfs"
	"github.com/yotam180/treasury/repository"
	"github.com/yotam180/treasury/server"
)

func main() {
	fs := altfs.NewFS([]string{"__test/"}, []string{"__test/"})
	bucket := repository.NewBucket(fs)

	server.Bucket = bucket
	server.Main.Run()
}
