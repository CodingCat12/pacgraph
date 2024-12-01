package main

import (
	"runtime"
	"time"
)

func logSpecs() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	logger.Debugf("Final Alloc = %v MB\n", memStats.Alloc/1024/1024)

	logger.Debugf("Operation took: %v", time.Since(startTime))
}
