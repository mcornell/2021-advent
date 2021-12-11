package advent11

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "11_puzzle.txt")
	data := util.ReadFile(puzzlePath)

	grid := NewGrid(data)
	total := grid.RunSteps(100)
	fmt.Printf("total: %d\n", total)

	grid = NewGrid(data)
	step := grid.FindSync()
	fmt.Printf("step: %d\n", step)
}
