package advent03

import (
	"2021-advent/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBits(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	ones, zeros := FindBits(report, 0)
	assert.Equal(t, 7, ones)
	assert.Equal(t, 5, zeros)
	ones, zeros = FindBits(report, 2)
	assert.Equal(t, 8, ones)
	assert.Equal(t, 4, zeros)
}

func TestFindGammaEpislonValues(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	gamma, epsilon := FindGammaEpsilonValues(report)
	assert.Equal(t, "10110", gamma)
	assert.Equal(t, "01001", epsilon)
}

func TestConvertGammaEpsilonToNumber(t *testing.T) {
	gamma, epsilon := ConvertGammaEpsilonToNumber("10110", "01001")
	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)
}

func TestFindValue(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	assert.Equal(t, 198, FindValue(report))
}
