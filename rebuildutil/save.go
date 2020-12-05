package rebuildutil

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func GenScenarios() error {
	for i := DefaultSetCtrl.Start; i <= DefaultSetCtrl.End; i++ {
		sce := Scenario{
			Path:    DefaultSetCtrl.Name + "_" + strconv.Itoa(i), // base name
			Content: make([]byte, 2048),
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
		sce.Path = strings.ReplaceAll(sce.Path, ".", "")
		sce.Path += ".txt"
		sce.Path = path.Join(DefaultSetCtrl.SavePathBase, sce.Path)
		DefaultSceCtrl.SceOutList = append(DefaultSceCtrl.SceOutList, sce)
	}
	return nil
}

func SaveScenarios() error {
	for idx, sce := range DefaultSceCtrl.SceOutList {
		logrus.Info(fmt.Sprintf("Saving scenario %v file at %v", idx+1, sce.Path))
	}
	return nil
}
