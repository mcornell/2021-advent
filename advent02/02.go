package advent02

import (
	"2021-advent/util"
	"strconv"
	"strings"
)

func ProcessInstruction(instruction string, x_position int, y_position int) (int, int) {
	move := strings.Split(instruction, " ")
	amount, _ := strconv.Atoi(move[1])
	if move[0] == "forward" {
		x_position += amount
	} else if move[0] == "down" {
		y_position += amount
	} else {
		y_position -= amount
	}
	return x_position, y_position

}

func ProcessInstructionTwo(instruction string, x_position int, y_position int, aim_position int) (int, int, int) {
	move := strings.Split(instruction, " ")
	amount, _ := strconv.Atoi(move[1])
	if move[0] == "forward" {
		x_position += amount
		y_position += amount * aim_position
	} else if move[0] == "down" {
		aim_position += amount
	} else {
		aim_position -= amount
	}
	return x_position, y_position, aim_position

}

func ProcessCourse(path string) (int, int) {
	instructions := util.ReadFile(path)
	x := 0
	y := 0
	for _, instruction := range instructions {
		x, y = ProcessInstruction(instruction, x, y)
	}
	return x, y
}

func ProcessCourseTwo(path string) (int, int, int) {
	instructions := util.ReadFile(path)
	x := 0
	y := 0
	aim := 0
	for _, instruction := range instructions {
		x, y, aim = ProcessInstructionTwo(instruction, x, y, aim)
	}
	return x, y, aim
}
