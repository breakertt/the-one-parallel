package util

import (
	"fmt"
	"strings"
)

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
