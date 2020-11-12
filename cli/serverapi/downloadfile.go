package serverapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/user/treasury_cli/config"
)

/*
DownloadFile is used to download file at srcPath on remote into dstPath
*/
func DownloadFile(repo, version, srcPath, dstPath string) error {
	url, err := url.Parse(config.Config.ServerURL)
	url.Path = path.Join(url.Path, "api", "repos", repo, "releases", version, "files", srcPath)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(data))
	}

	err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, res.Body)

	return err
}
