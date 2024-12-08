package main

import (
	"github.com/CodingCat12/pacgraph/pkg/config"
	"github.com/CodingCat12/pacgraph/pkg/data"
	"github.com/CodingCat12/pacgraph/pkg/log"
)

func main() {
	var err error

	config.DefaultConfig, err = config.LoadConfig("config.json")
	if err != nil {
		log.Logger.Warnf("failed to load config file: %v: falling back to defaults", err)
	}

	config.ParseArgs(&config.AdjustedConfig, config.DefaultConfig)
	log.LoggerSetup()

	packages, err := data.GetData()
	if err != nil {
		log.Logger.Fatalf("failed to retrieve package data, %v", err)
	}

	err = data.ConvertValues(packages)
	if err != nil {
		log.Logger.Fatalf("failed to write package data: %v", err)
	}

	err = data.ConvertArrays(packages)
	if err != nil {
		log.Logger.Fatalf("failed to write package array data: %v", err)
	}

	log.Logger.Infof("successfully wrote %v packages", len(packages))
	log.Logger.Infof("output written to directory: %v", config.AdjustedConfig.Paths.CsvDir)

	log.LogSpecs()
}
