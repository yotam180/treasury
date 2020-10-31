package serverapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"treasury/pkg/config"
)

type serverListFileResponse struct {
	Files []string `json:"files"`
}

/*
ListFiles is used to list all files on remote
*/
func ListFiles(repo, version string) ([]string, error) {

	url, err := url.Parse(config.Config.ServerURL)
	url.Path = path.Join(url.Path, "api", "repos", repo, "releases", version)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(data))
	}

	response := serverListFileResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Files, nil
}
