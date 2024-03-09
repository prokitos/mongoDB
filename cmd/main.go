package main

import (
	"module/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)
	app.MainServer()

}
