package main

import (
	"encoding/json"
	"flag"
	"os"
)

var defaultConfig Config
var adjustedConfig Config

func loadConfig(configFilePath string) (Config, error) {
	configFile, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	var config Config

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func parseArgs() Config {
	result := Config{}

	flag.BoolVar(&result.DebugMode, "debug", defaultConfig.DebugMode, "Enable debug mode")
	flag.IntVar(&result.BatchSize, "batchsize", defaultConfig.BatchSize, "How many rows to write at once (default: 5000)")
	flag.StringVar(&result.Paths.PackageFile, "packagesfile", defaultConfig.Paths.PackageFile, "")
	flag.StringVar(&result.Paths.GroupsFile, "groupsfile", defaultConfig.Paths.GroupsFile, "")
	flag.StringVar(&result.Paths.LicensesFile, "licensesfile", defaultConfig.Paths.LicensesFile, "")
	flag.StringVar(&result.Paths.ConflictsFile, "conflictsfile", defaultConfig.Paths.ConflictsFile, "")
	flag.StringVar(&result.Paths.ProvidesFile, "providesfile", defaultConfig.Paths.ProvidesFile, "Path to the provides CSV file")
	flag.StringVar(&result.Paths.ReplacesFile, "replacesfile", defaultConfig.Paths.ReplacesFile, "Path to the replaces CSV file")
	flag.StringVar(&result.Paths.DependsFile, "dependsfile", defaultConfig.Paths.DependsFile, "Path to the depends CSV file")
	flag.StringVar(&result.Paths.OptDependsFile, "optdependsfile", defaultConfig.Paths.OptDependsFile, "Path to the optional depends CSV file")
	flag.StringVar(&result.Paths.MakeDependsFile, "makedependsfile", defaultConfig.Paths.MakeDependsFile, "Path to the make depends CSV file")
	flag.StringVar(&result.Paths.CheckDependsFile, "checkdependsfile", defaultConfig.Paths.CheckDependsFile, "Path to the check depends CSV file")

	flag.Parse()

	return result
}

type Config struct {
	DebugMode bool `json:"debugMode"`
	BatchSize int  `json:"batchSize"`
	Paths     struct {
		PackageFile      string `json:"packageFile"`
		GroupsFile       string `json:"groupsFile"`
		LicensesFile     string `json:"licensesFile"`
		ConflictsFile    string `json:"conflictsFile"`
		ProvidesFile     string `json:"providesFile"`
		ReplacesFile     string `json:"replacesFile"`
		DependsFile      string `json:"dependsFile"`
		OptDependsFile   string `json:"optDependsFile"`
		MakeDependsFile  string `json:"makeDependsFile"`
		CheckDependsFile string `json:"checkDependsFile"`
	} `json:"paths"`
}
