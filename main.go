package main

import (
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

var logger = logrus.New()

func main() {
	argParser()
	if debugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.FatalLevel)
	}

	data, err := getData()
	if err != nil {
		logger.Fatalf("failed to retrieve package data, %v", err)
	}

	writeHeaders("packages/csv/packages.csv")
	convertValues(data)
	convertArrays(data)
	logSpecs()
}
