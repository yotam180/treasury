package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/user/treasury_cli/command"
	"github.com/user/treasury_cli/config"

	flag "github.com/spf13/pflag"
)

type cliArguments struct {
	isUpload   bool
	isDownload bool
	repo       string
	version    string
	srcPattern string
	dstPattern string
}

var cliArgs cliArguments

func initCLIArgs() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags] src [dst]\n", os.Args[0])
		fmt.Println("\nFlags")
		flag.PrintDefaults()
		fmt.Println("\nSpecials\n  {*} - Wildcard\n  {r}, {repo} - inplacing repo\n  {v}, {version} - inplacing version")
	}

	flag.BoolVarP(&cliArgs.isUpload, "upload", "u", false, "Set to upload files")
	flag.BoolVarP(&cliArgs.isDownload, "download", "d", false, "Set to download files")
	flag.StringVarP(&cliArgs.repo, "repo", "r", "", "Repo to upload/download to/from")
	flag.StringVarP(&cliArgs.version, "version", "v", "", "Version to upload/download to/from")
	flag.Parse()

	args := flag.Args()
	if len(args) == 1 {
		cliArgs.srcPattern = args[0]
	} else if len(args) == 2 {
		cliArgs.srcPattern, cliArgs.dstPattern = args[0], args[1]
	} else {
		fmt.Println("Wrong number of arguments")
		flag.Usage()
		os.Exit(1)
	}

	if cliArgs.isUpload == cliArgs.isDownload {
		fmt.Println("Please choose upload/download exclusively")
		flag.Usage()
		os.Exit(1)
	}

	repoPattern := regexp.MustCompile(config.Config.RepoPattern)
	if !repoPattern.MatchString(cliArgs.repo) {
		fmt.Println("Wrong repo format")
		flag.Usage()
		os.Exit(1)
	}

	versionPattern := regexp.MustCompile(config.Config.VersionPattern)
	if !versionPattern.MatchString(cliArgs.version) {
		fmt.Println("Wrong version format")
		flag.Usage()
		os.Exit(1)
	}

	// TODO - add unallowed glob/filepath.match metacharacters on src and dst
}

func applySpecials(str string) string {
	str = strings.ReplaceAll(str, "{r}", cliArgs.repo)
	str = strings.ReplaceAll(str, "{repo}", cliArgs.repo)
	str = strings.ReplaceAll(str, "{v}", cliArgs.version)
	str = strings.ReplaceAll(str, "{version}", cliArgs.version)
	return strings.ReplaceAll(str, "{*}", "*")
}

func main() {
	initCLIArgs()

	var err error

	if cliArgs.isUpload {
		err = command.Upload(cliArgs.repo, cliArgs.version, applySpecials(cliArgs.srcPattern), applySpecials(cliArgs.dstPattern))
	} else if cliArgs.isDownload {
		err = command.Download(cliArgs.repo, cliArgs.version, applySpecials(cliArgs.srcPattern), applySpecials(cliArgs.dstPattern))
	}

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
