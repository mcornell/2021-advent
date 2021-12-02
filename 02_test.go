package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {

	parsed := ReadFile("./02_test_input.txt")
	assert.Equal(t, "up 3", parsed[3])

}

func TestProcessInstruction(t *testing.T) {
	x, y := ProcessInstruction("forward 8", 5, 5)
	assert.Equal(t, 13, x)
	assert.Equal(t, 5, y)
	x, y = ProcessInstruction("down 8", 5, 5)
	assert.Equal(t, 5, x)
	assert.Equal(t, 13, y)
	x, y = ProcessInstruction("up 8", 5, 5)
	assert.Equal(t, 5, x)
	assert.Equal(t, -3, y)

}

func TestProcessCourse(t *testing.T) {
	x, y := ProcessCourse("./02_test_input.txt")
	assert.Equal(t, 15, x)
	assert.Equal(t, 10, y)

	x, y = ProcessCourse("./02_puzzle.txt")
	fmt.Println(x * y)
}

func TestProcessInstructionTwo(t *testing.T) {
	x, y, aim := ProcessInstructionTwo("forward 8", 5, 5, 0)
	assert.Equal(t, 13, x)
	assert.Equal(t, 5, y)
	assert.Equal(t, 0, aim)
	x, y, aim = ProcessInstructionTwo("down 8", 5, 5, 0)
	assert.Equal(t, 5, x)
	assert.Equal(t, 5, y)
	assert.Equal(t, 8, aim)
	x, y, aim = ProcessInstructionTwo("up 8", 5, 5, 5)
	assert.Equal(t, 5, x)
	assert.Equal(t, 5, y)
	assert.Equal(t, -3, aim)
	x, y, aim = ProcessInstructionTwo("forward 8", 5, 5, 5)
	assert.Equal(t, 13, x)
	assert.Equal(t, 45, y)
	assert.Equal(t, 5, aim)
}

func TestProcessCourseTwo(t *testing.T) {
	x, y, aim := ProcessCourseTwo("./02_test_input.txt")
	assert.Equal(t, 15, x)
	assert.Equal(t, 60, y)
	assert.Equal(t, 10, aim)

	x, y, aim = ProcessCourseTwo("./02_puzzle.txt")
	fmt.Println(x * y)
	fmt.Println(aim)
}

// forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
// down 5 adds 5 to your aim, resulting in a value of 5.
// forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
// up 3 decreases your aim by 3, resulting in a value of 2.
// down 8 adds 8 to your aim, resulting in a value of 10.
// forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60
