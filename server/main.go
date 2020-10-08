package main

import (
	"fmt"

	"github.com/yotam180/treasury/repo"

	"github.com/yotam180/treasury/altfs"
)

func main() {
	fs := altfs.NewFS([]string{"__test/", "__test copy/"}, []string{})

	r := repo.New(fs, "My Product")
	versions, err := r.ListReleases()
	if err != nil {
		fmt.Println("Error: %w", err)
		return
	}

	fmt.Println(versions)
}
