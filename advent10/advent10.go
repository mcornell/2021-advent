package advent10

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "10_puzzle.txt")
	data := util.ReadFile(puzzlePath)

	score := FindScore(data)
	fmt.Printf("score: %d\n", score)

	score = FindCompletedScore(data)
	fmt.Printf("score: %d\n", score)
}
