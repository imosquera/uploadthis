package main

import (
	"github.com/imosquera/uploadthis"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
)

func main() {
	uploadthis.ParseOpts()
	for _, monitorDir := range uploadthis.Settings.MonitorDirs {
		uploadFiles := monitor.GetUploadFiles(monitorDir.Path)
		prehooks := hooks.GetPrehooks(monitorDir.PreHooks)
		for _, prehook := range prehooks {
			uploadFiles, _ = prehook.RunPrehook(uploadFiles)
		}
		//upload files
		//post commit hooks
	}
}
