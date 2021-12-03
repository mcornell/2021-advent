package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {

	parsed := ReadFile("./test_input.txt")
	assert.Equal(t, "up 3", parsed[3])

}
