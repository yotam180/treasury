package repo

import (
	"fmt"

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

func (release Release) String() string {
	return release.Repo.Name + " " + release.Version
}
