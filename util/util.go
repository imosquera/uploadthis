package util

import (
	"log"
	"os"
	"path"
)

var RootProjectPath string

func init() {
	RootProjectPath = SetupRootProjectPath()
}
func SetupRootProjectPath() string {
	realGopath := os.Getenv("ORIG_GOPATH")
	realGopath = path.Join(realGopath, "src/github.com/imosquera/uploadthis")
	return realGopath
}

func LogPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
