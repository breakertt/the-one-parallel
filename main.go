package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/imaginebreake/the-one-multi-thread/config"
	"github.com/imaginebreake/the-one-multi-thread/splitutil"
	"github.com/sirupsen/logrus"
)

func init() {
	flag.StringVar(&config.CurrentConfig.Index, "b", "1", "indexs for batch run, can a value or range like 1:6")
	flag.StringVar(&config.CurrentConfig.ScenarioFile, "c", "default_settings.txt", "file for scenario config")
}

func usage() {
	fmt.Fprintf(os.Stdout, "\nUsage %s\n\n", "the-one-multi-thread")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if err := config.CurrentConfig.ValidateParseConfig(); err != nil {
		logrus.Fatal(err)
	}

	splitutil.DefaultSceCtl.SceSrc = splitutil.Scenario{
		Start: config.CurrentConfig.IndexRange.StartIndex,
		End:   config.CurrentConfig.IndexRange.EndIndex,
		Path:  config.CurrentConfig.ScenarioFile,
	}

	if err := splitutil.GenScenarios(); err != nil {
		logrus.Fatal(err)
	}
}
