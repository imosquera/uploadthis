package commands

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Commander interface {
	SetName(string)
	SetUploadFiles(uploadFiles []string)
	Prepare(workDir string)
	Run() ([]string, error)
}

type Command struct {
	Name           string
	UploadFiles    []string
	statePersistor StatePersistor
}

func NewFileStateCommander() *Command {
	return &Command{
		statePersistor: &FileStatePersistor{},
	}
}

func (self *Command) SetName(name string) {
	self.Name = name
}

func (self *Command) SetUploadFiles(uploadFiles []string) {
	self.UploadFiles = uploadFiles
}

func (self *Command) Prepare(workDir string) {
	self.statePersistor.SetWorkDir(workDir)
	self.UploadFiles = self.statePersistor.SetActive(self.UploadFiles)
	self.statePersistor.AppendResume(self.UploadFiles)
}

type StatePersistor interface {
	SetWorkDir(string)
	SetActive([]string) []string
	AppendResume([]string)
}

type FileStatePersistor struct {
	WorkDir string
}

func (self *FileStatePersistor) SetWorkDir(workDir string) {
	self.WorkDir = workDir
}

func (self *FileStatePersistor) AppendResume(filePaths []string) {

}

//set the active files by moving them into their own directory
func (self *FileStatePersistor) SetActive(filePaths []string) []string {
	newUploadFiles := make([]string, 0, len(filePaths))
	for _, uploadFileInfo := range filePaths {
		filename := path.Base(uploadFileInfo)
		workFile := path.Join(self.WorkDir, filename)
		os.Rename(uploadFileInfo, workFile)
	}
	return newUploadFiles
}

func MakeDir(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Panic(err)
	}
}

func GetUploadFiles(dirPath string) []string {
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
