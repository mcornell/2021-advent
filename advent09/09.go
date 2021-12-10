package advent09

import (
	"strconv"
	"strings"
)

type Grid struct {
	field    [][]Cell
	row, col int
}

func NewGrid(data []string) *Grid {
	grid := Grid{
		row: len(data),
		col: len(data[0]),
	}
	for _, row := range data {
		cellRow := []Cell{}
		for _, val := range strings.Split(row, "") {
			height, _ := strconv.Atoi(val)
			cellRow = append(cellRow, *NewCell(height))
		}
		grid.field = append(grid.field, cellRow)
	}

	return &grid
}

type Cell struct {
	height int
	is_low bool
}

func NewCell(height int) *Cell {
	return &Cell{
		height: height,
		is_low: false,
	}
}

func (grid *Grid) GetNeighbors(x, y int) []Cell {
	neighbors := []Cell{}
	for neighbor_y := -1; neighbor_y <= 1; neighbor_y++ {
		cell_y := neighbor_y + y
		if cell_y > -1 && cell_y < grid.row+1 {
			for neighbor_x := 1; neighbor_x <= 1; neighbor_x++ {
				if neighbor_x == 0 && neighbor_y == 0 {
					continue
				}
				cell_x := neighbor_x + x
				if cell_x > -1 && cell_x < grid.col+1 {
					neighbors = append(neighbors, grid.field[cell_x][cell_y])
				}

			}
		}
	}
	return neighbors
}

func (grid *Grid) FindLowPoints() []Cell {
	lowPoints := []Cell{}
	for y := 0; y < len(grid.field); y++ {
		for x := 0; x < len(grid.field[y]); x++ {
			neighbors := grid.GetNeighbors(x, y)
			is_low := true
			for _, neighbor := range neighbors {
				if neighbor.height < grid.field[x][y].height {
					is_low = false
					break
				}
			}
			grid.field[x][y].is_low = is_low
			lowPoints = append(lowPoints, grid.field[x][y])
		}
	}
	return lowPoints
}
