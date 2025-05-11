package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemStats(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\n[%s]\n", stage)
	fmt.Printf("Allocated: %v KB\n", m.Alloc/1024)
	fmt.Printf("Total Allocated: %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Number of GCs: %v\n", m.NumGC)
}

func main() {
	printMemStats("Start")

	// Allocate a lot of memory to force GC
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 1024) // Allocate 1KB 1M times (~1GB)
	}

	printMemStats("After allocation")

	// Force GC manually
	fmt.Println("\nForcing GC...")
	runtime.GC()

	// Give it a sec to finish
	time.Sleep(1 * time.Second)

	printMemStats("After manual GC")
}
