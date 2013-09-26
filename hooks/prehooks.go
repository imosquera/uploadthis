package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/monitor"
	"github.com/imosquera/uploadthis/util"
	"io"
	"log"
	"os"
	"path"
	//"runtime/debug"
)

type Prehooker interface {
	RunPrehook(uploadFiles []monitor.UploadFileInfo) ([]monitor.UploadFileInfo, error)
}

type Prehook struct {
	monitor.UploadFileInfo
}

type registeredPrehooks map[string]Prehooker

func RegisterPrehook(name string, prehooker Prehooker) {
	registeredPrehooks[name] = prehooker
}

func GetPrehooks(prehooks []string) []Prehooker {
	prehookers := make([]Prehooker, 0, 2)
	for _, prehook := range prehooks {
		prehookers = append(prehookers, registeredPrehooks[prehook])
	}
	return prehookers
}

var compressFile = func(uploadFile monitor.UploadFileInfo) (monitor.UploadFileInfo, error) {
	inFile, err := os.Open(uploadFile.Path)
	if err != nil {
		log.Fatal("Error for file " + uploadFile.Path + " " + err.Error())
	}
	log.Println("Working on file", uploadFile.Path)
	gzipPath := path.Join(path.Dir(uploadFile.Path), "doing", uploadFile.Info.Name()+".gz")
	outFile, err := os.Create(gzipPath)
	util.LogPanic(err)

	gzipWriter, err := gzip.NewWriterLevel(outFile, gzip.BestCompression)
	util.LogPanic(err)

	bytesWritten, err := io.Copy(gzipWriter, inFile)
	util.LogPanic(err)

	if bytesWritten != uploadFile.Info.Size() {
		log.Fatal("Bytes written do not match inFile byte size")
	}
	gzipWriter.Close()
	info, err := outFile.Stat()
	util.LogPanic(err)

	return monitor.UploadFileInfo{Path: gzipPath, Info: info}, err
}

type CompressPrehook struct{ Prehook }

//this function will compress the infile and
//return a file pointer that is readible
func (c CompressPrehook) RunPrehook(uploadFiles []monitor.UploadFileInfo) ([]monitor.UploadFileInfo, error) {
	var err error
	newFiles := make([]monitor.UploadFileInfo, len(uploadFiles))
	for _, uploadFile := range uploadFiles {
		uploadFile, _ := compressFile(uploadFile)
		newFiles = append(newFiles, uploadFile)
	}
	return newFiles, err
}
