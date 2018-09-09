package main

import (
    log "github.com/sirupsen/logrus"
)

func main () {
    log.WithFields(log.Fields{
        "qqs": "q1",
    }).Info("Why are binaries different?")
}