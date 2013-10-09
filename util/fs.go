package util

import (
	"os"
)

var Fs fileSystem = OsFs{}

type fileSystem interface {
	Open(name string) (*os.File, error)
	Create(name string) (*os.File, error)
	Stat(name string) (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type OsFs struct{}

func (OsFs) Open(name string) (*os.File, error)    { return os.Open(name) }
func (OsFs) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func (OsFs) Create(name string) (*os.File, error)  { return os.Create(name) }
