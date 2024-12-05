package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var csvDir string = "packages"
var pkgFile string = filepath.Join(csvDir, "packages.csv")

var startTime time.Time = time.Now()

var logger = logrus.New()

func main() {
	defaultConfig, _ = loadConfig("config.json")
	parseArgs()
	if adjustedConfig.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.FatalLevel)
	}

	data, err := getData()
	if err != nil {
		logger.Fatalf("failed to retrieve package data, %v", err)
	}

	RemoveContents(csvDir)
	convertValues(data)
	convertArrays(data)
	logSpecs()
}

func RemoveContents(dirName string) error {
	dir, err := os.Open(dirName)
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirName, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func toString(value any) string {
	result := fmt.Sprintf("%v", value)
	return result
}
