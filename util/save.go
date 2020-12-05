package util

import (
	"fmt"
	"os"
	"time"

	"path"
	"strconv"

	"github.com/sirupsen/logrus"
)

func GenScenarios() error {
	for i := DefaultSetCtrl.Start; i <= DefaultSetCtrl.End; i++ {
		sce := Scenario{
			Path:    DefaultSetCtrl.Name + "_" + strconv.Itoa(i) + "_" + strconv.FormatInt(time.Now().Unix(), 10), // base name
			Content: make([]byte, 0),
		}
		for _, key := range DefaultSetCtrl.Keys {
			valuesLen := len(DefaultSetCtrl.Settings[key].values)
			valueIdx := (i - 1) % valuesLen
			value := DefaultSetCtrl.Settings[key].values[valueIdx]
			if valuesLen != 1 {
				sce.Path += "_" + key + "-" + value
			}
			line := fmt.Sprintf("%v = %v\n", key, value)
			sce.Content = append(sce.Content, []byte(line)...)
		}
		line := fmt.Sprintf("%v = %v\n", "Scenario.name", sce.Path)
		sce.Content = append([]byte(line), sce.Content...)
		sce.Path += ".txt"
		sce.Path = path.Join(DefaultSetCtrl.SavePathDir, sce.Path)
		DefaultSceCtrl.SceOutList = append(DefaultSceCtrl.SceOutList, sce)
	}
	return nil
}

func SaveScenarios() error {
	if _, err := os.Stat(DefaultSetCtrl.SavePathDir); err != nil {
		if err := os.MkdirAll(DefaultSetCtrl.SavePathDir, os.ModePerm); err != nil {
			return err
		}
	}
	for idx, sce := range DefaultSceCtrl.SceOutList {
		logrus.Info(fmt.Sprintf("Saving scenario %v file at %v", idx+1, sce.Path))
		if err := sce.saveContent(); err != nil {
			return fmt.Errorf("Scenario saving error: %v", err)
		}
	}
	return nil
}
