package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util"
	"io"
	"log"
	"path"
)

var registeredPrehooks map[string]commands.Commander

func RegisterPrehook(name string, prehook commands.Commander) {
	if registeredPrehooks == nil {
		registeredPrehooks = make(map[string]commands.Commander, 5)
	}
	registeredPrehooks[name] = prehook
}

func GetPrehookCommands(prehooks []string, prehookCommands map[string]commands.Commander) {
	for _, prehook := range prehooks {
		prehookCommands[prehook] = registeredPrehooks[prehook]
	}
}

//*********************
// COMPRESSOR
//*********************
type GzipFileCompressor struct{}

func (self *GzipFileCompressor) Compress(filepath string) (string, error) {
	inFile, err := util.Fs.Open(filepath)
	if err != nil {
		log.Fatal("Error for file " + filepath + " " + err.Error())
	}

	//create the gzip file
	gzipPath := path.Join(path.Dir(filepath), path.Base(filepath)+".gz")
	outFile, err := util.Fs.Create(gzipPath)
	util.LogPanic(err)

	//create a new gzip writer so we can copy the bytes
	gzipWriter, err := gzip.NewWriterLevel(outFile, gzip.BestCompression)
	util.LogPanic(err)
	_, err = io.Copy(gzipWriter, inFile)
	util.LogPanic(err)

	gzipWriter.Close()
	return gzipPath, err
}

type Compressor interface {
	Compress(filepath string) (string, error)
}

func NewCompressPrehook() *CompressPrehook {
	return &CompressPrehook{
		compressor: &GzipFileCompressor{},
	}
}

type CompressPrehook struct {
	commands.Command
	compressor Compressor
}

//this function will compress the infile and
//return a file pointer that is readible
func (c CompressPrehook) Run() ([]string, error) {
	var err error
	newFiles := make([]string, 0, len(c.UploadFiles))
	for _, uploadFile := range c.UploadFiles {
		uploadFile, _ := c.compressor.Compress(uploadFile)
		newFiles = append(newFiles, uploadFile)
	}
	return newFiles, err
}
