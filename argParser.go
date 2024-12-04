package main

import (
	"flag"
)

var debugMode bool
var batchSize int

func argParser() {
	flag.BoolVar(&debugMode, "debug", false, "Enable debug mode")
	flag.IntVar(&batchSize, "batchsize", 5000, "How many rows to write at once (default: 5000)")

	flag.Parse()
}
