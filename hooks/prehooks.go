package hooks

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

type Prehooker interface {
	RunPrehook(inFile io.Reader) io.Writer
}

type Prehook struct{}

type CompressPrehook struct{ Prehook }

func fatalLog(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

var NewCompressor = func(writer io.Writer) (*gzip.Writer, error) {
	return gzip.NewWriterLevel(writer, gzip.BestCompression)
}

var Copy = func(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

//this function will compress the infile and
//return a file pointer that is readible
func (c CompressPrehook) RunPrehook(inFile io.Reader, fileInfo os.FileInfo) io.Reader {
	gzipPath := "/tmp/" + fileInfo.Name() + ".gz"

	outFile, err := os.Create(gzipPath)
	fatalLog(err)

	gzipWriter, err := NewCompressor(outFile)
	fatalLog(err)

	bytesWritten, err := Copy(gzipWriter, inFile)
	fatalLog(err)
	defer gzipWriter.Close()

	if bytesWritten != fileInfo.Size() {
		log.Fatal("Bytes written doesn't match infile byte size")
	}
	newFile, _ := os.Open(gzipPath)
	return newFile
}
