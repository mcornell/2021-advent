package advent14

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {

	data := util.ReadFile("./14_test_input.txt")
	polymer, rules := ParseInput(data)

	assert.Equal(t, 4, len(polymer.p))
	assert.Equal(t, 16, len(rules))

}

func TestExecuteRules(t *testing.T) {

	data := util.ReadFile("./14_test_input.txt")
	polymer, rules := ParseInput(data)
	polymer.ExecuteRules(rules)
	assert.Equal(t, "NCNBCHB", polymer.p)
	polymer.ExecuteRules(rules)
	assert.Equal(t, "NBCCNBBBCBHCB", polymer.p)

}

func TestFindMostCommonElement(t *testing.T) {
	data := util.ReadFile("./14_test_input.txt")
	most, least := FindMostCommonElement(data)
	assert.Equal(t, 161, least)
	assert.Equal(t, 1749, most)
}
