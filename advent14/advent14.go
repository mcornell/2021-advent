package advent14

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "14_puzzle.txt")
	data := util.ReadFile(puzzlePath)
	most, least := FindMostCommonElement(data)

	fmt.Printf("Most: %d - Least: %d = %d\n", most, least, most-least)
}
