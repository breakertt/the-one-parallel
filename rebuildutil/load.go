package rebuildutil

import (
	"fmt"
	"strings"

	"github.com/imaginebreake/the-one-parallel/config"
)

type SettingCtrl struct {
	Start        int
	End          int
	Name         string
	Content      []string
	Keys         []string
	Settings     map[string]setting
	SavePathBase string
}

type setting struct {
	key     string
	values  []string
	content string
}

var DefaultSetCtrl SettingCtrl

func LoadInputScenario() error {
	if err := DefaultSceCtrl.SceSrc.readContent(); err != nil {
		return err
	}
	DefaultSceCtrl.SceFmt = DefaultSceCtrl.SceSrc
	DefaultSceCtrl.SceFmt.fmtContent(DefaultSceCtrl.SceSrc.Content)
	if err := SetupSetCtrl(); err != nil {
		return err
	}
	return nil
}

func SetupSetCtrl() error {
	DefaultSetCtrl.Name = config.CurrentConfig.ScenarioName
	DefaultSetCtrl.Settings = make(map[string]setting)
	DefaultSetCtrl.Content = strings.Split(string(DefaultSceCtrl.SceFmt.Content), "\n")
	DefaultSetCtrl.SavePathBase = fmt.Sprintf("%v_configs_tmp", DefaultSetCtrl.Name)
	DefaultSetCtrl.AnalyzeContent()
	return nil
}

func (c *SettingCtrl) AnalyzeContent() error {
	for _, line := range c.Content {
		if strings.Contains(line, "=") {
			kv := strings.Split(string(line), "=")
			if len(kv) != 2 {
				return fmt.Errorf("Invalid setting: %v", line)
			}
			if kv[0] == "Scenario.name" {
				continue
			}
			c.Settings[kv[0]] = setting{
				key:     kv[0],
				values:  SplitValue(kv[1]),
				content: line,
			}
			c.Keys = append(c.Keys, kv[0])
		}
	}
	return nil
}

func SplitValue(value string) []string {
	var values []string
	if strings.Contains(value, "[") && strings.Contains(value, "]") {
		value = strings.ReplaceAll(value, "[", "")
		value = strings.ReplaceAll(value, "]", "")
		valuesTmp := strings.Split(value, ";")

		for _, value := range valuesTmp {
			if value != "" {
				values = append(values, value)
			}
		}
	} else {
		values = append(values, value)
	}
	return values
}
