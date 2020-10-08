package main

import (
	"fmt"

	"github.com/yotam180/treasury/altfs"
)

func main() {
	fs := altfs.NewFS([]string{"test/"}, []string{})
	fmt.Println(fs.Exists("/b.txt"))
}
