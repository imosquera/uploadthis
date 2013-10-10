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
	//this is a function that should only be used during testing
	travisBuildDir := os.Getenv("TRAVIS_BUILD_DIR")
	if travisBuildDir != "" {
		return travisBuildDir
	} else {
		realGopath := os.Getenv("ORIG_GOPATH")
		realGopath = path.Join(realGopath, "src/github.com/imosquera/uploadthis")
		return realGopath
	}
}

func LogPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
