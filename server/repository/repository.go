package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/yotam180/treasury/altfs"
)

/*
Repo is a collection of releases for a project
*/
type Repo struct {
	altfs.FileSystem

	Name string
}

/*
Release is a specific release of a
*/
type Release struct {
	Repo *Repo

	Version string
}

/*
New creates a new repository object. It does not create the repository
*/
func New(fileSystem altfs.FileSystem, name string) *Repo {
	return &Repo{fileSystem, name}
}

/*
ListReleases returns an array of all releases in the repository.
*/
func (repo *Repo) ListReleases() ([]Release, error) {
	subDirs, err := repo.ListDir(repo.Name)
	if err != nil {
		return nil, fmt.Errorf("can't list versions: %w", err)
	}

	releases := make([]Release, 0)

	for _, dir := range subDirs {
		if dir.IsDir() {
			releases = append(releases, Release{repo, dir.Name()})
		}
	}

	return releases, nil
}

/*
CreateRelease opens a new release folder in a repo and returns the release object
*/
func (repo *Repo) CreateRelease(version string) (Release, error) {
	// TODO: Verify that this doesn't enter some other release files folder or anything, and that it is acceptable as a release.
	err := repo.Mkdir(path.Join(repo.Name, version))
	if err != nil {
		return Release{}, err
	}

	return Release{repo, version}, nil
}

/*
GetMetadata returns the metadata object for the release
*/
func (release Release) GetMetadata() map[string]interface{} {
	f, err := release.Repo.Open(path.Join(release.Path(), "metadata.json"))
	if err != nil {
		return map[string]interface{}{}
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return map[string]interface{}{}
	}

	var data = make(map[string]interface{}, 0)
	err = json.Unmarshal(content, &data)

	return data
}

/*
SetMetadata adds some metadata keys to the repository
*/
func (release Release) SetMetadata(newKeys map[string]interface{}) error {
	metadata := release.GetMetadata()

	for key, value := range newKeys {
		metadata[key] = value
	}

	f, err := release.Repo.Create(path.Join(release.Path(), "metadata.json"))
	if err != nil {
		return fmt.Errorf("can't open metadata file for writing: %w", err)
	}
	defer f.Close()

	marshalled, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("can't encode metadata as JSON object: %w", err)
	}

	_, err = f.Write(marshalled)
	if err != nil {
		return fmt.Errorf("can't write metadata to file: %w", err)
	}

	return nil
}

/*
Path eturns the root path of the release inside the file system
*/
func (release Release) Path() string {
	return path.Join(release.Repo.Name, release.Version)
}

func (release Release) String() string {
	return release.Repo.Name + " " + release.Version
}
