package advent07

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPositions(t *testing.T) {
	positions := util.ReadFile("./07_test_input.txt")[0]
	crab_subs := LoadPositions(positions)

	assert.Equal(t, 2, crab_subs[2])
	assert.Equal(t, 10, len(crab_subs))
}

func TestFindMinMax(t *testing.T) {
	positions := util.ReadFile("./07_test_input.txt")[0]
	crab_subs := LoadPositions(positions)

	min, max := MinMax(crab_subs)
	assert.Equal(t, 0, min)
	assert.Equal(t, 16, max)
}

func TestBruteForceLeastFuel(t *testing.T) {
	positions := util.ReadFile("./07_test_input.txt")[0]
	fuel, position := BruteForceLeastFuel(positions, false)

	assert.Equal(t, 37, fuel)
	assert.Equal(t, 2, position)
}

func TestBruteForceLeastFuelTriagleCost(t *testing.T) {
	positions := util.ReadFile("./07_test_input.txt")[0]
	fuel, position := BruteForceLeastFuel(positions, true)

	assert.Equal(t, 168, fuel)
	assert.Equal(t, 5, position)
}
