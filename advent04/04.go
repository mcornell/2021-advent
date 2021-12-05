package advent04

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Numbers [][]int
}

func NewCard(lines []string, lineIndex int) *Card {
	cardPointer := new(Card)

	fmt.Printf("Lines Index: %#v\n", lines[lineIndex:lineIndex+5])

	for idx, line := range lines[lineIndex : lineIndex+5] {
		fmt.Printf("Line %d: %v\n", idx, line)
		numStringArray := strings.Fields(line)
		fmt.Printf("NumStringSlice %d: %v\n", idx, numStringArray)
		var numArray []int
		for _, num := range numStringArray {
			val, _ := strconv.Atoi(num)
			numArray = append(numArray, val)
		}
		fmt.Printf("Num Array: %#v\n", numArray)
		cardPointer.Numbers = append(cardPointer.Numbers, numArray)
	}
	fmt.Printf("Numbers: %v\n", cardPointer.Numbers)
	return cardPointer
}

func ParseInput(report []string) ([]int, []Card) {
	stringNumbers := strings.Split(report[0], ",")
	var numbers []int
	for _, num := range stringNumbers {
		val, _ := strconv.Atoi(num)
		numbers = append(numbers, val)
	}
	println("cards")
	cards := []Card{}
	for lineNum := 2; lineNum < len(report); lineNum += 6 {
		fmt.Printf("line num: %v\n", lineNum)
		cards = append(cards, *NewCard(report, lineNum))
		fmt.Printf("Card! %v\n\n", cards[len(cards)-1])
	}
	println("cards!")
	return numbers, cards
}
