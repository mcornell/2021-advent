package advent08

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	puzzlePath := filepath.Join(pwd, "..", "08_puzzle.txt")
	data := util.ReadFile(puzzlePath)
	signals := ParseInput(data)
	constants := FindAllConstantsInInput(signals)

	fmt.Printf("Constants: %v\nLen: %d", constants, len(constants))

	sum := SumAllOutput(signals)
	fmt.Printf("Sum of Output: %d\n", sum)

}
