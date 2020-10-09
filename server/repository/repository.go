package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/yotam180/treasury/altfs"
)

/*
Bucket wraps a file system with multiple repositories
*/
type Bucket struct {
	altfs.FileSystem
}

/*
Repo is a collection of releases for a project
*/
type Repo struct {
	bucket *Bucket

	Name        string
	LastUpdated time.Time
}

/*
Release is a specific release of a
*/
type Release struct {
	repo   *Repo
	bucket *Bucket

	Version string
}

/*
NewBucket creates a new bucket for a file system.
*/
func NewBucket(fs altfs.FileSystem) *Bucket {
	return &Bucket{fs}
}

/*
NewRepo creates a new repository object. It does not create the repository physically, but just returns an object representing it.
TODO: Create it physically (?)
*/
func (bucket *Bucket) NewRepo(name string, updated time.Time) *Repo {
	return &Repo{bucket, name, updated}
}

/*
ListRepositories returns a string of repository objects in the bucket.
TODO: Add more metadata?
*/
func (bucket *Bucket) ListRepositories() ([]*Repo, error) {
	dirs, err := bucket.ListDir("/")
	if err != nil {
		return nil, fmt.Errorf("can't list repositories: %w", err)
	}

	repos := make([]*Repo, 0, len(dirs))
	for _, dir := range dirs {
		if dir.IsDir() {
			dir.ModTime()
			repos = append(repos, bucket.NewRepo(dir.Name(), dir.ModTime()))
		}
	}

	return repos, nil
}

/*
ListReleases returns an array of all releases in the repository.
*/
func (repo *Repo) ListReleases() ([]Release, error) {
	subDirs, err := repo.bucket.ListDir(repo.Name)
	if err != nil {
		return nil, fmt.Errorf("can't list versions: %w", err)
	}

	releases := make([]Release, 0)

	for _, dir := range subDirs {
		if dir.IsDir() {
			releases = append(releases, Release{repo, repo.bucket, dir.Name()})
		}
	}

	return releases, nil
}

/*
CreateRelease opens a new release folder in a repo and returns the release object
*/
func (repo *Repo) CreateRelease(version string) (Release, error) {
	// TODO: Verify that this doesn't enter some other release files folder or anything, and that it is acceptable as a release.
	err := repo.bucket.Mkdir(path.Join(repo.Name, version))
	if err != nil {
		return Release{}, err
	}

	return Release{repo, repo.bucket, version}, nil
}

/*
GetMetadata returns the metadata object for the release
*/
func (release Release) GetMetadata() map[string]interface{} {
	f, err := release.bucket.Open(path.Join(release.Path(), "metadata.json"))
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

	f, err := release.bucket.Create(path.Join(release.Path(), "metadata.json"))
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
AddFile adds a file to the release.
If the file exists, this fails with an error
*/
func (release Release) AddFile(fileName string, blob io.Reader) error {
	fileDirPath := path.Join(release.Path(), "files")
	filePath := path.Join(fileDirPath, fileName)

	if release.bucket.Exists(filePath) {
		return fmt.Errorf("file %s already exists in release %s", fileName, release.Version)
	}

	err := release.bucket.Mkdir(fileDirPath)
	if err != nil {
		return fmt.Errorf("cannot create release file directory: %w", err)
	}

	f, err := release.bucket.Create(filePath)
	if err != nil {
		return fmt.Errorf("cannot create file in release: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, blob)
	if err != nil {
		return fmt.Errorf("cannot copy file into release: %w", err)
	}

	return nil
}

/*
GetFile opens a specific file in a release for reading
*/
func (release Release) GetFile(fileName string) (altfs.ReadFile, error) {
	filePath := path.Join(release.Path(), "files", fileName)

	file, err := release.bucket.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("can't open release file: %w", err)
	}

	return file, nil
}

/*
ListFiles returns the list of files in a release
*/
func (release Release) ListFiles() []string {
	return release.repo.ListDirRecursive(path.Join(release.Path(), "files"))
}

/*
ListDirRecursive returns all files in a directory, recursively.
*/
func (repo Repo) ListDirRecursive(dirPath string) []string {
	prefix := dirPath
	dirPath = "/"
	q := queue.New(64)

	result := []string{}

	for {
		dirContent, err := repo.bucket.ListDir(path.Join(prefix, dirPath))
		if err != nil {
			return []string{}
		}

		for _, fileInfo := range dirContent {
			if fileInfo.IsDir() {
				q.Put(path.Join(dirPath, fileInfo.Name()))
			} else {
				result = append(result, path.Join(dirPath, fileInfo.Name()))
			}
		}

		if q.Len() == 0 {
			break
		} else {
			next, err := q.Get(1)
			if err != nil {
				return result // ?
			}
			dirPath = next[0].(string)
		}
	}

	return result
}

/*
Path eturns the root path of the release inside the file system
*/
func (release Release) Path() string {
	return path.Join(release.repo.Name, release.Version)
}

func (release Release) String() string {
	return release.repo.Name + " " + release.Version
}
