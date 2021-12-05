package advent04

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	println(pwd)
	puzzlePath := filepath.Join(pwd, "..", "03_puzzle.txt")
	fmt.Println(puzzlePath)
	// report := util.ReadFile(puzzlePath)

}
