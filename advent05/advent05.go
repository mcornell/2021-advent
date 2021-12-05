package advent05

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
	puzzlePath := filepath.Join(pwd, "..", "05_puzzle.txt")
	fmt.Println(puzzlePath)
	vents := util.ReadFile(puzzlePath)
	grid := NewVentGrid()
	grid.ProcessVents(vents)
	fmt.Printf("Avoid Total: %d", grid.AvoidTotal())
}
