package altfs

import (
	"io"
	"os"
)

/*
FileSystem is an abstraction of a file system providing a simple Open, Create, Exists? interface.
*/
type FileSystem interface {
	Open(name string) (ReadFile, error)
	Create(name string) (WriteFile, error)

	Stat(name string) (os.FileInfo, error)

	Exists(names string) bool

	Mkdir(dirPath string) error
	ListDir(dirPath string) ([]os.FileInfo, error) // TODO: Generalize return type?
}

/*
ReadFile is an interface for a file that can be read from
*/
type ReadFile interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker

	Name() string
}

/*
WriteFile is an interface for a file that can be written to
*/
type WriteFile interface {
	io.Closer
	io.Writer
	io.WriterAt
	io.Seeker

	Name() string
}
