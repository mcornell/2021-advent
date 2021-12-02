package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input_01 = `199
200
208
210
200
207
240
269
260
263`

func TestParseInput(t *testing.T) {

	parsed := ParseInput(test_input_01)
	assert.Equal(t, 10, len(parsed))
	assert.Equal(t, "260", parsed[8])
}

func TestCountInc(t *testing.T) {

	parsed := ParseInput(test_input_01)
	count := CountIncrease(parsed)
	assert.Equal(t, 7, count)
}
