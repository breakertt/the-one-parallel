package splitutil

import (
	"fmt"
	"strings"
)

type SettingCtrl struct {
	SceNamePrefix string
	MaxIndex      int
	Content       []string
	Keys          []string
	Settings      map[string]setting
	SavePathBase  string
}

type setting struct {
	key         string
	values      []string
	lineContent string
	lineNum     int
}

var DefaultSetCtrl SettingCtrl

func SetupSetCtrl() error {
	DefaultSetCtrl.MaxIndex = 1

	contentBytes := DefaultSceCtrl.SceFmt.Content
	DefaultSetCtrl.Content = strings.Split(string(contentBytes), "\n")

	DefaultSetCtrl.AnalyzeContent()
	return nil
}

func (c *SettingCtrl) AnalyzeContent() error {
	lineNum := 0
	for _, line := range c.Content {
		if strings.Contains(line, "=") {
			kv := strings.Split(string(line), "=")
			if len(kv) != 2 {
				return fmt.Errorf("Invalid setting: %v", line)
			}

			var tmpSetting setting

			tmpSetting.lineNum = lineNum
			lineNum++

			tmpSetting.key = kv[0]

			value := kv[1]
			if strings.Contains(value, "[") && strings.Contains(value, "]") {
				value = strings.ReplaceAll(value, "[", "")
				value = strings.ReplaceAll(value, "]", "")
				values := strings.Split(value, ";")
				fmt.Println(len(values))
			} else {

			}

		} else {
			continue
		}
	}
	return nil
}

// func (c *SettingCtrl) SetupSceneName() error {
// 	basename := path.Base(DefaultSceCtrl.SceSrc.Path)
// 	name := strings.TrimSuffix(basename, filepath.Ext(basename))
// 	return nil
// }
