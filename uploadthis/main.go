package main

import (
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"log"
)

func main() {
	//this setups the logger so that it prints file numbers
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//this parsing options and does some additional configurations
	conf.ParseOpts()
	for _, monitorDir := range conf.Settings.MonitorDirs {
		conf.MakeWorkDirs(monitorDir.Path)
		uploadFiles := monitor.GetUploadFiles(monitorDir.Path)
		prehooks := hooks.GetPrehooks(monitorDir.PreHooks)
		for _, prehook := range prehooks {
			uploadFiles, _ = prehook.RunPrehook(uploadFiles)
		}
		//upload files
		//post commit hooks
	}
}
