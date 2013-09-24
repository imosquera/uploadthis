package main

import (
	"fmt"
	"github.com/imosquera/uploadthis"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"log"
)

func main() {
	//this setups the logger so that it prints nicely
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//this parsing options and does some additional configurations
	uploadthis.ParseOpts()
	for _, monitorDir := range uploadthis.Settings.MonitorDirs {
		hooks.MakeWorkDirs(monitorDir.Path)
		uploadFiles := monitor.GetUploadFiles(monitorDir.Path)
		prehooks := hooks.GetPrehooks(monitorDir.PreHooks)
		for _, prehook := range prehooks {
			fmt.Println("%#v", uploadFiles)
			uploadFiles, _ = prehook.RunPrehook(uploadFiles)
		}
		//upload files
		//post commit hooks
	}
}
