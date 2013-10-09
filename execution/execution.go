package execution

import (
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"path"
)

// func MakeDirWithWarning(dirPath string) {
// 	err := os.Mkdir(dirPath, 0755)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func MakeWorkDir(dirPath string) {
// 	MakeDirWithWarning(dirPath)
// }

// func MoveToWorkDir(workDir string, uploadFiles []string) []string {
// 	newUploadFiles := make([]string, 0, len(uploadFiles))
// 	for _, uploadFileInfo := range uploadFiles {
// 		filename := path.Base(uploadFileInfo)
// 		workFile := path.Join(workDir, filename)
// 		os.Rename(uploadFileInfo, workFile)
// 	}
// 	return newUploadFiles
// }

type CommandProducer interface {
	CreateCommandList(monitorDir *conf.MonitorDir) map[string]commands.Commander
}

type ConfigCommandProducer struct{}

func (self *ConfigCommandProducer) CreateCommandList(monitorDir *conf.MonitorDir) map[string]commands.Commander {
	commanders := make(map[string]commands.Commander, 5)
	for _, monitorDir := range conf.Settings.MonitorDirs {
		hooks.GetPrehookCommands(monitorDir.PreHooks, commanders)
		//upload file commands
		//post commit commands
	}
	return commanders
}

type CommandExecutor interface {
	ExecuteCommands(map[string]commands.Commander, conf.MonitorDir)
}

type SequentialCommandExecutor struct{}

func (self *SequentialCommandExecutor) ExecuteCommands(commands map[string]commands.Commander, monitorDir conf.MonitorDir) {
	uploadFiles := monitor.GetUploadFiles(monitorDir.Path)
	for name, command := range commands {
		workDir := path.Join(monitorDir.Path, name)
		command.SetUploadFiles(uploadFiles)
		command.Prepare(workDir)
		uploadFiles, _ = command.Run()
	}
}

func NewDefaultCommandManager() CommandManager {
	return DefaultCommandManager{
		Executor: &SequentialCommandExecutor{},
		Producer: &ConfigCommandProducer{},
	}
}

type CommandManager interface {
	ExecuteCommandsForMonitors()
}

type DefaultCommandManager struct {
	Producer CommandProducer
	Executor CommandExecutor
}

func (self DefaultCommandManager) ExecuteCommandsForMonitors() {
	for _, monitorDir := range conf.Settings.MonitorDirs {
		commandList := self.Producer.CreateCommandList(&monitorDir)
		self.Executor.ExecuteCommands(commandList, monitorDir)
	}
}
