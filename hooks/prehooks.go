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

type Prehook struct {
	uploadFiles []string
}

func (self *Prehook) Prepare(workDir string) {
	//uploadFiles = MoveToWorkDir(workDir, uploadFiles)
	//paths := monitor.GetUploadFiles(workDir)
	//uploadFiles = append(uploadFiles, paths...)
	//uploadFiles, _ = command.Run(uploadFiles)
}

func (self *Prehook) SetUploadFiles(uploadFiles []string) {
	self.uploadFiles = uploadFiles
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
	Prehook
	compressor Compressor
}

//this function will compress the infile and
//return a file pointer that is readible
func (c CompressPrehook) Run() ([]string, error) {
	var err error
	newFiles := make([]string, 0, len(c.uploadFiles))
	for _, uploadFile := range c.uploadFiles {
		uploadFile, _ := c.compressor.Compress(uploadFile)
		newFiles = append(newFiles, uploadFile)
	}
	return newFiles, err
}
