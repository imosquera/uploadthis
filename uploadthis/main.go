package main

import (
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"log"
	"path"
)

func main() {
	//this setups the logger so that it prints file numbers
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//this parsing options and does some additional configurations
	conf.ParseOpts()
	executeCommandsForMonitors()
}

var createCommandList = func(monitorDir *conf.MonitorDir) map[string]commands.Commander {
	commanders := make(map[string]commands.Commander, 5)
	for _, monitorDir := range conf.Settings.MonitorDirs {
		hooks.GetPrehookCommands(monitorDir.PreHooks, commanders)
		//upload file commands
		//post commit commands
	}
	return commanders
}

var executeCommands = func(commanders map[string]commands.Commander, monitorDir conf.MonitorDir) {
	uploadFiles := monitor.GetUploadFiles(monitorDir.Path)
	for name, command := range commanders {
		workDir := path.Join(monitorDir.Path, name)
		uploadFiles = commands.MoveToWorkDir(workDir, uploadFiles)
		uploadFiles = commands.UpdateResumeFiles(workDir, uploadFiles)
		uploadFiles, _ = command.Run(uploadFiles)
	}
}

var executeCommandsForMonitors = func() {
	log.Println("%v", conf.Settings.MonitorDirs)
	for _, monitorDir := range conf.Settings.MonitorDirs {
		log.Println("%v", monitorDir)
		commandList := createCommandList(&monitorDir)
		executeCommands(commandList, monitorDir)
	}
}
