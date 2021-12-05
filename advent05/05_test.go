package advent05

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessLine(t *testing.T) {
	grid := NewVentGrid()
	grid.ProcessLine("1,1 -> 1,3")

	assert.Equal(t, 0, grid.Intensity[0][1])
	assert.Equal(t, 1, grid.Intensity[1][1])
	assert.Equal(t, 1, grid.Intensity[1][3])
	grid.ProcessLine("9,7 -> 7,7")
	assert.Equal(t, 1, grid.Intensity[9][7])
	assert.Equal(t, 1, grid.Intensity[8][7])

}

func TestProcessVents(t *testing.T) {
	grid := NewVentGrid()
	vents := util.ReadFile("./05_test_input.txt")
	grid.ProcessVents(vents)

	assert.Equal(t, 2, grid.Intensity[0][9])
	assert.Equal(t, 2, grid.Intensity[2][9])
	assert.Equal(t, 1, grid.Intensity[3][9])
	assert.Equal(t, 1, grid.Intensity[5][9])
	assert.Equal(t, 0, grid.Intensity[6][9])
	assert.Equal(t, 0, grid.Intensity[0][4])
	assert.Equal(t, 1, grid.Intensity[1][4])
	assert.Equal(t, 2, grid.Intensity[3][4])
}

func TestAvoidTotal(t *testing.T) {
	grid := NewVentGrid()
	vents := util.ReadFile("./05_test_input.txt")
	grid.ProcessVents(vents)
	assert.Equal(t, 5, grid.AvoidTotal())
}
