package advent09

import (
	"2021-advent/util"
	"fmt"
	"sort"
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
	for row_idx, row := range data {
		cellRow := []Cell{}
		for col_idx, val := range strings.Split(row, "") {
			height, _ := strconv.Atoi(val)
			cellRow = append(cellRow, *NewCell(height, row_idx, col_idx))
		}
		grid.field = append(grid.field, cellRow)
	}

	return &grid
}

type Cell struct {
	height, row, col             int
	IsLow, IsBasin, BasinChecked bool
}

func NewCell(height, row, col int) *Cell {
	return &Cell{
		row:          row,
		col:          col,
		height:       height,
		IsLow:        false,
		IsBasin:      height < 9,
		BasinChecked: false,
	}
}

func (cell *Cell) String() string {
	return fmt.Sprintf("Cell At: %d, %d Height: %d, BasinChecked: %t", cell.row, cell.col, cell.height, cell.BasinChecked)
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
				// fmt.Printf("Low Point is: %v at %d, %d\n", theCell, row, col)
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

func FindBasin(cell Cell, grid *Grid) int {
	basin_size := 0
	if cell.IsBasin && !cell.BasinChecked {
		basin_size = 1
		grid.field[cell.row][cell.col].BasinChecked = true
		neighbors := grid.GetNeighbors(cell.row, cell.col)
		for _, neighbor := range neighbors {
			updated_neighbor := grid.GetCellAt(neighbor.row, neighbor.col)
			if !updated_neighbor.BasinChecked {
				basin_size += FindBasin(updated_neighbor, grid)
			}
		}
	}
	return basin_size
}

func FindAllBasins(data []string) []int {
	basinSizes := []int{}
	grid := NewGrid(data)
	lowPoints := grid.FindLowPoints()
	for _, lowPoint := range lowPoints {
		basinSizes = append(basinSizes, FindBasin(lowPoint, grid))
	}
	return basinSizes
}

func FindTopThreeBasins(data []string) int {
	basins := FindAllBasins(data)
	sort.Ints(basins)

	result := basins[len(basins)-1]
	result *= basins[len(basins)-2]
	result *= basins[len(basins)-3]
	fmt.Printf("top 3: %d, %d, %d result: %d\n", basins[len(basins)-1], basins[len(basins)-2], basins[len(basins)-3], result)
	return result
}
