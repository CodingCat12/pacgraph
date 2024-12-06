package log

import (
	"runtime"
	"time"
)

var StartTime time.Time = time.Now()

func LogSpecs() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	Logger.Debugf("Final Alloc = %v MB\n", memStats.Alloc/1024/1024)

	Logger.Debugf("Operation took: %v", time.Since(StartTime))
}
