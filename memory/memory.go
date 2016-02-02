package memory

import (
	"runtime"
	"time"
)

var (
	diffTime             float64
	lastSampleTime       time.Time         = time.Now()
	lastPauseNs          uint64            = 0
	lastNumGc            uint32            = 0
	memStats             *runtime.MemStats = &runtime.MemStats{}
	now                  time.Time
	nsInMs               float64 = float64(time.Millisecond)
	nbGc                 uint32  = 0
	pauseSinceLastSample uint64  = 0
)

// Returns memory usage statistics
func GetRuntimeStats() (result map[string]float64) {
	runtime.ReadMemStats(memStats)

	now = time.Now()
	diffTime = now.Sub(lastSampleTime).Seconds()

	result = map[string]float64{
		"alloc":          float64(memStats.Alloc),
		"frees":          float64(memStats.Frees),
		"gc.pause_total": float64(memStats.PauseTotalNs) / nsInMs,
		"heap.alloc":     float64(memStats.HeapAlloc),
		"heap.objects":   float64(memStats.HeapObjects),
		"mallocs":        float64(memStats.Mallocs),
		"stack":          float64(memStats.StackInuse),
	}

	if lastPauseNs > 0 {
		pauseSinceLastSample = memStats.PauseTotalNs - lastPauseNs
		result["gc.pause_per_second"] = float64(pauseSinceLastSample) / nsInMs / diffTime
	}

	lastPauseNs = memStats.PauseTotalNs

	nbGc = memStats.NumGC - lastNumGc
	if lastNumGc > 0 {
		result["gc.gc_per_second"] = float64(nbGc) / diffTime
	}

	// Collect GC pauses
	if nbGc > 0 {
		if nbGc > 256 {
			nbGc = 256
		}

		var i uint32

		for i = 0; i < nbGc; i++ {
			idx := int((memStats.NumGC-uint32(i))+255) % 256
			pause := float64(memStats.PauseNs[idx])
			result["gc.pause"] = pause / nsInMs
		}
	}

	// Store last values
	lastNumGc = memStats.NumGC
	lastSampleTime = now

	return result
}
