package splitutil

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	return nil
}

func (s *Scenario) fmtContent(srcContent []byte) error {
	// Delete comment lines
	commentRe := regexp.MustCompile(`#.*(\r\n|\n)`)
	content := commentRe.ReplaceAll(srcContent, []byte(""))

	// Delte back-slash
	backslashRe := regexp.MustCompile(`\\( |\t)*(\r\n|\n)( |\t)*`)
	content = backslashRe.ReplaceAll(content, []byte(""))

	fmt.Print(string(content))
	return nil
}

func (s *Scenario) saveContent() error {
	return nil
}

func GenScenarios() error {
	if err := DefaultSceCtl.SceSrc.readContent(); err != nil {
		return err
	}

	DefaultSceCtl.SceFmt = DefaultSceCtl.SceSrc
	DefaultSceCtl.SceFmt.fmtContent(DefaultSceCtl.SceSrc.Content)

	// fmt.Print(string(DefaultSceCtl.SceFmt.Content))
	return nil
}
