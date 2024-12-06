package config

import (
	"encoding/json"
	"flag"
	"os"
)

var DefaultConfig Config
var AdjustedConfig Config

func LoadConfig(configFilePath string) (Config, error) {
	fallbackConfig := Config{
		DebugMode: false,
		BatchSize: 5000,
		Paths: struct {
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
		}{
			PackageFile:      "packages/packages.csv",
			GroupsFile:       "packages/groups.csv",
			LicensesFile:     "packages/licenses.csv",
			ConflictsFile:    "packages/conflicts.csv",
			ProvidesFile:     "packages/provides.csv",
			ReplacesFile:     "packages/replaces.csv",
			DependsFile:      "packages/depends.csv",
			OptDependsFile:   "packages/optdepends.csv",
			MakeDependsFile:  "packages/makedepends.csv",
			CheckDependsFile: "packages/checkdepends.csv",
		},
	}

	configFile, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return fallbackConfig, err
	}
	defer configFile.Close()

	var config Config

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return fallbackConfig, err
	}

	return config, nil
}

func ParseArgs(adjustedConf *Config, defaultConf Config) {
	flag.BoolVar(&adjustedConf.DebugMode, "debug", defaultConf.DebugMode, "Enable debug mode")
	flag.IntVar(&adjustedConf.BatchSize, "batchsize", defaultConf.BatchSize, "How many rows to write at once (default: 5000)")
	flag.StringVar(&adjustedConf.Paths.PackageFile, "packagesfile", defaultConf.Paths.PackageFile, "")
	flag.StringVar(&adjustedConf.Paths.GroupsFile, "groupsfile", defaultConf.Paths.GroupsFile, "")
	flag.StringVar(&adjustedConf.Paths.LicensesFile, "licensesfile", defaultConf.Paths.LicensesFile, "")
	flag.StringVar(&adjustedConf.Paths.ConflictsFile, "conflictsfile", defaultConf.Paths.ConflictsFile, "")
	flag.StringVar(&adjustedConf.Paths.ProvidesFile, "providesfile", defaultConf.Paths.ProvidesFile, "Path to the provides CSV file")
	flag.StringVar(&adjustedConf.Paths.ReplacesFile, "replacesfile", defaultConf.Paths.ReplacesFile, "Path to the replaces CSV file")
	flag.StringVar(&adjustedConf.Paths.DependsFile, "dependsfile", defaultConf.Paths.DependsFile, "Path to the depends CSV file")
	flag.StringVar(&adjustedConf.Paths.OptDependsFile, "optdependsfile", defaultConf.Paths.OptDependsFile, "Path to the optional depends CSV file")
	flag.StringVar(&adjustedConf.Paths.MakeDependsFile, "makedependsfile", defaultConf.Paths.MakeDependsFile, "Path to the make depends CSV file")
	flag.StringVar(&adjustedConf.Paths.CheckDependsFile, "checkdependsfile", defaultConf.Paths.CheckDependsFile, "Path to the check depends CSV file")

	flag.Parse()
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
