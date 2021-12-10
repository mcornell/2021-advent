package advent09

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Each number corresponds to the height of a particular location,
where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points -
the locations that are lower than any of its adjacent locations.
Most locations have four adjacent locations (up, down, left, and right);
locations on the edge or corner of the map have three or two adjacent
locations, respectively. (Diagonal locations do not count as adjacent.)
*/
func TestLoadPuzzle(t *testing.T) {
	data := util.ReadFile("./09_test_input.txt")
	grid := NewGrid(data)
	assert.Equal(t, 5, grid.row)
	assert.Equal(t, 10, grid.col)
	assert.Equal(t, 5, len(grid.field))
	assert.Equal(t, 10, len(grid.field[0]))
}
