package altfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

/*
AltFS is a type representing an AltFS structure, with a list of read mounts and write mounts.
*/
type AltFS struct {
	reads  []string
	writes []string
}

// Assert that AltFS correctly implements FileSystem
var _ FileSystem = (*AltFS)(nil)

/*
NewFS will create a new AltFS structure over a list of read-enabled mounts and write-enabled mounts.
*/
func NewFS(reads []string, writes []string) AltFS {
	if reads == nil {
		reads = make([]string, 0)
	}

	if writes == nil {
		writes = make([]string, 0)
	}

	return AltFS{reads, writes}
}

// Open tries to open a file similarly to os.Open
func (fs AltFS) Open(name string) (ReadFile, error) {
	if len(fs.reads) == 0 {
		return nil, fmt.Errorf("write-only file system")
	}

	for _, mount := range fs.reads {
		file, err := os.Open(path.Join(mount, removeLeadingSlash(name)))
		if err == nil {
			return file, nil
		}
	}

	return nil, fmt.Errorf("can't open file %s on any of the read-enabled mount points", name)
}

// Create tries to create a file similarly to os.Create
func (fs AltFS) Create(name string) (WriteFile, error) {
	if len(fs.writes) == 0 {
		return nil, fmt.Errorf("read-only file system")
	}

	for _, mount := range fs.writes {
		file, err := os.Create(fs.mkpath(mount, name))
		if err == nil {
			return file, nil
		}
	}

	return nil, fmt.Errorf("can't open file %s on any of the write-enable mount points", name)
}

// Stat stats a file in the file system
func (fs AltFS) Stat(name string) (os.FileInfo, error) {
	var lastError error = nil

	for _, mount := range fs.reads {
		info, err := os.Stat(fs.mkpath(mount, name))
		if err == nil {
			return info, nil
		}
		lastError = err
	}

	return nil, lastError
}

// Exists checks if a file exists in one of the READ mount points.
func (fs AltFS) Exists(name string) bool {
	for _, mount := range fs.reads {
		if _, err := os.Stat(fs.mkpath(mount, name)); err == nil {
			return true
		}
	}

	return false
}

/*
Mkdir creates a directory
*/
func (fs AltFS) Mkdir(dirPath string) error {

	if len(fs.writes) == 0 {
		return fmt.Errorf("read-only file system")
	}

	for _, mount := range fs.writes {
		err := os.MkdirAll(fs.mkpath(mount, dirPath), os.ModePerm)
		if err == nil {
			return nil
		}
		fmt.Println(err)
	}

	return fmt.Errorf("can't create directory on any of the mount points")
}

/*
ListDir returns all files and subdirectories in a directory in the file system, across all read mount-points
*/
func (fs AltFS) ListDir(dirPath string) ([]os.FileInfo, error) {
	results := make(map[string]os.FileInfo, 0)

	someSuccess := false

	for _, mount := range fs.reads {
		list, err := ioutil.ReadDir(fs.mkpath(mount, dirPath))
		if err != nil {
			continue // TODO: Is this a good decision?
		}

		someSuccess = true
		for _, result := range list {
			if _, contains := results[result.Name()]; !contains {
				results[result.Name()] = result
			}
		}
	}

	if !someSuccess {
		return nil, fmt.Errorf("could not find directory %s", dirPath)
	}

	names := make([]string, 0, len(results))
	for _, result := range results {
		names = append(names, result.Name())
	}

	sort.Strings(names)

	arrayResults := make([]os.FileInfo, len(results))
	for index, name := range names {
		arrayResults[index] = results[name]
	}

	return arrayResults, nil
}

func (fs AltFS) mkpath(mount, filePath string) string {
	return path.Join(mount, removeLeadingSlash(filePath))
}

func removeLeadingSlash(filePath string) string {
	switch {
	case strings.HasPrefix(filePath, "/"),
		strings.HasPrefix(filePath, "\\"):
		return filePath[1:]

	default:
		return filePath
	}
}
