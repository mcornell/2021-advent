package advent08

import "strings"

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
