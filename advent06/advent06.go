package advent06

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	println(pwd)
	puzzlePath := filepath.Join(pwd, "..", "06_puzzle.txt")
	fmt.Println(puzzlePath)
	initial_state := util.ReadFile(puzzlePath)
	fish := SetupState(initial_state[0])

	fish_count := ProcessDays(fish, 80)
	fmt.Printf("Count: %d\n", fish_count)
	fish = SetupState(initial_state[0])

	fish_count = ProcessDays(fish, 256)
	fmt.Printf("Count: %d\n", fish_count)

}

// Alloc = 1059 MiB        TotalAlloc = 7419 MiB   Sys = 2429 MiB  NumGC = 52
// Increment Day
// Alloc = 1011 MiB        TotalAlloc = 7678 MiB   Sys = 2429 MiB  NumGC = 52
// Increment Day
// Alloc = 1928 MiB        TotalAlloc = 9124 MiB   Sys = 3404 MiB  NumGC = 54
// Increment Day
// Alloc = 2053 MiB        TotalAlloc = 9382 MiB   Sys = 3404 MiB  NumGC = 54
// signal: killed
