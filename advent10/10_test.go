package advent10

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLineCorrupted(t *testing.T) {

	data := util.ReadFile("./10_test_input.txt")
	isCorrupted, corrupted := IsLineCorrupted(data[0])
	assert.Equal(t, "", corrupted)
	assert.False(t, isCorrupted)
	isCorrupted, corrupted = IsLineCorrupted(data[2])
	assert.Equal(t, "}", corrupted)
	assert.True(t, isCorrupted)
}

func TestFindScore(t *testing.T) {
	data := util.ReadFile("./10_test_input.txt")
	score := FindScore(data)
	assert.Equal(t, 26397, score)
}

func TestCompleteIncompleteLine(t *testing.T) {
	data := util.ReadFile("./10_test_input.txt")

	completed := CompleteIncompleteLine(data[0])
	assert.Equal(t, "}}]])})]", completed)

}

func TestFindCompletedScore(t *testing.T) {
	data := util.ReadFile("./10_test_input.txt")
	score := FindCompletedScore(data)
	assert.Equal(t, 288957, score)
}
