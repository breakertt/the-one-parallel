package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/imaginebreake/the-one-multi-thread/config"
)

func init() {
	flag.StringVar(&config.CurrentConfig.Index, "b", "1", "indexs for batch run, can a value or range like 1:6")
	flag.StringVar(&config.CurrentConfig.ConfigFile, "c", "default_settings.txt", "file for scenario config")
}

func usage() {
	fmt.Fprintf(os.Stdout, "\nUsage %s\n\n", "the-one-multi-thread")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	err := config.CurrentConfig.ValidateParseConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("%v", config.CurrentConfig)
}
