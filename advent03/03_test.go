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

func TestSeparateValuesByBit(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	ones, zeros := SeparateValuesByBit(report, 0)
	assert.Equal(t, 7, len(ones))
	assert.Equal(t, 5, len(zeros))
	ones, zeros = SeparateValuesByBit(report, 2)
	assert.Equal(t, 8, len(ones))
	assert.Equal(t, 4, len(zeros))
}

func TestFindOxygenValues(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	oxygen := FindOxygenValues(report)
	assert.Equal(t, "10111", oxygen)
}

func TestFindCO2Values(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	c02 := FindCO2Values(report)
	assert.Equal(t, "01010", c02)
}

func TestFindGammaEpislonValues(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	gamma, epsilon := FindGammaEpsilonValues(report)
	assert.Equal(t, "10110", gamma)
	assert.Equal(t, "01001", epsilon)
}

func TestConvertValuesToNumbers(t *testing.T) {
	gamma, epsilon := ConvertValuesToNumbers("10110", "01001")
	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)
}

func TestFindValue(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	assert.Equal(t, 198, FindValue(report))
}

func TestFindOxyCO2Values(t *testing.T) {
	report := util.ReadFile("./03_test_input.txt")
	assert.Equal(t, 230, FindOxyCO2Values(report))
}
