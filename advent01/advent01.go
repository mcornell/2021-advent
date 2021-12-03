package advent01

import (
	"fmt"
)

func RunPuzzle() {
	parsed := ParseInput(puzzle_input_01)
	count := CountIncrease(parsed)
	fmt.Println(count)
	count = CountIncreaseSlidingWindow(parsed)
	fmt.Println(count)
}
