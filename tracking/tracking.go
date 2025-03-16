package tracking

import (
	"arena"
	"runtime"
	"sync"
	"unsafe"
)

var structHeapUsage sync.Map

func GetAllocations(ar *arena.Arena) *map[string]uint64 {
	// Allocate the map inside the arena
	allocMap := arena.New[map[string]uint64](ar)
	*allocMap = make(map[string]uint64)

	// Populate map with tracked struct allocations
	structHeapUsage.Range(func(key, value interface{}) bool {
		// Arena-allocate string keys to avoid heap usage
		keyArena := arena.New[string](ar)
		*keyArena = key.(string)

		(*allocMap)[*keyArena] = value.(uint64)
		return true
	})

	// Get current memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Store memory stats inside the arena map
	totalKey := arena.New[string](ar)
	*totalKey = "TotalMemoryUsed"
	(*allocMap)[*totalKey] = memStats.Alloc

	return allocMap
}

func TrackHeapAlloc(structName string, obj any) {
	size := int(unsafe.Sizeof(obj)) // Only stack size

	// Get current heap allocation
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	heapUsed := m.HeapAlloc

	// Store struct heap usage
	structHeapUsage.Store(structName, heapUsed)

	// Debug log
	println("[Tracking] " + structName + " allocated " + string(size) + " bytes")
}
