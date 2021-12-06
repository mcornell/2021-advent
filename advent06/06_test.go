package advent06

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLanternFish(t *testing.T) {
	fish := NewLanternFish(3, 1)
	newFish := fish.ProcessDay()
	assert.Equal(t, 0, newFish)
	assert.Equal(t, 2, fish.Timer)
	newFish = fish.ProcessDay()
	assert.Equal(t, 0, newFish)
	assert.Equal(t, 1, fish.Timer)
	newFish = fish.ProcessDay()
	assert.Equal(t, 0, newFish)
	assert.Equal(t, 0, fish.Timer)
	newFish = fish.ProcessDay()
	assert.Equal(t, 1, newFish)
	assert.Equal(t, 6, fish.Timer)
}

func TestSetupState(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	assert.Equal(t, 4, len(fish))
	assert.Equal(t, 1, fish[2].Timer)
}

func TestIncrementDay(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	fish = IncrementDay(fish)
	assert.Equal(t, 4, len(fish))
	assert.Equal(t, 0, fish[2].Timer)
	fish = IncrementDay(fish)
	assert.Equal(t, 5, len(fish))
	assert.Equal(t, 0, fish[3].Timer)
	assert.Equal(t, 8, fish[4].Timer)
	fish = IncrementDay(fish)
	fish = IncrementDay(fish)
	assert.Equal(t, 8, fish[6].Timer)
	assert.Equal(t, 2, fish[6].Count)

}

func TestProcessDays(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	fish_count := ProcessDays(fish, 80)
	assert.Equal(t, 5934, fish_count)

	fish = SetupState(initial_state[0])
	fish_count = ProcessDays(fish, 256)
	assert.Equal(t, 26984457539, fish_count)

}
