package advent08

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Each entry consists of ten unique signal patterns, a | delimiter,
and finally the four digit output value. Within an entry,
the same wire/segment connections are used (but you don't know what the connections actually are).
The unique signal patterns correspond to the ten different ways the submarine
tries to render a digit using the current wire/segment connections.
Because 7 is the only digit that uses three segments, dab in the above example means
that to render a 7, signal lines d, a, and b are on. Because 4 is the only digit that
 uses four segments, eafb means that to render a 4, signal lines e, a, f, and b are on.
*/
func TestLoadPuzzle(t *testing.T) {
	data := util.ReadFile("./08_test_input.txt")
	signals := ParseInput(data)
	assert.Equal(t, 10, len(signals))
	assert.Equal(t, 10, len(signals[2].Patterns))
	assert.Equal(t, 4, len(signals[2].Output))
}

func TestFindConstantNumbers(t *testing.T) {
	data := util.ReadFile("./08_test_input.txt")
	signals := ParseInput(data)
	one, four, seven, eight := FindConstantNumbers(signals[0])
	assert.Equal(t, "", one)
	assert.Equal(t, "", seven)
	assert.Equal(t, "gcbe", four)
	assert.Equal(t, "", eight)
}
