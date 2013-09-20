package hooks

import (
	"compress/gzip"
	"errors"
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
func (c CompressPrehook) RunPrehook(inFile io.Reader, fileInfo os.FileInfo) (gzipReader io.Reader, err error) {
	gzipPath := "/tmp/" + fileInfo.Name() + ".gz"

	outFile, err := os.Create(gzipPath)

	gzipWriter, err := NewCompressor(outFile)

	bytesWritten, err := Copy(gzipWriter, inFile)
	defer gzipWriter.Close()

	if bytesWritten != fileInfo.Size() {
		err = errors.New("Bytes written does not match InFile byte size")
	}
	gzipReader, _ = os.Open(gzipPath)
	return gzipReader, err
}
