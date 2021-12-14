package advent13

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "13_puzzle.txt")
	data := util.ReadFile(puzzlePath)
	paper := NewPaper(data)

	paper.Fold()

	fmt.Printf("dot Count: %d\n", len(paper.Dots))
	for len(paper.Instructions) > 0 {
		paper.Fold()
	}
	fmt.Printf("Message: %v\n", paper)
}
