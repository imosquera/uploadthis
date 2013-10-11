package main

import (
	log "github.com/cihub/seelog"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/execution"
)

func main() {
	defer log.Flush()
	log.Info("uploadthis has started up...")
	//this parsing options and does some additional configurations
	conf.ParseOpts()

	//this begins the execution chain
	commandManager := execution.NewDefaultCommandManager()
	commandManager.ExecuteCommandsForMonitors()
}
