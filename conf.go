package uploadthis

import (
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

var Settings *UploadthisConfig

type UploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey string
	}
	WatchFile string
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
