package commands

import (
	"github.com/imosquera/uploadthis/monitor"
	"log"
	"os"
	"path"
)

type Commander interface {
	Run(uploadFiles []string) ([]string, error)
}

var MakeDirWithWarning = func(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Println(err)
	}
}

func MakeWorkDir(dirPath string) {
	MakeDirWithWarning(dirPath)
}

var MoveToWorkDir = func(workDir string, uploadFiles []string) []string {
	newUploadFiles := make([]string, 0, len(uploadFiles))
	for _, uploadFileInfo := range uploadFiles {
		filename := path.Base(uploadFileInfo)
		workFile := path.Join(workDir, filename)
		os.Rename(uploadFileInfo, workFile)
	}
	return newUploadFiles
}

var UpdateResumeFiles = func(workDir string, uploadFiles []string) []string {
	paths := monitor.GetUploadFiles(workDir)
	return append(uploadFiles, paths...)
}
