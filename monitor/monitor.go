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
	alldirs := make([]UploadFileInfo, 5)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		alldirs = append(alldirs, UploadFileInfo{Path: path, Info: info})
		return nil
	})
	return alldirs
}
