package util

import (
	"io"
	"os"
)

var Fs fileSystem = OsFs{}

type fileSystem interface {
	Open(name string) (OSFile, error)
	Create(name string) (*os.File, error)
	Stat(name string) (os.FileInfo, error)
}

type OSFile interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type OsFs struct{}

func (OsFs) Open(name string) (OSFile, error) { return os.Open(name) }

//TODO: we have to convert the following into an interface
func (OsFs) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (OsFs) Create(name string) (*os.File, error)  { return os.Create(name) }
