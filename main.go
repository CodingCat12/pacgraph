package main

import (
	"path/filepath"
	"time"
)

var jsonDir string = filepath.Join("packages", "json")
var csvDir string = filepath.Join("packages", "csv")

var startTime time.Time = time.Now()

func main() {
	argParser()
	convertValues()
	convertArrays()

	if debugMode {
		logSpecs()
	}
}
