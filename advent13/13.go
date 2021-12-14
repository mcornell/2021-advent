package advent13

import (
	"fmt"
	"strconv"
	"strings"
)

type Dot struct {
	x, y int
}
type Instruction struct {
	index  int
	fold_x bool
}

func (i *Instruction) String() string {
	return fmt.Sprintf("Fold along x: %t at index %d", i.fold_x, i.index)
}

type Paper struct {
	Dots         []Dot
	Instructions []Instruction
	grid         [][]bool
}

func NewPaper(data []string) *Paper {
	dots := []Dot{}
	instructions := []Instruction{}
	max_x := 0
	max_y := 0

	for _, line := range data {
		if strings.Contains(line, ",") {
			dot_split := strings.Split(line, ",")
			x, _ := strconv.Atoi(dot_split[0])
			y, _ := strconv.Atoi(dot_split[1])
			dot := Dot{
				x: x,
				y: y,
			}
			dots = append(dots, dot)
			if x > max_x {
				max_x = x
			}
			if y > max_y {
				max_y = y
			}
		} else if strings.Contains(line, "=") {
			fields := strings.Fields(line)
			fold := strings.Split(fields[2], "=")
			fold_x := true
			if strings.Contains(fold[0], "y") {
				fold_x = false
			}
			index, _ := strconv.Atoi(fold[1])
			instruction := Instruction{
				fold_x: fold_x,
				index:  index,
			}
			instructions = append(instructions, instruction)
		}
	}
	var grid = make([][]bool, max_y+1)
	for x := range grid {
		grid[x] = make([]bool, max_x+1)
	}
	for _, dot := range dots {
		grid[dot.y][dot.x] = true
	}

	paper := Paper{
		Instructions: instructions,
		Dots:         dots,
		grid:         grid,
	}
	return &paper
}

func (p *Paper) String() string {
	retval := "grid :\n\n"
	for _, row := range p.grid {
		for _, val := range row {
			if val {
				retval += "#"
			} else {
				retval += "."
			}
		}
		retval += "\n"
	}

	return retval + "\n"
}

func (p *Paper) Fold() string {
	instruction := p.Instructions[0]
	index := instruction.index
	if instruction.fold_x {
		fmt.Println(instruction)
		for fold, x := index+1, index-1; x > -1; fold, x = fold+1, x-1 {

			for row_idx := range p.grid {
				if p.grid[row_idx][fold] {
					p.grid[row_idx][x] = true
				}
			}
		}
		var new_dots = []Dot{}
		var new_grid = make([][]bool, len(p.grid))
		for x := range new_grid {
			new_grid[x] = make([]bool, index)
		}
		for row_idx, row := range new_grid {
			for col_idx := range row {
				if p.grid[row_idx][col_idx] {
					new_grid[row_idx][col_idx] = true
					dot := Dot{
						x: col_idx,
						y: row_idx,
					}
					new_dots = append(new_dots, dot)
				}
			}
		}

		p.grid = new_grid
		p.Dots = new_dots

	} else {
		for fold, y := index+1, index-1; y > -1; fold, y = fold+1, y-1 {
			for col_idx, val := range p.grid[fold] {
				if val {
					p.grid[y][col_idx] = true
					dot := Dot{
						x: col_idx,
						y: y,
					}
					p.Dots = append(p.Dots, dot)
				}
			}
		}
		p.grid = p.grid[:len(p.grid)-index-1]
	}
	// fmt.Println(p)
	p.Instructions = p.Instructions[1:]
	return instruction.String()
}
