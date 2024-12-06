package main

import (
	"github.com/CodingCat12/pacgraph/pkg/config"
	"github.com/CodingCat12/pacgraph/pkg/data"
	"github.com/CodingCat12/pacgraph/pkg/helper"
	"github.com/CodingCat12/pacgraph/pkg/log"
)

var csvDir string = "packages"

func main() {
	var err error

	config.DefaultConfig, err = config.LoadConfig("config.json")
	if err != nil {
		log.Logger.Warnf("failed to load config file, falling back to defaults")
	}

	config.ParseArgs(&config.AdjustedConfig, config.DefaultConfig)
	log.LoggerSetup()

	packages, err := data.GetData()
	if err != nil {
		log.Logger.Fatalf("failed to retrieve package data, %v", err)
	}

	helper.RemoveContents(csvDir)
	data.ConvertValues(packages)
	data.ConvertArrays(packages)
	log.Logger.Debugf("processed %v packages", len(packages))
	log.LogSpecs()
}
