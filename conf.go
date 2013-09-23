package uploadthis

import (
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

var Settings uploadthisConfig

var opts struct {
	ConfigPath string `short:"c" long:"config" description:"config path"`
	AccesssKey string `long:"accesskey" short:"a" description:"aws access key"`
	SecretKey  string `long:"secretkey" short:"s" description:"Call phone number"`
	Usage      bool   `long:"usage" short:"u" description:"Print usage"`
}

type monitorDir struct {
	Path      string
	Bucket    string
	PreHooks  []string
	PostHooks []string
}

type uploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey string
	}
	DoingDir    string
	DoneDir     string
	MonitorDirs []monitorDir
}

//this is here for mocking purposes
var optsParser = flags.Parse

func ParseOpts() {
	optsParser(&opts)

	if opts.ConfigPath != "" {
		loadConfig(opts.ConfigPath)
	}

	if opts.AccesssKey != "" && opts.SecretKey != "" {
		Settings.Auth.AccessKey = opts.AccesssKey
		Settings.Auth.SecretKey = opts.SecretKey
	}

}

func loadConfig(path string) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	configString, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	goyaml.Unmarshal(configString, &Settings)
}
