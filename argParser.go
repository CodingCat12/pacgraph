package main

import (
	"flag"
	"path/filepath"
)

var debugMode bool
var batchSize int
var pkgFile string

func argParser() {
	flag.BoolVar(&debugMode, "debug", false, "Enable debug mode")
	flag.IntVar(&batchSize, "batchsize", 5000, "How many rows to write at once (default: 5000)")
	flag.StringVar(&pkgFile, "outfile", filepath.Join(csvDir, "packages.csv"), "The name of the output file")

	flag.Parse()
}
