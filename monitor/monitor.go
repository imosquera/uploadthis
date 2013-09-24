package monitor

import (
	"os"
	"path/filepath"
)

type UploadFileInfo struct {
	Path string
	Info os.FileInfo
}

var GetUploadFiles = func(path string) []UploadFileInfo {
	allFiles := make([]UploadFileInfo, 0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		allFiles = append(allFiles, UploadFileInfo{Path: path, Info: info})
		return nil
	})
	return allFiles
}
