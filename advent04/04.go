package advent04

import (
	"strconv"
	"strings"
)

type Card struct {
	Numbers [][]int
	Winner  bool
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

func (c *Card) CalculateScore(drawn int) int {
	total := 0
	for _, row := range c.Numbers {
		for _, col := range row {
			if col > -1 {
				total += col
			}
		}
	}
	return total * drawn
}

func NewCard(lines []string, lineIndex int) *Card {
	cardPointer := new(Card)
	cardPointer.Winner = false

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

func ParseInput(game_input []string) ([]int, []Card) {
	stringNumbers := strings.Split(game_input[0], ",")
	var numbers []int
	for _, num := range stringNumbers {
		val, _ := strconv.Atoi(num)
		numbers = append(numbers, val)
	}
	// println("cards")
	cards := []Card{}
	for lineNum := 2; lineNum < len(game_input); lineNum += 6 {
		// fmt.Printf("line num: %v\n", lineNum)
		cards = append(cards, *NewCard(game_input, lineNum))
		// fmt.Printf("Card! %v\n\n", cards[len(cards)-1])
	}
	// println("cards!")
	return numbers, cards
}

func DrawNumber(numbers []int, cards []Card) (int, []int, []Card) {
	num := numbers[0]
	for _, card := range cards {
		card.MarkNumber(num)
	}
	return num, numbers[1:], cards
}

func IsWinner(cards []Card) int {
	win_index := -1
	for idx, card := range cards {
		if !card.Winner && card.CheckWinner() {
			win_index = idx
			cards[idx].Winner = true
			break
		}
	}
	return win_index
}

func FlagWinner(cards []Card) int {
	retval := -1
	for idx, card := range cards {
		if !card.Winner && card.CheckWinner() {
			cards[idx].Winner = true
			retval = idx
		}
	}
	return retval
}

func RunGame(game []string) (int, int) {
	numbers, cards := ParseInput(game)
	var drawn = -2
	var winner = -1
	for {
		drawn, numbers, cards = DrawNumber(numbers, cards)
		winner = IsWinner(cards)
		if winner > -1 {
			break
		}
	}
	return winner, cards[winner].CalculateScore(drawn)
}

func FindLoser(game []string) (int, int) {
	numbers, cards := ParseInput(game)
	drawn := -2
	winner := -1
	for {
		drawn, numbers, cards = DrawNumber(numbers, cards)
		// fmt.Printf("Drawn: %d\n", drawn)
		winner = FlagWinner(cards)
		if winner > -1 {
			// fmt.Printf("winner found\n")
			all_wins := true
			losers := []int{}
			for idx, card := range cards {
				if !card.Winner {
					all_wins = false
					losers = append(losers, idx)
				}
			}
			// fmt.Printf("losers: %v\n", losers)
			if all_wins {
				break
			}
		}
	}

	return winner, cards[winner].CalculateScore(drawn)
}
