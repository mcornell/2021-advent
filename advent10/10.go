package advent10

import (
	"2021-advent/util"
	"sort"
	"strings"
)

var Chunks = map[string]string{
	"(": ")",
	"<": ">",
	"[": "]",
	"{": "}",
}

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

func (s *Stack) Empty() []string {
	reversed_contents := []string{}
	for !s.IsEmpty() {
		reversed_contents = append(reversed_contents, s.Pop())
	}
	return reversed_contents
}

func IsLineCorrupted(line string) (bool, string) {
	line_stack := new(Stack)
	corrupted := ""
	isCorrupted := false

	opens := make([]string, len(Chunks))
	i := 0
	for key := range Chunks {
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
			if Chunks[top] == it {
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

func CompleteIncompleteLine(line string) string {
	line_stack := new(Stack)

	opens := make([]string, len(Chunks))
	i := 0
	for key := range Chunks {
		opens[i] = key
		i++
	}

	line_chars := strings.Split(line, "")

	for _, it := range line_chars {
		// fmt.Printf("next char %s\n", it)
		top := line_stack.Peek()
		if util.Contains(opens, it) || top == "" {
			line_stack.Push(it)
		} else {
			line_stack.Pop()
		}

	}
	// fmt.Printf("Line Stack: %v\n", line_stack)
	completed := ""
	for _, it := range line_stack.Empty() {
		completed += Chunks[it]
	}
	return completed
}

func FindCompletedScore(data []string) int {
	completed_chars := map[string]int{
		")": 1,
		">": 4,
		"]": 2,
		"}": 3,
	}

	scores := []int{}
	for _, line := range data {
		isCorrupted, _ := IsLineCorrupted(line)
		if !isCorrupted {
			a_score := 0
			completed := CompleteIncompleteLine(line)
			for _, it := range strings.Split(completed, "") {
				a_score *= 5
				a_score += completed_chars[it]
			}
			scores = append(scores, a_score)
		}
	}
	// fmt.Printf("Scores: %v\n", scores)
	sort.Ints(scores)
	return scores[len(scores)/2]
}
