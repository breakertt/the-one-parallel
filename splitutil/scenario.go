package splitutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
)

type Scenario struct {
	Start   int
	End     int
	Path    string
	Content []byte
}

type ScenarioCtrl struct {
	SceSrc     Scenario
	SceFmt     Scenario
	SceOutList []Scenario
}

var DefaultSceCtrl ScenarioCtrl

func (s *Scenario) readContent() error {
	var err error
	if s.Content, err = ioutil.ReadFile(s.Path); err != nil {
		return fmt.Errorf("Scnario file %v read error: %v", s.Path, err)
	}
	return nil
}

func (s *Scenario) fmtContent(srcContent []byte) error {
	content := make([]byte, len(srcContent))
	copy(content, srcContent)

	// Delte CRLF newline
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

	// Delete space
	content = bytes.ReplaceAll(content, []byte(" "), []byte(""))

	// Delete tab
	content = bytes.ReplaceAll(content, []byte("\t"), []byte(""))

	// Delete comment lines
	commentRe := regexp.MustCompile(`#.*\n`)
	content = commentRe.ReplaceAll(content, []byte(""))

	// Delte back-slash
	backslashRe := regexp.MustCompile(`\\\n`)
	content = backslashRe.ReplaceAll(content, []byte(""))

	// Delte multiple newline
	newlinesRe := regexp.MustCompile(`\n+`)
	content = newlinesRe.ReplaceAll(content, []byte("\n"))

	s.Content = content

	return nil
}

func (s *Scenario) saveContent() error {
	return nil
}

func LoadInputScenario() error {
	if err := DefaultSceCtrl.SceSrc.readContent(); err != nil {
		return err
	}

	DefaultSceCtrl.SceFmt = DefaultSceCtrl.SceSrc
	DefaultSceCtrl.SceFmt.fmtContent(DefaultSceCtrl.SceSrc.Content)

	// fmt.Print(string(DefaultSceCtrl.SceFmt.Content))
	if err := SetupSetCtrl(); err != nil {
		return err
	}

	return nil
}

func GenScenarios() error {
	return nil
}
