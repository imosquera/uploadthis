package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util"
	log "github.com/cihub/seelog"
	"io"
	"path"
	"os"
	"path/filepath"
	"crypto/rand"
	"encoding/hex"
)

//*********************
// COMPRESSOR
//*********************
type GzipFileCompressor struct{}

func (self *GzipFileCompressor) Compress(filepath string) (string, error) {
	inFile, err := util.Fs.Open(filepath)
	if err != nil {
		log.Critical("Error for file " + filepath + " " + err.Error())
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

	err = os.Remove(filepath)
	util.LogPanic(err)

	gzipWriter.Close()
	return gzipPath, err
}

type Compressor interface {
	Compress(filepath string) (string, error)
}

func NewCompressPrehook() *CompressPrehook {
	return &CompressPrehook{
		commands.NewFileStateCommand(),
		&GzipFileCompressor{},
	}
}

type CompressPrehook struct {
	*commands.Command
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

//**************
// Renamer
//**************
type Renamer interface {
	Rename(filepath string) (string, error)
	SetProperties(...interface {})
}

type UniqueSuffixRenamer struct{
	Renamer
	suffixLength int
}

func (self *UniqueSuffixRenamer) Rename(filename string) (string, error) {
	var extension = filepath.Ext(filename)
	var name = filename[0:len(filename)-len(extension)]
	suffix, _ := self.GenerateRandomString(self.suffixLength)
	var newName = name + "-" + suffix + extension

	err := os.Rename(filename, newName)
	if err != nil {
		log.Critical("Error for file " + filename + " " + err.Error())
		util.LogPanic(err)
	}
	return newName, err
}

func (self *UniqueSuffixRenamer) GenerateRandomString(length int) (string, error) {
	uuid := make([]byte, length)
	_, err := rand.Read(uuid)
	if err != nil {
		log.Critical("Error generating random string " + err.Error())
		util.LogPanic(err)
		return "", err
	}
	return hex.EncodeToString(uuid), nil
}

type RenamePrehook struct {
	*commands.Command
	renamer Renamer
}

func NewRenamePrehook() *RenamePrehook {
	return &RenamePrehook{
		commands.NewFileStateCommand(),
		&UniqueSuffixRenamer{suffixLength: 2},
	}
}

//this function will rename the infile and
//return a new file name
func (self RenamePrehook) Run() ([]string, error) {
	var err error
	newFiles := make([]string, 0, len(self.UploadFiles))
	for _, uploadFile := range self.UploadFiles {
		newFile, the_err := self.renamer.Rename(uploadFile)
		log.Info("Rename: ", uploadFile, " -> ", newFile)
		if the_err != nil {
			err = the_err
		}
		newFiles = append(newFiles, newFile)
	}
	return newFiles, err
}
