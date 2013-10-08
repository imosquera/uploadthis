package monitor

import (
	"io/ioutil"
	"path"
)

var GetUploadFiles = func(dirPath string) []string {
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
