package advent04

import (
	"strconv"
	"strings"
)

type Card struct {
	Numbers [][]int
}

func (c *Card) MarkNumber(number int) {
	for row_idx, row := range c.Numbers {
		for col_idx, col := range row {
			if col == number {
				c.Numbers[row_idx][col_idx] = -1
			}
		}
	}
}

func (c *Card) CheckWinner() bool {
	if c.CheckRows() {
		return true
	}
	return c.CheckColumns()
}

func (c *Card) CheckRows() bool {
	retVal := false
	for _, row := range c.Numbers {
		for _, col := range row {
			retVal = true
			if col > -1 {
				retVal = false
				break
			}
		}
		if retVal {
			break
		}
	}
	return retVal
}

func (c *Card) CheckColumns() bool {
	retVal := false
	column_count := len(c.Numbers[0])
	for col := 0; col < column_count; col++ {
		for _, row := range c.Numbers {
			retVal = true
			if row[col] > -1 {
				retVal = false
				break
			}
		}
		if retVal {
			break
		}
	}
	return retVal
}

func NewCard(lines []string, lineIndex int) *Card {
	cardPointer := new(Card)

	// fmt.Printf("Lines Index: %#v\n", lines[lineIndex:lineIndex+5])

	for _, line := range lines[lineIndex : lineIndex+5] {
		// fmt.Printf("Line %d: %v\n", idx, line)
		numStringArray := strings.Fields(line)
		// fmt.Printf("NumStringSlice %d: %v\n", idx, numStringArray)
		var numArray []int
		for _, num := range numStringArray {
			val, _ := strconv.Atoi(num)
			numArray = append(numArray, val)
		}
		// fmt.Printf("Num Array: %#v\n", numArray)
		cardPointer.Numbers = append(cardPointer.Numbers, numArray)
	}
	// fmt.Printf("Numbers: %v\n", cardPointer.Numbers)
	return cardPointer
}

func ParseInput(report []string) ([]int, []Card) {
	stringNumbers := strings.Split(report[0], ",")
	var numbers []int
	for _, num := range stringNumbers {
		val, _ := strconv.Atoi(num)
		numbers = append(numbers, val)
	}
	// println("cards")
	cards := []Card{}
	for lineNum := 2; lineNum < len(report); lineNum += 6 {
		// fmt.Printf("line num: %v\n", lineNum)
		cards = append(cards, *NewCard(report, lineNum))
		// fmt.Printf("Card! %v\n\n", cards[len(cards)-1])
	}
	// println("cards!")
	return numbers, cards
}

func DrawNumber(numbers []int, cards []Card) ([]int, []Card) {
	num := numbers[0]
	for _, card := range cards {
		card.MarkNumber(num)
	}
	return numbers[1:], cards
}
