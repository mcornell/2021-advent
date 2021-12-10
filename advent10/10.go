package advent10

import (
	"2021-advent/util"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	index := len(*s) - 1
	popped := (*s)[index]
	*s = (*s)[:index]
	return popped
}

func IsLineCorrupted(line string) (bool, string) {
	line_stack := new(Stack)
	corrupted := ""
	isCorrupted := false
	chunks := map[string]string{
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
	}
	opens := make([]string, len(chunks))
	i := 0
	for key := range chunks {
		opens[i] = key
		i++
	}

	line_chars := strings.Split(line, "")
	// fmt.Printf("Line Stack: %#v\n", line_stack)

	for _, it := range line_chars {
		// fmt.Printf("next char %s\n", it)
		top := line_stack.Peek()
		if util.Contains(opens, it) || top == "" {
			line_stack.Push(it)
		} else {
			if chunks[top] == it {
				line_stack.Pop()
			} else {
				isCorrupted = true
				corrupted = it
				break
			}
		}
		// fmt.Printf("Line Stack: %v\n", line_stack)
	}
	return isCorrupted, corrupted
}

func FindScore(data []string) int {
	illegal_chars := map[string]int{
		")": 3,
		">": 25137,
		"]": 57,
		"}": 1197,
	}

	score := 0
	for _, line := range data {
		isCorrupted, corrupted := IsLineCorrupted(line)
		if isCorrupted {
			score += illegal_chars[corrupted]
		}
	}
	return score
}
