package advent07

import (
	"math"
	"strconv"
	"strings"
)

func LoadPositions(input string) []int {
	list_of_positions := strings.Split(input, ",")
	positions := []int{}
	for _, position_str := range list_of_positions {
		position, _ := strconv.Atoi(position_str)
		positions = append(positions, position)
	}
	return positions
}

func MinMax(int_array []int) (int, int) {
	max := int_array[0]
	min := int_array[0]
	for _, value := range int_array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func BruteForceLeastFuel(positions string, triangular_cost bool) (int, int) {
	crab_subs := LoadPositions(positions)
	min, max := MinMax(crab_subs)
	cheapest_position := 0
	cheapest_cost := 210000000
	for position := min; position < max+1; position++ {
		cost := 0
		for _, sub := range crab_subs {
			sub_cost := int(math.Abs(float64(sub - position)))
			if triangular_cost {
				sub_cost = ((sub_cost + 1) * sub_cost) / 2
			}
			cost += sub_cost
		}
		if cost < cheapest_cost {
			cheapest_position = position
			cheapest_cost = cost
		}
	}
	return cheapest_cost, cheapest_position
}
