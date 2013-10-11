package execution

import (
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"github.com/imosquera/uploadthis/upload"
	"path"
)

type CommandProducer interface {
	CreateCommandList(monitorDir *conf.MonitorDir) map[string]commands.Commander
}

type ConfigCommandProducer struct{}

//builds a command list based on the monitor dirs in the yaml config
func (self *ConfigCommandProducer) CreateCommandList(monitorDir *conf.MonitorDir) map[string]commands.Commander {
	commanders := make(map[string]commands.Commander, 5)
	for _, monitorDir := range conf.Settings.MonitorDirs {
		hooks.GetHookCommands(hooks.PREHOOK, monitorDir.PreHooks, commanders)
		commanders["upload"] = upload.NewUploadCommand(monitorDir.Bucket)
		hooks.GetHookCommands(hooks.POSTHOOK, monitorDir.PreHooks, commanders)
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
