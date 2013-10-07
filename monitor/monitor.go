package monitor

import (
	"io/ioutil"
	"os"
	"path"
)

type UploadFileInfo struct {
	Path string
	Info os.FileInfo
}

var GetUploadFiles = func(dirPath string) []UploadFileInfo {
	allFiles := make([]UploadFileInfo, 0)
	files, _ := ioutil.ReadDir(dirPath)
	for _, dirFile := range files {
		if !dirFile.IsDir() {
			filePath := path.Join(dirPath, dirFile.Name())
			allFiles = append(allFiles, UploadFileInfo{Path: filePath, Info: dirFile})
		}
	}
	return allFiles
}
