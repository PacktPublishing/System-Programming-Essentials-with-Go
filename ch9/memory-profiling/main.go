package main

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	// ...

	f, err := os.Create("memprofile.out")
	if err != nil {
		// Handle error
	}
	defer f.Close()
	runtime.GC()
	pprof.WriteHeapProfile(f)

	// ... (Rest of your code)
}
