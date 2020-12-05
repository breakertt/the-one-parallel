package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/imaginebreake/the-one-parallel/config"
	"github.com/imaginebreake/the-one-parallel/util"
	"github.com/sirupsen/logrus"
)

func init() {
	flag.StringVar(&config.CurrentConfig.Index, "b", "1", "indexs for batch run, can a value or range like 1:6")
	flag.StringVar(&config.CurrentConfig.ScenarioFile, "c", "default_settings.txt", "file for scenario config")
	flag.StringVar(&config.CurrentConfig.ScenarioName, "n", "default_settings", "name for scenario")
	flag.IntVar(&config.CurrentConfig.MaxParallel, "p", 4, "number of parallel the one simulator")
}

func usage() {
	fmt.Fprintf(os.Stdout, "\nUsage %s\n\n", "the-one-parallel")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if err := config.CurrentConfig.ValidateParseConfig(); err != nil {
		logrus.Fatal(err)
	}

	util.DefaultSetCtrl.Start = config.CurrentConfig.IndexRange.StartIndex
	util.DefaultSetCtrl.End = config.CurrentConfig.IndexRange.EndIndex
	util.DefaultSceCtrl.SceSrc = util.Scenario{
		Path: config.CurrentConfig.ScenarioFile,
	}

	if err := util.LoadInputScenario(); err != nil {
		logrus.Fatal(err)
	}

	if err := util.GenScenarios(); err != nil {
		logrus.Fatal(err)
	}

	if err := util.SaveScenarios(); err != nil {
		logrus.Fatal(err)
	}

	if err := util.RunScenarios(); err != nil {
		logrus.Fatal(err)
	}
}
