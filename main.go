package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/imaginebreake/the-one-parallel/config"
	"github.com/imaginebreake/the-one-parallel/rebuildutil"
	"github.com/sirupsen/logrus"
)

func init() {
	flag.StringVar(&config.CurrentConfig.Index, "b", "1", "indexs for batch run, can a value or range like 1:6")
	flag.StringVar(&config.CurrentConfig.ScenarioFile, "c", "default_settings.txt", "file for scenario config")
	flag.StringVar(&config.CurrentConfig.ScenarioName, "n", "default_settings", "name for scenario")
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

	rebuildutil.DefaultSetCtrl.Start = config.CurrentConfig.IndexRange.StartIndex
	rebuildutil.DefaultSetCtrl.End = config.CurrentConfig.IndexRange.EndIndex
	rebuildutil.DefaultSceCtrl.SceSrc = rebuildutil.Scenario{
		Path: config.CurrentConfig.ScenarioFile,
	}

	if err := rebuildutil.LoadInputScenario(); err != nil {
		logrus.Fatal(err)
	}

	if err := rebuildutil.GenScenarios(); err != nil {
		logrus.Fatal(err)
	}

	if err := rebuildutil.SaveScenarios(); err != nil {
		logrus.Fatal(err)
	}
}
