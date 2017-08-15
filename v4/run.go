package main

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/my-docker/v2/container"
)

func Run(tty bool, comArray []string) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New Parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	sendInitCommand(comArray, writePipe)

	parent.Wait()

	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
