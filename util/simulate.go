package util

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"sync"

	"github.com/imaginebreake/the-one-parallel/config"
	"github.com/sirupsen/logrus"
)

func RunScenarios() error {
	wg := &sync.WaitGroup{}
	limiter := make(chan bool, config.CurrentConfig.MaxParallel)

	for idx, sce := range DefaultSceCtrl.SceOutList {
		wg.Add(1)
		limiter <- true
		go runTheOne(idx, sce, limiter, wg)
	}

	wg.Wait()
	return nil
}

func runTheOne(idx int, sce Scenario, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	logrus.Info(fmt.Sprintf("Simulation %v:%v - %v start", idx+1, len(DefaultSceCtrl.SceOutList), path.Base(sce.Path)))
	cmd := exec.Command("one.bat", "-b", "1", sce.Path)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info(fmt.Sprintf("Simulation %v:%v - %v fin", idx+1, len(DefaultSceCtrl.SceOutList), path.Base(sce.Path)))
	fmt.Println(string(stdout.Bytes()))
	<-limiter
}
