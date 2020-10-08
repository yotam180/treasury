package altfs

import (
	"fmt"
	"os"
	"path"
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

	if fs.reads == nil {
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

	if fs.writes == nil {
		return nil, fmt.Errorf("read-only file system")
	}

	for _, mount := range fs.writes {
		file, err := os.Create(path.Join(mount, removeLeadingSlash(name)))
		if err == nil {
			return file, nil
		}
	}

	return nil, fmt.Errorf("can't open file %s on any of the write-enable mount points", name)
}

// Exists checks if a file exists in one of the READ mount points.
func (fs AltFS) Exists(name string) bool {
	for _, mount := range fs.reads {
		if _, err := os.Stat(path.Join(mount, removeLeadingSlash(name))); err == nil {
			return true
		}
	}

	return false
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
