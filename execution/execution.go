package execution

import (
	log "github.com/cihub/seelog"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/upload"
	"github.com/imosquera/uploadthis/util"
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
		commanders["upload"] = upload.NewUploadCommand(monitorDir)
		hooks.GetHookCommands(hooks.POSTHOOK, monitorDir.PostHooks, commanders)
	}
	return commanders
}

type CommandExecutor interface {
	ExecuteCommands(map[string]commands.Commander, conf.MonitorDir)
}

type SequentialCommandExecutor struct{}

func (self *SequentialCommandExecutor) ExecuteCommands(commandList map[string]commands.Commander, monitorDir conf.MonitorDir) {
	uploadFiles := util.GetFilesFromDir(monitorDir.Path)
	for name, command := range commandList {
		log.Infof("Using command:%s for monitor path: %s", name, monitorDir.Path)
		command.SetMonitor(monitorDir)
		command.SetName(name)
		command.SetUploadFiles(uploadFiles)
		command.Prepare()
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
	log.Infof("Found %d monitor dirs", len(conf.Settings.MonitorDirs))
	for _, monitorDir := range conf.Settings.MonitorDirs {
		log.Info("Working on monitor dir: ", monitorDir.Path)
		commandList := self.Producer.CreateCommandList(&monitorDir)
		self.Executor.ExecuteCommands(commandList, monitorDir)
	}
}
