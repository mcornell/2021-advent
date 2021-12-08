package advent08

import (
	"fmt"
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

func FindConstantNumbers(signal Signal) []string {
	constants := []string{}

	for _, code := range signal.Output {
		if len(code) == 2 || len(code) == 3 || len(code) == 4 || len(code) == 7 {
			constants = append(constants, code)
		}
	}
	return constants
}

func FindAllConstantsInInput(signals []Signal) []string {
	constants := []string{}
	for _, signal := range signals {
		constants = append(constants, FindConstantNumbers(signal)...)
		fmt.Println(constants)
	}
	return constants
}
