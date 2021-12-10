package advent09

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "09_puzzle.txt")
	data := util.ReadFile(puzzlePath)

	risk := FindTotalRisk(data)
	fmt.Printf("Risk: %d\n", risk)

}
