package command

import (
	"fmt"
	"path/filepath"
	"strings"
	"treasury/pkg/serverapi"
)

/*
Upload executes an uploading command
*/
func Upload(repo, version, srcPattern, dstPattern string) error {
	srcPaths, err := filepath.Glob(srcPattern)
	if err != nil {
		return err
	}

	if dstPattern == "" {
		dstPattern = "./"
	}

	isDstPatternDirectory := strings.HasSuffix(dstPattern, "/")

	if len(srcPaths) == 0 {
		return fmt.Errorf("Cannot find \"%s\" on local", srcPattern)
	}

	if len(srcPaths) > 1 && !isDstPatternDirectory {
		return fmt.Errorf("Cannot upload multiple files to the same location \"%s\"", dstPattern)
	}

	for _, srcPath := range srcPaths {
		dstPath := ""
		if isDstPatternDirectory {
			dstPath = filepath.ToSlash(filepath.Join(dstPattern, filepath.Base(srcPath)))
		} else {
			dstPath = dstPattern
		}

		err = serverapi.UploadFile(repo, version, srcPath, dstPath)
		if err != nil {
			return err
		}
	}

	return nil
}
