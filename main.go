package main

import (
	"fmt"
)

func main() {
	fmt.Println("ho ho ho")
	parsed := ParseInput(puzzle_input_01)
	count := CountIncrease(parsed)
	fmt.Println(count)
	count = CountIncreaseSlidingWindow(parsed)
	fmt.Println(count)
	fmt.Println("ho ho ho")
	x, y := ProcessCourse("./02_puzzle.txt")
	fmt.Println(x * y)
}
