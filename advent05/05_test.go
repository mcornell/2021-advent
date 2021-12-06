package advent05

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessLine(t *testing.T) {
	grid := NewVentGrid()
	grid.ProcessLineStraight("1,1 -> 1,3")

	assert.Equal(t, 0, grid.Intensity[0][1])
	assert.Equal(t, 1, grid.Intensity[1][1])
	assert.Equal(t, 1, grid.Intensity[1][3])
	grid.ProcessLineStraight("9,7 -> 7,7")
	assert.Equal(t, 1, grid.Intensity[9][7])
	assert.Equal(t, 1, grid.Intensity[8][7])

}

func TestProcessVentsStraight(t *testing.T) {
	grid := NewVentGrid()
	vents := util.ReadFile("./05_test_input.txt")
	grid.ProcessVentsStraight(vents)

	assert.Equal(t, 2, grid.Intensity[0][9])
	assert.Equal(t, 2, grid.Intensity[2][9])
	assert.Equal(t, 1, grid.Intensity[3][9])
	assert.Equal(t, 1, grid.Intensity[5][9])
	assert.Equal(t, 0, grid.Intensity[6][9])
	assert.Equal(t, 0, grid.Intensity[0][4])
	assert.Equal(t, 1, grid.Intensity[1][4])
	assert.Equal(t, 2, grid.Intensity[3][4])
}

func TestProcessVentsAll(t *testing.T) {
	grid := NewVentGrid()
	vents := util.ReadFile("./05_test_input.txt")
	grid.ProcessVentsAll(vents)

	assert.Equal(t, 2, grid.Intensity[0][9])
	assert.Equal(t, 2, grid.Intensity[2][9])
	assert.Equal(t, 1, grid.Intensity[3][9])
	assert.Equal(t, 1, grid.Intensity[5][9])
	assert.Equal(t, 0, grid.Intensity[6][9])
	assert.Equal(t, 0, grid.Intensity[0][4])
	assert.Equal(t, 1, grid.Intensity[1][4])
	assert.Equal(t, 2, grid.Intensity[3][4])
	assert.Equal(t, 3, grid.Intensity[4][4])
	assert.Equal(t, 1, grid.Intensity[0][0])
}

func TestAvoidTotal(t *testing.T) {
	grid := NewVentGrid()
	vents := util.ReadFile("./05_test_input.txt")
	grid.ProcessVentsStraight(vents)
	assert.Equal(t, 5, grid.AvoidTotal())
	grid2 := NewVentGrid()
	grid2.ProcessVentsAll(vents)
	assert.Equal(t, 12, grid2.AvoidTotal())
}
