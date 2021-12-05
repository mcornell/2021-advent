package advent04

import (
	"2021-advent/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBits(t *testing.T) {
	report := util.ReadFile("./04_test_input.txt")
	numbers, cards := ParseInput(report)
	fmt.Printf("Cards: %v\n", cards)
	assert.Equal(t, 27, len(numbers))
	assert.Equal(t, 3, len(cards))
}

func TestDrawNumber(t *testing.T) {
	report := util.ReadFile("./04_test_input.txt")
	numbers, cards := ParseInput(report)
	drawn, numbers, cards := DrawNumber(numbers, cards)
	assert.Equal(t, 26, len(numbers))
	assert.Equal(t, 4, numbers[0])
	assert.Equal(t, 7, drawn)
	assert.Equal(t, -1, cards[0].Numbers[2][4])
}

func TestCheckWinner(t *testing.T) {
	report := util.ReadFile("./04_test_input.txt")
	numbers, cards := ParseInput(report)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	_, numbers, cards = DrawNumber(numbers, cards)
	drawn, _, cards := DrawNumber(numbers, cards)
	assert.False(t, cards[0].CheckWinner())
	assert.False(t, cards[1].CheckWinner())
	fmt.Printf("Card Winner: %#v", cards[2].Numbers)
	assert.True(t, cards[2].CheckWinner())
	assert.Equal(t, 4512, cards[2].CalculateScore(drawn))
}

func TestCheckColumns(t *testing.T) {
	card_input := []string{"0 -1 2 3 4", "-1 -1 6 7 8", "5 -1 10 11 12", "9 -1 14 15 16", "13 -1 18 19 20"}
	card := NewCard(card_input, 0)
	assert.True(t, card.CheckColumns())
}

func TestRunGame(t *testing.T) {
	game := util.ReadFile("./04_test_input.txt")
	winner_index, score := RunGame(game)
	assert.Equal(t, 4512, score)
	assert.Equal(t, 2, winner_index)
}

func TestFindLoser(t *testing.T) {
	game := util.ReadFile("./04_test_input.txt")
	loser_index, score := FindLoser(game)
	assert.Equal(t, 1924, score)
	assert.Equal(t, 1, loser_index)
}
