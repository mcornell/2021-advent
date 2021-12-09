package advent08

import (
	"fmt"
	"sort"
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
		fmt.Println(line)
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

	fmt.Printf("map: %v\n", number_mapping)
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

func FindOneCoordinates(one []string, sixAndNine []string) (string, string, string, string) {
	top_right := ""
	bottom_right := ""
	six := ""
	nine := ""
	for _, num := range sixAndNine {
		num_slice := strings.Split(num, "")
		if contains(num_slice, one[0]) && contains(num_slice, one[1]) {
			nine = num
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
	return top_right, bottom_right, six, nine
}

func DetermineUnknownValues(known map[int]string, unknown []string) map[int]string {
	one := strings.Split(known[1], "")
	six_and_nine := []string{}
	for _, item := range unknown {
		if len(item) == 6 {
			six_and_nine = append(six_and_nine, item)
		}
	}
	top_right, bottom_right, six, nine := FindOneCoordinates(one, six_and_nine)

	top := FindUnknownItem(known[1], known[7])[0]
	// four_items := FindUnknownItem(known[1], known[4])

	// left_top := FindUnknownItem(known[1], known[4])[0]
	// middle := FindUnknownItem(known[1], known[4])[1]

	fmt.Printf("rt: %s rb: %s t: %s, six: %s, nine: %s\n", top_right, bottom_right, top, six, nine)
	return known
}

func FindAllConstantsInInput(signals []Signal) []string {
	constants := []string{}
	for _, signal := range signals {
		constants = append(constants, FindAllConstantsInLine(signal)...)
		fmt.Println(constants)
	}
	return constants
}
