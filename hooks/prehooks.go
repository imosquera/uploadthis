package hooks

import (
	"compress/gzip"
	"errors"
	"github.com/imosquera/uploadthis/monitor"
	"io"
	"os"
)

type Prehooker interface {
	RunPrehook(uploadFiles []monitor.UploadFileInfo) ([]monitor.UploadFileInfo, error)
}

type Prehook struct{}

type CompressPrehook struct{ Prehook }

func GetPrehooks(prehooks []string) []Prehooker {
	prehookers := make([]Prehooker, 2)
	for _, prehook := range prehooks {
		if prehook == "compress" {
			prehookers = append(prehookers, CompressPrehook{})
		}
	}
	return prehookers
}

var compressFile = func(uploadFile monitor.UploadFileInfo) (monitor.UploadFileInfo, error) {
	inFile, _ := os.Open(uploadFile.Path)
	gzipPath := "/tmp/" + uploadFile.Info.Name() + ".gz"
	outFile, err := os.Create(gzipPath)
	gzipWriter, err := gzip.NewWriterLevel(inFile, gzip.BestCompression)
	bytesWritten, err := io.Copy(gzipWriter, inFile)

	if bytesWritten != uploadFile.Info.Size() {
		err = errors.New("Bytes written does not match InFile byte size")
	}
	gzipWriter.Close()
	info, err := outFile.Stat()
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
