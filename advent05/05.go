package advent05

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type VentGrid struct {
	Intensity [1000][1000]int
}

func NewVentGrid() *VentGrid {
	grid := new(VentGrid)

	return grid
}

func NewCoordinates(text string) *Coordinates {
	values := strings.Split(text, ",")
	coords := new(Coordinates)
	coords.X, _ = strconv.Atoi(values[0])
	coords.Y, _ = strconv.Atoi(values[1])
	return coords
}

func (grid *VentGrid) ProcessLine(line string) {
	points := strings.Split(line, " -> ")
	start := NewCoordinates(points[0])
	end := NewCoordinates(points[1])
	fmt.Printf("Start: %v End %v\n", start, end)
	var from = 0
	var to = 0
	if start.X == end.X {
		if start.Y < end.Y {
			from = start.Y
			to = end.Y
		} else {
			to = start.Y
			from = end.Y
		}
		for y := from; y < to+1; y++ {
			grid.Intensity[start.X][y] += 1
		}
	} else if start.Y == end.Y {
		if start.X < end.X {
			from = start.X
			to = end.X
		} else {
			to = start.X
			from = end.X
		}

		for x := from; x < to+1; x++ {
			grid.Intensity[x][start.Y] += 1
		}
	}

}

func (grid *VentGrid) ProcessVents(lines []string) {
	for _, line := range lines {
		grid.ProcessLine(line)
	}
}

func (grid *VentGrid) AvoidTotal() int {
	avoidThese := 0
	for _, x := range grid.Intensity {
		for _, val := range x {
			if val > 1 {
				avoidThese++
			}
		}
	}
	return avoidThese
}
