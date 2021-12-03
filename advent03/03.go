package advent03

import "strconv"

func FindBits(report []string, bit_index int) (int, int) {
	ones := 0
	for _, line := range report {
		if len(line) > bit_index && string(line[bit_index]) == "1" {
			ones++
		}
	}
	return ones, len(report) - ones
}

func SeparateValuesByBit(report []string, bit_index int) ([]string, []string) {
	ones := []string{}
	zeros := []string{}
	for _, line := range report {
		if len(line) > bit_index {
			if string(line[bit_index]) == "1" {
				ones = append(ones, line)
			} else {
				zeros = append(zeros, line)
			}
		}
	}
	return ones, zeros
}

func FindGammaEpsilonValues(report []string) (string, string) {
	gamma := ""
	epsilon := ""
	for column := 0; column < len(report[0]); column++ {
		ones, zeros := FindBits(report, column)
		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			epsilon += "1"
			gamma += "0"
		}
	}
	return gamma, epsilon
}

func FindOxygenValues(report []string) string {
	oxygen := report
	for column := 0; column < len(report[0]); column++ {
		if len(oxygen) == 1 {
			break
		}
		ones, zeros := SeparateValuesByBit(oxygen, column)
		if len(ones) >= len(zeros) {
			oxygen = ones
		} else {
			oxygen = zeros
		}
	}
	return oxygen[0]
}

func FindCO2Values(report []string) string {
	co2 := report
	for column := 0; column < len(report[0]); column++ {
		if len(co2) == 1 {
			break
		}
		ones, zeros := SeparateValuesByBit(co2, column)
		if len(ones) < len(zeros) {
			co2 = ones
		} else {
			co2 = zeros
		}
	}
	return co2[0]
}

func ConvertValuesToNumbers(gamma string, epsilon string) (int, int) {
	gamma_int, _ := strconv.ParseInt(gamma, 2, 32)
	epsilon_int, _ := strconv.ParseInt(epsilon, 2, 32)
	return int(gamma_int), int(epsilon_int)
}

func FindValue(report []string) int {
	gamma, epsilon := FindGammaEpsilonValues(report)
	gamma_int, epsilon_int := ConvertValuesToNumbers(gamma, epsilon)
	return gamma_int * epsilon_int
}

func FindOxyCO2Values(report []string) int {
	oxygen := FindOxygenValues(report)
	co2 := FindCO2Values(report)
	oxy_int, co2_int := ConvertValuesToNumbers(oxygen, co2)
	return oxy_int * co2_int
}
