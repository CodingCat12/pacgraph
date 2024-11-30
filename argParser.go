package main

import (
	"flag"
)

var debugMode bool
var batchSize int

func argParser() {
	flag.BoolVar(&debugMode, "debug", false, "Enable debug mode")
	flag.IntVar(&batchSize, "batchsize", 2000, "How many lines to write at once")

	flag.Parse()
}
