package hooks

import (
	"github.com/imosquera/uploadthis/commands"
)

const (
	PREHOOK = iota
	POSTHOOK
)

var registeredPrehooks map[string]commands.Commander
var registeredPosthooks map[string]commands.Commander

func init() {
	registeredPrehooks = make(map[string]commands.Commander, 5)
	registeredPosthooks = make(map[string]commands.Commander, 5)
}

func GetHookMap(hookType int) map[string]commands.Commander {
	var hookMap map[string]commands.Commander
	if hookType == PREHOOK {
		hookMap = registeredPrehooks
	} else {
		hookMap = registeredPosthooks
	}
	return hookMap
}

func RegisterHook(hookType int, name string, prehook commands.Commander) {
	hookMap := GetHookMap(hookType)
	hookMap[name] = prehook
}

func GetHookCommands(hookType int, hooks []string, hookCommands map[string]commands.Commander) {
	hookMap := GetHookMap(hookType)
	for _, hook := range hooks {
		if theHook, ok := hookMap[hook]; ok {
			hookCommands[hook] = theHook
		} else {
			panic("couldn't find key:" + hook + " in hook map")
		}
	}
}

func init() {
	//register prehooks
	RegisterHook(PREHOOK, "compress", NewCompressPrehook())
	RegisterHook(POSTHOOK, "archive", NewArchiveCommand())
}
