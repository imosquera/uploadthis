package hooks

import (
	"log"
	"os"
	"path"
)

func MakeDirWithWarning(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Println(err)
	}
}
func MakeWorkDirs(dirPath string) {
	doing := path.Join(dirPath, "doing")
	done := path.Join(dirPath, "done")
	MakeDirWithWarning(doing)
	MakeDirWithWarning(done)
}
