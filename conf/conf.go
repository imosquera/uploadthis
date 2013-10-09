package conf

import (
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

var Settings UploadthisConfig

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

type UploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey string
	}
	MonitorDirs []MonitorDir
}

var configLoader ConfigLoader = &YamlConfigLoader{}

func ParseOpts() {

	flags.Parse(&opts)

	if opts.ConfigPath != "" {
		configLoader.LoadConfig(opts.ConfigPath, &Settings)
	}

	if opts.AccesssKey != "" && opts.SecretKey != "" {
		Settings.Auth.AccessKey = opts.AccesssKey
		Settings.Auth.SecretKey = opts.SecretKey
	}
}

type ConfigLoader interface {
	LoadConfig(path string, settings interface{})
}

type YamlConfigLoader struct{}

func (self *YamlConfigLoader) LoadConfig(path string, settings interface{}) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Panic("can't open config file", err)
	}
	configString, err := ioutil.ReadAll(file)
	err = goyaml.Unmarshal(configString, settings)
	if err != nil {
		log.Panic("can't unmarshal the yaml file", err)
	}
}
