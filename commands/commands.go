package commands

import (
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/util"
	"os"
	"path"
)

type Commander interface {
	SetMonitor(conf.MonitorDir)
	SetName(string)
	SetUploadFiles([]string)
	Prepare()
	Run() ([]string, error)
}

type Command struct {
	monitorDir     conf.MonitorDir
	UploadFiles    []string
	statePersistor StatePersistor
	Name           string
}

func NewFileStateCommand() *Command {
	return &Command{
		statePersistor: &FileStatePersistor{},
	}
}

func (self *Command) SetMonitor(monitor conf.MonitorDir) {
	self.monitorDir = monitor
}

func (self *Command) SetName(name string) {
	self.Name = name
}

func (self *Command) SetUploadFiles(uploadFiles []string) {
	self.UploadFiles = uploadFiles
}

func (self *Command) Prepare() {
	workDir := path.Join(self.monitorDir.Path, self.Name)
	self.statePersistor.SetWorkDir(workDir)
	self.statePersistor.SetActive(self.UploadFiles)
	self.UploadFiles = self.statePersistor.GetActive()
}

type StatePersistor interface {
	SetWorkDir(string)
	SetActive([]string) []string
	GetActive() []string
}

type FileStatePersistor struct {
	WorkDir string
}

func (self *FileStatePersistor) SetWorkDir(workDir string) {
	self.WorkDir = workDir
}

func (self *FileStatePersistor) GetActive() []string {
	return util.GetFilesFromDir(self.WorkDir)
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
