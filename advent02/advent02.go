package advent02

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	println(pwd)
	puzzlePath := filepath.Join(pwd, "..", "02_puzzle.txt")
	fmt.Println(puzzlePath)
	x, y := ProcessCourse(puzzlePath)
	fmt.Println(x * y)

	x, y, aim := ProcessCourseTwo(puzzlePath)
	fmt.Println(x * y)
	fmt.Println(aim)

}
