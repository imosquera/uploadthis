package uploadthis

import (
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

var Settings *UploadthisConfig

var opts struct {
	ConfigPath string `short:"c" long:"config" description:"config path"`
	AccesssKey string `long:"accesskey" short:"a" description:"aws access key"`
	SecretKey  string `long:"secretkey" short:"s" description:"Call phone number"`
	Usage      bool   `long:"usage" short:"u" description:"Print usage"`
}

type UploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey string
	}
	WatchFile string
}

func ParseOpts() {
	flags.Parse(&opts)
	if opts.ConfigPath != "" {
		loadConfig(opts.ConfigPath)
	}
	if opts.AccesssKey != "" {
		Settings.Auth.AccessKey = opts.AccesssKey
	}
	if opts.SecretKey != "" {
		Settings.Auth.SecretKey = opts.SecretKey
	}
}
func loadConfig(path string) {
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
