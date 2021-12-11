package advent11

import (
	"fmt"
	"strconv"
	"strings"
)

type Octopus struct {
	row, col int
	energy   int
	flashed  bool
}

func NewOctopus(energy, row, col int) *Octopus {
	octo := Octopus{
		row:     row,
		col:     col,
		energy:  energy,
		flashed: false,
	}
	return &octo
}

func (o *Octopus) String() string {
	return fmt.Sprintf("%d, %d: %d - %t", o.row, o.col, o.energy, o.flashed)
}

type Grid [10][10]Octopus

func NewGrid(data []string) *Grid {
	grid := new(Grid)
	for row_idx, row := range data {
		for col_idx, val := range strings.Split(row, "") {
			energy, _ := strconv.Atoi(val)
			grid[row_idx][col_idx] = *NewOctopus(energy, row_idx, col_idx)
		}
	}
	return grid
}

func (g *Grid) String() string {
	printed := "\n"
	for _, row := range g {
		for _, val := range row {
			printed += fmt.Sprint(val.energy)
		}
		printed += "\n"
	}
	return printed
}

func (g *Grid) TriggerStep() int {
	for row_idx, row := range g {
		for col_idx := range row {
			g[row_idx][col_idx].energy += 1
		}
	}

	g.FlashOctopii()
	total_flashed := 0
	for row_idx, row := range g {
		for col_idx := range row {
			if g[row_idx][col_idx].energy > 9 {
				g[row_idx][col_idx].energy = 0
				g[row_idx][col_idx].flashed = false
				total_flashed += 1
			}
		}
	}
	return total_flashed
}

func (g *Grid) FlashOctopii() {
	for row_idx, row := range g {
		for col_idx := range row {
			octo := g[row_idx][col_idx]
			if octo.energy > 9 && !octo.flashed {
				g.Flash(octo)
			}
		}
	}
	flashedAlready := 0
	canFlash := 0
	for row_idx, row := range g {
		for col_idx := range row {
			octo := g[row_idx][col_idx]
			if octo.energy > 9 {
				if octo.flashed {
					flashedAlready += 1
				} else {
					canFlash += 1
				}
			}
		}
	}
	if canFlash > 0 {
		// fmt.Printf("Flashed in this round of this step: %d\n", flashedAlready)
		// fmt.Printf("Needs to Flash in this round of this step: %d\n", canFlash)
		g.FlashOctopii()
	}
}

func (g *Grid) Flash(octo Octopus) {
	g[octo.row][octo.col].flashed = true
	for row := octo.row - 1; row < octo.row+2; row++ {
		for col := octo.col - 1; col < octo.col+2; col++ {
			if row > -1 && row < 10 && col > -1 && col < 10 {
				g[row][col].energy += 1
			}
		}
	}
}

func (g *Grid) RunSteps(stepCount int) int {
	total := 0
	for step := 0; step < stepCount; step++ {
		step_total := g.TriggerStep()
		total += step_total
		// fmt.Printf("Grid after %d - Flashed %d: %v", step, step_total, g)
	}
	return total
}

func (g *Grid) FindSync() int {
	step := 0
	for {
		step_total := g.TriggerStep()
		step++
		// fmt.Printf("Grid after %d - Flashed %d: %v", step, step_total, g)
		if step_total == 100 {
			break
		}

	}
	return step
}
