package advent08

import (
	"sort"
	"strconv"
	"strings"
)

type Signal struct {
	Patterns []string
	Output   []string
}

func NewSignal(pattern_string string, output_string string) *Signal {
	a_signal := new(Signal)
	a_signal.Patterns = strings.Fields(pattern_string)
	a_signal.Output = strings.Fields(output_string)
	return a_signal
}

func ParseInput(data []string) []Signal {
	signals := []Signal{}
	for _, line := range data {
		line_split := strings.Split(line, "|")
		signals = append(signals, *NewSignal(line_split[0], line_split[1]))

	}
	return signals
}

func FindAllConstantsInLine(signal Signal) []string {
	constants := []string{}

	for _, code := range signal.Output {
		if len(code) == 2 || len(code) == 3 || len(code) == 4 || len(code) == 7 {
			constants = append(constants, code)
		}
	}
	return constants
}

func SortSignalValue(value string) string {
	chars := strings.Split(value, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func FindKnownNumbersInLine(signal Signal) (map[int]string, []string) {
	number_mapping := make(map[int]string)
	unknown := []string{}

	for _, value := range signal.Patterns {
		if len(value) == 2 {
			number_mapping[1] = value
		} else if len(value) == 3 {
			number_mapping[7] = value
		} else if len(value) == 4 {
			number_mapping[4] = value
		} else if len(value) == 7 {
			number_mapping[8] = value
		} else {
			unknown = append(unknown, value)
		}
	}

	return number_mapping, unknown

}

func contains(search []string, it string) bool {
	for _, a := range search {
		if it == a {
			return true
		}
	}
	return false
}

func FindUnknownItem(known string, unknown string) []string {
	var known_slice []string = strings.Split(known, "")

	unknown_values := []string{}
	for _, c := range strings.Split(unknown, "") {
		if !contains(known_slice, c) {
			unknown_values = append(unknown_values, c)
		}
	}
	return unknown_values
}

func FindOneCoordinates(one []string, sixAndNine []string) (string, string, string, []string) {
	top_right := ""
	bottom_right := ""
	six := ""
	nine_or_zero := []string{}
	for _, num := range sixAndNine {
		num_slice := strings.Split(num, "")
		if contains(num_slice, one[0]) && contains(num_slice, one[1]) {
			nine_or_zero = append(nine_or_zero, num)
		} else {
			six = num
		}
	}
	if contains(strings.Split(six, ""), one[0]) {
		top_right = one[1]
		bottom_right = one[0]
	} else {
		top_right = one[0]
		bottom_right = one[1]
	}
	return top_right, bottom_right, six, nine_or_zero
}

func FindThreeValue(one []string, unknown []string) (string, []string) {
	three := ""
	fiveOrTwo := []string{}
	for _, val := range unknown {

		if len(val) == 5 {
			chars := strings.Split(val, "")
			if contains(chars, one[0]) && contains(chars, one[1]) {
				three = val
			} else {
				fiveOrTwo = append(fiveOrTwo, val)
			}
		}
	}
	return three, fiveOrTwo
}

func FindSetDifference(subSet string, superSet string) []string {
	difference := []string{}
	subSetMap := make(map[string]bool)
	for _, val := range strings.Split(subSet, "") {
		subSetMap[val] = true
	}
	for _, val := range strings.Split(superSet, "") {
		if !subSetMap[val] {
			difference = append(difference, val)
		}
	}
	return difference
}

func FindIntersectionValue(left []string, right []string) (string, string, string) {
	intersection := ""
	leftoverLeft := ""
	leftoverRight := ""
	leftSet := make(map[string]bool)
	for _, val := range left {
		leftSet[val] = true
	}
	for _, val := range right {
		if leftSet[val] {
			intersection = val
		} else {
			leftoverRight = val
		}
	}
	delete(leftSet, intersection)
	for k := range leftSet {
		leftoverLeft = k
	}

	return intersection, leftoverLeft, leftoverRight
}

func FindNineAndZero(nine_or_zero []string, middle string) (string, string) {
	nine := ""
	zero := ""
	for _, val := range nine_or_zero {
		if contains(strings.Split(val, ""), middle) {
			nine = val
		} else {
			zero = val
		}
	}

	return nine, zero
}

func FindTwoAndFive(five_and_two []string, top_right string) (string, string, string) {
	two := ""
	five := ""
	bottom_left := ""
	for _, val := range five_and_two {
		if contains(strings.Split(val, ""), top_right) {
			two = val
		} else {
			five = val
		}
	}
	_, two_rem, _ := FindIntersectionValue(strings.Split(two, ""), strings.Split(five, ""))
	_, bottom_left, _ = FindIntersectionValue(strings.Split(two_rem, ""), []string{top_right})

	return two, five, bottom_left

}

func DetermineUnknownValues(known map[int]string, unknown []string) map[int]string {
	one := strings.Split(known[1], "")
	six_nine_and_zero := []string{}
	for _, item := range unknown {
		if len(item) == 6 {
			six_nine_and_zero = append(six_nine_and_zero, item)
		}
	}

	top_right, _, six, nine_or_zero := FindOneCoordinates(one, six_nine_and_zero)
	known[6] = six

	_ = FindUnknownItem(known[1], known[7])[0]

	// 3 contains 1 but 5 and 2 do not.
	var five_and_two = []string{}
	known[3], five_and_two = FindThreeValue(one, unknown)

	// Logic 7 is subset of 3. Find the two letters which aren't known.

	middle_and_bottom := FindSetDifference(known[7], known[3])
	// 4 is subset of 1 Find two letters which aren't known

	middle_and_top_left := FindSetDifference(known[1], known[4])
	middle, _, _ := FindIntersectionValue(middle_and_bottom, middle_and_top_left)
	known[9], known[0] = FindNineAndZero(nine_or_zero, middle)

	known[2], known[5], _ = FindTwoAndFive(five_and_two, top_right)

	return known
}

func FindAllConstantsInInput(signals []Signal) []string {
	constants := []string{}
	for _, signal := range signals {
		constants = append(constants, FindAllConstantsInLine(signal)...)
	}
	return constants
}

func DecodeOutput(signal Signal) string {
	constants, remainder := FindKnownNumbersInLine(signal)
	known := DetermineUnknownValues(constants, remainder)
	knownByValue := make(map[string]int)
	for key, value := range known {
		knownByValue[SortSignalValue(value)] = key
	}
	decoded := ""
	for _, value := range signal.Output {
		sorted := SortSignalValue(value)
		decoded += strconv.Itoa(knownByValue[sorted])
	}
	return decoded

}

func SumAllOutput(signals []Signal) int {
	sum := 0
	for _, signal := range signals {
		output := DecodeOutput(signal)
		val, _ := strconv.Atoi(output)
		sum += val

	}
	return sum
}
