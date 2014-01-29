package conf

import (
	log "github.com/cihub/seelog"
	"github.com/imosquera/uploadthis/util"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
	"path"
)

var Settings UploadthisConfig

var opts struct {
	ConfigPath string `long:"config"    short:"c" description:"config path"`
	AccessKey  string `long:"accesskey" short:"a" description:"AWS access key"`
	SecretKey  string `long:"secretkey" short:"s" description:"AWS secret key"`
	IAMURL     string `long:"imaurl"    short:"i" description:"IAM security credentials URL"`
	Usage      bool   `long:"usage"     short:"u" description:"Print usage"`
}

type MonitorDir struct {
	Path          string
	TimeThreshold int    `yaml:"time_threshold"`
	Bucket        string
	PreHooks      []string
	PostHooks     []string
}

type UploadthisConfig struct {
	Auth struct {
		AccessKey, SecretKey, IAMURL string
	}
	MonitorDirs []MonitorDir
	Logdir      string
}

var configLoader ConfigLoader = &YamlConfigLoader{}
var loggerConfig LoggerConfig = &SeeLogConfig{defaultLogDir: "/var/log"}

//parses command line options into a UploadthisConfig structure
func ParseOpts() {
	//this setups the logger so that it prints file numbers
	flags.Parse(&opts)

	if opts.ConfigPath != "" {
		log.Info("loading config from: ", opts.ConfigPath)
		configLoader.LoadConfig(opts.ConfigPath, &Settings)
	} else {

	}

	if opts.AccessKey != "" && opts.SecretKey != "" {
		Settings.Auth.AccessKey = opts.AccessKey
		Settings.Auth.SecretKey = opts.SecretKey
	}

	if opts.IAMURL != "" {
		Settings.Auth.IAMURL = opts.IAMURL
	}

	loggerConfig.ConfigLogger(Settings.Logdir)
}

type LoggerConfig interface {
	ConfigLogger(string)
}

type SeeLogConfig struct {
	defaultLogDir string
}

func (self *SeeLogConfig) ConfigLogger(settingsLogDir string) {
	var logDir string
	if settingsLogDir != "" {
		logDir = settingsLogDir
	} else {
		logDir = self.defaultLogDir
	}

	if _, err := os.Stat(logDir); err != nil {
		err = os.MkdirAll(logDir, 0665)
		util.LogPanic(err)
	}

	logPath := path.Join(logDir, "uploadthis.log")
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0665)
	util.LogPanic(err)

	newLogger, err := log.LoggerFromWriterWithMinLevel(logFile, log.InfoLvl)
	util.LogPanic(err)

	log.ReplaceLogger(newLogger)
}

type ConfigLoader interface {
	LoadConfig(path string, settings interface{})
}

type YamlConfigLoader struct{}

func (self *YamlConfigLoader) LoadConfig(path string, settings interface{}) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		panic("can't open config file" + err.Error())
	}
	configString, err := ioutil.ReadAll(file)
	err = goyaml.Unmarshal(configString, settings)
	if err != nil {
		panic("can't unmarshal the yaml file" + err.Error())
	}
}
