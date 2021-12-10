package advent09

import (
	"2021-advent/util"
	"fmt"
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
	IsLow  bool
}

func NewCell(height int) *Cell {
	return &Cell{
		height: height,
		IsLow:  false,
	}
}

func (grid *Grid) GetCellAt(row, col int) Cell {
	return grid.field[row][col]
}

func (grid *Grid) GetNeighbors(row, col int) []Cell {
	neighbors := []Cell{}
	// Check Above
	above := row - 1
	if above > -1 {
		neighbors = append(neighbors, grid.GetCellAt(above, col))
	}
	// Check Below
	below := row + 1
	if below < grid.row {
		neighbors = append(neighbors, grid.GetCellAt(below, col))
	}
	// Check Left
	left := col - 1
	if left > -1 {
		neighbors = append(neighbors, grid.GetCellAt(row, left))
	}
	// Check Right
	right := col + 1
	if right < grid.col {
		neighbors = append(neighbors, grid.GetCellAt(row, right))
	}

	// for neighbor_row := -1; neighbor_row <= 1; neighbor_row++ {
	// 	cell_row := neighbor_row + row
	// 	if cell_row > -1 && cell_row < grid.row {
	// 		fmt.Printf("Neighbor Row: %v\n", grid.field[cell_row])
	// 		// for neighbor_x := -1; neighbor_x <= 1; neighbor_x++ {
	// 		// 	if neighbor_x == 0 && neighbor_row == 0 {
	// 		// 		continue
	// 		// 	}
	// 		// 	cell_col := neighbor_x + col
	// 		// 	if cell_col > -1 && cell_col < grid.col {
	// 		// 		fmt.Printf("Neighbor Found: %d, %d, %v\n", cell_row, cell_col, grid.GetCellAt(cell_col, cell_row))
	// 		// 		neighbors = append(neighbors, grid.GetCellAt(cell_col, cell_row))
	// 		// 	}

	// 		// }
	// 	}
	// }
	return neighbors
}

func (grid *Grid) FindLowPoints() []Cell {
	lowPoints := []Cell{}
	for row := 0; row < len(grid.field); row++ {
		for col := 0; col < len(grid.field[row]); col++ {
			neighbors := grid.GetNeighbors(row, col)
			theCell := grid.GetCellAt(row, col)
			height_array := []int{}
			for _, neighbor := range neighbors {
				height_array = append(height_array, neighbor.height)
			}
			min, _ := util.MinMax(height_array)
			if theCell.height < min {
				theCell.IsLow = true
				fmt.Printf("Low Point is: %v at %d, %d\n", theCell, row, col)
				lowPoints = append(lowPoints, theCell)
			}
		}
	}
	return lowPoints
}

func FindTotalRisk(data []string) int {
	grid := NewGrid(data)
	lowPoints := grid.FindLowPoints()
	risk := 0
	for _, cell := range lowPoints {
		risk += cell.height + 1
	}
	return risk
}
