package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	return strings.Split(str, "\n")

}

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

func ProcessCourse(path string) (int, int) {
	instructions := ReadFile(path)
	x := 0
	y := 0
	for _, instruction := range instructions {
		x, y = ProcessInstruction(instruction, x, y)
		fmt.Println(fmt.Sprint(x, y))
	}
	return x, y
}

// func main() {
// 	fmt.Println("ho ho ho")
// 	x, y := ProcessCourse("./02_puzzle.txt")
// 	fmt.Println(x * y)

// }
