package advent06

import (
	"strconv"
	"strings"
)

type LanternFish struct {
	Timer int
	Count int
}

func NewLanternFish(timer int, count int) *LanternFish {
	fish := new(LanternFish)
	fish.Timer = timer
	fish.Count = count

	return fish
}

func (fish *LanternFish) ProcessDay() int {
	if fish.Timer == 0 {
		fish.Timer = 6
		return fish.Count
	} else {
		fish.Timer--
	}
	return 0
}

func SetupState(state_string string) []LanternFish {
	fish := []LanternFish{}
	initial_times := strings.Split(state_string, ",")
	fish_count := make(map[int]int)
	for _, timer := range initial_times {
		timer_int, _ := strconv.Atoi(timer)
		fish_count[timer_int] += 1
	}
	for timer, count := range fish_count {
		fish = append(fish, *NewLanternFish(timer, count))
	}
	return fish
}

func IncrementDay(fish []LanternFish) []LanternFish {
	newFish := 0
	for idx, a_fish := range fish {
		newFish += a_fish.ProcessDay()
		fish[idx] = a_fish
	}
	if newFish > 0 {
		fish = append(fish, *NewLanternFish(8, newFish))
	}
	return fish
}

func ProcessDays(fish []LanternFish, days int) int {
	for i := 0; i < days; i++ {
		fish = IncrementDay(fish)
		fish_count := make(map[int]int)
		for _, fish_group := range fish {
			fish_count[fish_group.Timer] += fish_group.Count
		}
		new_fish_array := []LanternFish{}
		for timer, count := range fish_count {
			new_fish_array = append(new_fish_array, *NewLanternFish(timer, count))
		}
		fish = new_fish_array
	}
	totalFish := 0
	for _, a_fish := range fish {
		totalFish += a_fish.Count
	}

	return totalFish
}
