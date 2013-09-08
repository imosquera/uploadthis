package uploadthis

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

var Settings *UploadthisConfig

type UploadthisConfig struct {
	WatchDir string
}

func LoadConfig(path string) {
	Settings = new(UploadthisConfig)
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	configString, err := ioutil.ReadAll(file)
	println(string(configString))
	if err != nil {
		log.Fatal(err)
	}
	goyaml.Unmarshal(configString, Settings)
}
