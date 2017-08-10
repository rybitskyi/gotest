package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/rybitskyi/gotest-example/service"
)

func main() {
	log.Infof("Say: " + service.SayHello("Nick"));
}