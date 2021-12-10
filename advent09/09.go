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
	height         int
	IsLow, IsBasin bool
	row, col       int
}

func NewCell(height, row, col int) *Cell {
	return &Cell{
		height:  height,
		IsLow:   false,
		IsBasin: height < 9,
		row:     row,
		col:     col,
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

func (grid *Grid) FindBasin(lowPoint Cell) int {
	// fmt.Printf("Starting at Cell: %v\n", lowPoint)
	basin_size := 1
	start_row := lowPoint.row
	start_col := lowPoint.col
	// to top
	// fmt.Printf("Top\n")
	for row := start_row; row > -1; row-- {
		if grid.GetCellAt(row, start_col).IsBasin {
			basin_size++
		} else {
			break
		}
		// to left
		// fmt.Println("left")
		for col := start_col - 1; col > -1; col-- {
			// fmt.Printf("Checking %d, %d: ", row, col)
			if grid.GetCellAt(row, col).IsBasin {
				basin_size++
				// fmt.Println("Is Basin")
			} else {
				// fmt.Println("Is Wall")
				break
			}
		}
		// to right
		// fmt.Println("right -> will skip starting point")
		for col := start_col + 1; col < grid.col; col++ {
			// fmt.Printf("Checking %d, %d ", row, col)
			if grid.GetCellAt(row, col).IsBasin {
				basin_size++
				// fmt.Println("Is Basin")
			} else {
				// fmt.Println("Is Wall")
				break
			}
		}
	}
	//to bottom
	// fmt.Println("bottom")
	for row := start_row + 1; row < grid.row; row++ {
		if grid.GetCellAt(row, start_col).IsBasin {
			basin_size++
		} else {
			break
		}
		// to left
		// fmt.Println("left")
		for col := start_col - 1; col > -1; col-- {
			// fmt.Printf("Checking %d, %d ", row, col)
			if grid.GetCellAt(row, col).IsBasin {
				basin_size++
				// fmt.Println("Is Basin")
			} else {
				// fmt.Println("Is Wall")
				break
			}
		}
		// to right
		// fmt.Println("right")
		for col := start_col + 1; col < grid.col; col++ {
			// fmt.Printf("Checking %d, %d ", row, col)
			if grid.GetCellAt(row, col).IsBasin {
				basin_size++
				// fmt.Println("Is Basin")
			} else {
				// fmt.Println("Is Wall")
				break
			}
		}
	}
	return basin_size
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

func FindAllBasins(data []string) []int {
	basinSizes := []int{}
	grid := NewGrid(data)
	lowPoints := grid.FindLowPoints()
	for _, lowPoint := range lowPoints {
		basinSizes = append(basinSizes, grid.FindBasin(lowPoint))
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
