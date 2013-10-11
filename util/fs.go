package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
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

//more utility functions to make a dir
func MakeDir(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Panic(err)
	}
}

//just gets a directory listing
func GetFilesFromDir(dirPath string) []string {
	allFiles := make([]string, 0)
	files, _ := ioutil.ReadDir(dirPath)
	for _, dirFile := range files {
		if !dirFile.IsDir() {
			filePath := path.Join(dirPath, dirFile.Name())
			allFiles = append(allFiles, filePath)
		}
	}
	return allFiles
}
