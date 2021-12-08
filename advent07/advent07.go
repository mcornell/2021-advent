package advent07

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
	puzzlePath := filepath.Join(pwd, "..", "07_puzzle.txt")
	fmt.Println(puzzlePath)
	positions := util.ReadFile(puzzlePath)
	cost, position := BruteForceLeastFuel(positions[0], false)

	fmt.Printf("Cost: %d\nPostion: %d\n", cost, position)
	cost, position = BruteForceLeastFuel(positions[0], true)
	fmt.Printf("Cost: %d\nPostion: %d\n", cost, position)
}
