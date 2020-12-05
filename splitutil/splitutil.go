package splitutil

import (
	"fmt"
	"io/ioutil"
)

type Scenario struct {
	Start      int
	End        int
	Path       string
	Content    []byte
	Name       string
	SavePrefix string
}

type ScenarioCtl struct {
	SceSrc     Scenario
	SceFmt     Scenario
	SceOutList []Scenario
}

var DefaultSceCtl ScenarioCtl

func (s *Scenario) readContent() error {
	var err error
	if s.Content, err = ioutil.ReadFile(s.Path); err != nil {
		return fmt.Errorf("Scnario file %v read error: %v", s.Path, err)
	}
	fmt.Println(s.Content)
	return nil
}

func (s *Scenario) saveContent() error {
	return nil
}

func GenScenarios() error {
	if err := DefaultSceCtl.SceSrc.readContent(); err != nil {
		return err
	}
	fmt.Printf("%v", string(DefaultSceCtl.SceSrc.Content))
	return nil

}
