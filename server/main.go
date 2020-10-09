package main

import (
	"fmt"
	"os"

	"github.com/yotam180/treasury/repository"

	"github.com/yotam180/treasury/altfs"
)

func main() {
	fs := altfs.NewFS([]string{"__test/", "__test copy/"}, []string{"__test/"})

	r := repository.New(fs, "My Product")
	// versions, err := r.ListReleases()
	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// 	return
	// }

	release, err := r.CreateRelease("0.3.0")
	if err != nil {
		fmt.Println("Could not create release:", err.Error())
		return
	}

	fmt.Println(release.GetMetadata())
	release.SetMetadata(map[string]interface{}{
		"author": "Yotam S",
	})
	fmt.Println(release.GetMetadata())

	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = release.AddFile("main_exe.go", f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(release.Version)
}
