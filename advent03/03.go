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

func ConvertGammaEpsilonToNumber(gamma string, epsilon string) (int, int) {
	gamma_int, _ := strconv.ParseInt(gamma, 2, 32)
	epsilon_int, _ := strconv.ParseInt(epsilon, 2, 32)
	return int(gamma_int), int(epsilon_int)
}

func FindValue(report []string) int {
	gamma, epsilon := FindGammaEpsilonValues(report)
	gamma_int, epsilon_int := ConvertGammaEpsilonToNumber(gamma, epsilon)
	return gamma_int * epsilon_int
}
