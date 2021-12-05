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
