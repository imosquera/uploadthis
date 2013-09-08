package main

import (
	"flag"
	"fmt"
	"github.com/imosquera/uploadthis"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(3)
}

func main() {
	var (
		configFile string
	)

	flag.StringVar(&configFile, "c", "/etc/samplepath", "path to config file")
	flag.Parse()
	if configFile == "/etc/samplepath" {
		Usage()
	}

	uploadthis.LoadConfig("sample-config.yaml")
	println(uploadthis.Settings.WatchFile)
}
