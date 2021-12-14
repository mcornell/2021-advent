package advent13

import (
	"2021-advent/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaper(t *testing.T) {

	data := util.ReadFile("./13_test_input.txt")
	paper := NewPaper(data)
	fmt.Println(paper)
	assert.Equal(t, 2, len(paper.Instructions))
	assert.Equal(t, 18, len(paper.Dots))

}

func TestFold(t *testing.T) {

	data := util.ReadFile("./13_test_input.txt")
	paper := NewPaper(data)
	fmt.Println(paper)
	paper.Fold()
	fmt.Println(paper)
	assert.Equal(t, 1, len(paper.Instructions))
	assert.Equal(t, 17, len(paper.Dots))
	paper.Fold()
	fmt.Println(paper)
	assert.Equal(t, 0, len(paper.Instructions))
	assert.Equal(t, 16, len(paper.Dots))

}
