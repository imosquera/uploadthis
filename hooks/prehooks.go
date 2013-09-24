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

type CompressPrehook struct{ Prehook }

func GetPrehooks(prehooks []string) []Prehooker {
	prehookers := make([]Prehooker, 0)
	for _, prehook := range prehooks {
		if prehook == "compress" {
			prehookers = append(prehookers, CompressPrehook{})
		}
	}
	return prehookers
}

var compressFile = func(uploadFile monitor.UploadFileInfo) (monitor.UploadFileInfo, error) {
	inFile, err := os.Open(uploadFile.Path)
	if err != nil {
		log.Fatal("Error for file " + uploadFile.Path + " " + err.Error())
	}

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
