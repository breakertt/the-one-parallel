package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Index      string
	IndexRange struct {
		StartIndex int
		EndIndex   int
	}
	ScenarioFile string
	ScenarioName string
}

var CurrentConfig Config

var (
	ErrorInvalidIndex        = errors.New("Invalid index range")
	ErrorInvalidScenarioFile = errors.New("Invalid config file")
)

func (c *Config) ValidateParseConfig() error {
	if err := c.parseIndex(); err != nil {
		return err
	}

	if err := c.isScenarioFileExist(); err != nil {
		return err
	}

	if err := c.setupName(); err != nil {
		return err
	}

	return nil
}

func (c *Config) parseIndex() error {
	if strings.Contains(c.Index, ":") {
		Idxs := strings.Split(c.Index, ":")
		if len(Idxs) != 2 {
			return ErrorInvalidIndex
		}

		start, err := strconv.Atoi(Idxs[0])
		if err != nil {
			return fmt.Errorf("%v: %v", ErrorInvalidIndex, err)
		}

		end, err := strconv.Atoi(Idxs[1])
		if err != nil {
			return fmt.Errorf("%v: %v", ErrorInvalidIndex, err)
		}

		if start < 1 || start > end {
			return ErrorInvalidIndex
		}

		c.IndexRange.StartIndex, c.IndexRange.EndIndex = start, end
	} else {
		end, err := strconv.Atoi(c.Index)
		if err != nil {
			return fmt.Errorf("%v: %v", ErrorInvalidIndex, err)
		}

		if end < 1 {
			return ErrorInvalidIndex
		}

		c.IndexRange.StartIndex, c.IndexRange.EndIndex = 1, end
	}
	return nil
}

func (c *Config) isScenarioFileExist() error {
	if _, err := os.Stat(c.ScenarioFile); err != nil {
		return fmt.Errorf("%v: %v", ErrorInvalidScenarioFile, err)
	}
	return nil
}

func (c *Config) setupName() error {
	if c.ScenarioName == "default_settings" {
		basename := path.Base(c.ScenarioFile)
		c.ScenarioName = strings.TrimSuffix(basename, filepath.Ext(basename))
	}
	return nil
}
