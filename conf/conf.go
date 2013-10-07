package conf

import (
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
	"path"
)

var Settings uploadthisConfig

var opts struct {
	ConfigPath string `short:"c" long:"config" description:"config path"`
	AccesssKey string `long:"accesskey" short:"a" description:"aws access key"`
	SecretKey  string `long:"secretkey" short:"s" description:"Call phone number"`
	Usage      bool   `long:"usage" short:"u" description:"Print usage"`
}

type MonitorDir struct {
	Path      string
	Bucket    string
	PreHooks  []string
	PostHooks []string
}

type uploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey string
	}
	MonitorDirs []MonitorDir
}

var optsParser = flags.Parse

var ParseOpts = func() {

	optsParser(&opts)

	if opts.ConfigPath != "" {
		loadConfig(opts.ConfigPath)
	}

	if opts.AccesssKey != "" && opts.SecretKey != "" {
		Settings.Auth.AccessKey = opts.AccesssKey
		Settings.Auth.SecretKey = opts.SecretKey
	}
}

var loadConfig = func(path string) {
	file, err := os.Open(path) // For read access.
	e, _ := os.Getwd()
	println(e)

	if err != nil {
		log.Panic("can't open config file", err)
	}
	configString, err := ioutil.ReadAll(file)
	err = goyaml.Unmarshal(configString, &Settings)
	if err != nil {
		log.Panic("can't unmarshal the yaml file", err)
	}
}

var MakeDirWithWarning = func(dirPath string) {
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Println(err)
	}
}

var MakeWorkDirs = func(dirPath string) {
	doing := path.Join(dirPath, "doing")
	done := path.Join(dirPath, "done")
	MakeDirWithWarning(doing)
	MakeDirWithWarning(done)
}