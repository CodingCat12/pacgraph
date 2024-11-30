package main

import (
	"log"
	"runtime"
	"time"
)

func logSpecs() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Printf("Final Alloc = %v MB\n", memStats.Alloc/1024/1024)

	log.Printf("Operation took: %v", time.Since(startTime))
}
