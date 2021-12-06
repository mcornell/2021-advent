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

func (grid *VentGrid) ProcessLineStraight(line string) {
	points := strings.Split(line, " -> ")
	start := NewCoordinates(points[0])
	end := NewCoordinates(points[1])
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

func (grid *VentGrid) ProcessLineAll(line string) {
	points := strings.Split(line, " -> ")
	start := NewCoordinates(points[0])
	end := NewCoordinates(points[1])
	fmt.Printf("Start: %v,%v End %v,%v\n", start.X, start.Y, end.X, end.Y)

	xIncrement := true
	yIncrement := true
	xStable := false
	yStable := false
	if start.X == end.X {
		xStable = true
		xIncrement = false
	} else if start.X > end.X {
		xIncrement = false
	}
	if start.Y == end.Y {
		yStable = true
		yIncrement = false
	} else if start.Y > end.Y {
		yIncrement = false
	}
	x := start.X
	y := start.Y
	for {
		grid.Intensity[x][y] += 1
		if x == end.X && y == end.Y {
			break
		}
		if xIncrement {
			x++
		} else if !xStable {
			x--
		}
		if yIncrement {
			y++
		} else if !yStable {
			y--
		}
	}

}

func (grid *VentGrid) ProcessVentsStraight(lines []string) {
	for _, line := range lines {
		grid.ProcessLineStraight(line)
	}
}

func (grid *VentGrid) ProcessVentsAll(lines []string) {
	for _, line := range lines {
		grid.ProcessLineAll(line)
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
