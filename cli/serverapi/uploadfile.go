package serverapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/user/treasury_cli/config"
)

/*
UploadFile is Used to upload a file at srcPath into dstPath on remote
*/
func UploadFile(repo, version, srcPath, dstPath string) error {
	url, err := url.Parse(config.Config.ServerURL)
	url.Path = path.Join(url.Path, "api", "repos", repo, "releases", version, "upload")
	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)
	formData, err := writer.CreateFormFile("file", dstPath)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(formData, srcFile)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url.String(), payload)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

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

	return nil
}
