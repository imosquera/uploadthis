package main

import (
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/execution"
	"log"
)

func main() {
	//this setups the logger so that it prints file numbers
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//this parsing options and does some additional configurations
	conf.ParseOpts()

	commandManager := execution.NewDefaultCommandManager()
	commandManager.ExecuteCommandsForMonitors()
}
