package advent11

import (
	"2021-advent/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleFlash(t *testing.T) {

	data := util.ReadFile("./11_test_input.txt")
	grid := NewGrid(data)
	fmt.Printf("Grid0: %v", grid)
	grid.TriggerStep()
	fmt.Printf("Grid1: %v", grid)
	grid.TriggerStep()
	fmt.Printf("Grid2: %v", grid)
	grid.TriggerStep()
	fmt.Printf("Grid3: %v", grid)
	grid.TriggerStep()
	fmt.Printf("Grid4: %v", grid)
	assert.Equal(t, 0, grid[0][4].energy)
	assert.Equal(t, 3, grid[0][5].energy)
}

func TestRunSteps(t *testing.T) {
	data := util.ReadFile("./11_test_input.txt")
	grid := NewGrid(data)
	total := grid.RunSteps(100)
	assert.Equal(t, 1656, total)
}

func TestFindSync(t *testing.T) {
	data := util.ReadFile("./11_test_input.txt")
	grid := NewGrid(data)
	step := grid.FindSync()
	assert.Equal(t, 195, step)
}
