package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var jsonDir string = filepath.Join("packages", "json")
var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

func main() {
	argParser()

	jsonFiles, err := os.ReadDir(jsonDir)
	if err != nil {
		log.Fatalf("error opening directory: %v\n", jsonDir)
	}

	convertValues(jsonFiles)
	convertArrays(jsonFiles)

	if debugMode {
		logSpecs()
	}
}
