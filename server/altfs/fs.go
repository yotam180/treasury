package altfs

import "io"

/*
FileSystem is an abstraction of a file system providing a simple Open, Create, Exists? interface.
*/
type FileSystem interface {
	Open(name string) (ReadFile, error)
	Create(name string) (WriteFile, error)
	Exists(names string) bool
}

/*
ReadFile is an interface for a file that can be read from
*/
type ReadFile interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
}

/*
WriteFile is an interface for a file that can be written to
*/
type WriteFile interface {
	io.Closer
	io.Writer
	io.WriterAt
	io.Seeker
}
