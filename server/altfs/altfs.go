package altfs

/*
AltFS is a type representing an AltFS structure, with a list of read mounts and write mounts.
*/
type AltFS struct {
	reads  []string
	writes []string
}

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
	panic("Not implemented")
}

// Create tries to create a file similarly to os.Create
func (fs AltFS) Create(name string) (WriteFile, error) {
	panic("Not implemented")
}

// Exists checks if a file exists in the file system.
func (fs AltFS) Exists(name string) bool {
	panic("Not implemented")
}
