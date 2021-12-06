package advent06

import (
	"2021-advent/util"
	"fmt"
	"strconv"
	"strings"
)

type LanternFish struct {
	Timer int
}

func NewLanternFish(timer int) *LanternFish {
	fish := new(LanternFish)
	fish.Timer = timer

	return fish
}

func (fish *LanternFish) ProcessDay() *LanternFish {
	if fish.Timer == 0 {
		fish.Timer = 6
		return NewLanternFish(8)
	} else {
		fish.Timer--
	}
	return nil
}

func SetupState(state_string string) []LanternFish {
	fish := []LanternFish{}
	initial_times := strings.Split(state_string, ",")
	for _, timer := range initial_times {
		timer_int, _ := strconv.Atoi(timer)
		fish = append(fish, *NewLanternFish(timer_int))
	}
	util.PrintMemUsage("SetupState")
	return fish
}

func IncrementDay(fish []LanternFish) []LanternFish {
	for idx, a_fish := range fish {
		newFish := a_fish.ProcessDay()
		fish[idx] = a_fish
		if newFish != nil {
			fish = append(fish, *newFish)
		}
	}
	util.PrintMemUsage("Increment Day")
	return fish
}

func ProcessDays(fish []LanternFish, days int) int {
	for i := 0; i < days; i++ {
		fish = IncrementDay(fish)
	}
	fmt.Println("Done")
	return len(fish)
}
