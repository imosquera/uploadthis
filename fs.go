package uploadthis

import (
	"io"
	"os"
)

var Fs fileSystem = osFS{}

type fileSystem interface {
	Open(name string) (file, error)
	Create(name string) (file, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Open(name string) (file, error)   { return os.Open(name) }
func (osFS) Create(name string) (file, error) { return os.Create(name) }
func (osFS) Copy(name string) (file, error)   { return os.Create(name) }
