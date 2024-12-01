package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var jsonDir string = filepath.Join("packages", "json")
var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

var logger = logrus.New()

func main() {
	argParser()
	if debugMode {
		logger.SetLevel(logrus.DebugLevel)
	}

	jsonFiles, err := os.ReadDir(jsonDir)
	if err != nil {
		logger.Fatalf("error opening directory: %v\n", jsonDir)
	}

	convertValues(jsonFiles)
	convertArrays(jsonFiles)
	logSpecs()
}
