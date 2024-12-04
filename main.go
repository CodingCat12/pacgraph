package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var csvDir string = filepath.Join("packages", "csv")
var pkgFile string = filepath.Join(csvDir, "packages.json")

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

	RemoveContents(csvDir)
	writeHeaders("packages/csv/packages.csv")
	convertValues(data)
	convertArrays(data)
	logSpecs()
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
