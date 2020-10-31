package command

import (
	"fmt"
	"path/filepath"
	"strings"
	"treasury/pkg/serverapi"
)

/*
Download executes a downloading command command
*/
func Download(repo, version, srcPattern, dstPattern string) error {
	repoFilePaths, err := serverapi.ListFiles(repo, version)
	if err != nil {
		return err
	}

	srcPaths := []string{}
	for _, repoFilePath := range repoFilePaths {
		repoFilePath = strings.TrimPrefix(repoFilePath, "/")
		isMatch, err := filepath.Match(srcPattern, repoFilePath)
		if err != nil {
			return err
		}

		if isMatch {
			srcPaths = append(srcPaths, repoFilePath)
		}
	}

	if dstPattern == "" {
		dstPattern = filepath.FromSlash("./")
	}

	isDstPatternDirectory := strings.HasSuffix(dstPattern, filepath.FromSlash("/"))

	if len(srcPaths) == 0 {
		return fmt.Errorf("Cannot find \"%s\" on server", srcPattern)
	}

	if len(srcPaths) > 1 && !isDstPatternDirectory {
		return fmt.Errorf("Cannot download multiple files to the same location \"%s\"", dstPattern)
	}

	for _, srcPath := range srcPaths {
		dstPath := ""
		if isDstPatternDirectory {
			dstPath = filepath.Join(dstPattern, filepath.Base(filepath.FromSlash(srcPath)))
		} else {
			dstPath = dstPattern
		}

		err = serverapi.DownloadFile(repo, version, srcPath, dstPath)
		if err != nil {
			return err
		}
	}

	return nil
}
