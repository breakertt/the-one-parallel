package util

import (
	"fmt"
	"strings"

	"github.com/imaginebreake/the-one-parallel/config"
)

type SettingCtrl struct {
	Start       int
	End         int
	Name        string
	Content     []string
	Keys        []string
	Settings    map[string]setting
	SavePathDir string
}

type setting struct {
	key     string
	values  []string
	content string
}

var DefaultSetCtrl SettingCtrl

func SetupSetCtrl() error {
	DefaultSetCtrl.Name = config.CurrentConfig.ScenarioName
	DefaultSetCtrl.Settings = make(map[string]setting)
	DefaultSetCtrl.Content = strings.Split(string(DefaultSceCtrl.SceFmt.Content), "\n")
	DefaultSetCtrl.SavePathDir = fmt.Sprintf("%v_configs_tmp", DefaultSetCtrl.Name)
	DefaultSetCtrl.AnalyzeContent()
	return nil
}
