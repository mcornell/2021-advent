package advent06

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLanternFish(t *testing.T) {
	fish := NewLanternFish(3)
	newFish := fish.ProcessDay()
	assert.Nil(t, newFish)
	assert.Equal(t, 2, fish.Timer)
	newFish = fish.ProcessDay()
	assert.Nil(t, newFish)
	assert.Equal(t, 1, fish.Timer)
	newFish = fish.ProcessDay()
	assert.Nil(t, newFish)
	assert.Equal(t, 0, fish.Timer)
	newFish = fish.ProcessDay()
	assert.NotNil(t, newFish)
	assert.Equal(t, 6, fish.Timer)
	assert.Equal(t, 8, newFish.Timer)
}

func TestSetupState(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	assert.Equal(t, 5, len(fish))
	assert.Equal(t, 1, fish[3].Timer)
}

func TestIncrementDay(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	fish = IncrementDay(fish)
	assert.Equal(t, 5, len(fish))
	assert.Equal(t, 0, fish[3].Timer)
	fish = IncrementDay(fish)
	assert.Equal(t, 6, len(fish))
	assert.Equal(t, 0, fish[4].Timer)
	assert.Equal(t, 8, fish[5].Timer)

}

func TestProcessDays(t *testing.T) {
	initial_state := util.ReadFile("./06_test_input.txt")
	fish := SetupState(initial_state[0])
	fish_count := ProcessDays(fish, 80)
	assert.Equal(t, 5934, fish_count)

	fish = SetupState(initial_state[0])
	fish_count = ProcessDays(fish, 256)
	assert.Equal(t, 5934, fish_count)

}
