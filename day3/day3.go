package day3

import (
	"strconv"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetTotalMaxJoltage() (int, error) {
	banks, err := utils.ReadLines("day3/input.txt")
	if err != nil {
		return 0, err
	}

	total := 0
	for _, bank := range banks {
		maxIndex := getMaxIndex(bank, 0, len(bank)-1)
		nextMaxIndex := getMaxIndex(bank, maxIndex+1, len(bank))

		asStr := string(bank[maxIndex]) + string(bank[nextMaxIndex])
		asInt, err := strconv.Atoi(asStr)
		if err != nil {
			return 0, err
		}
		total += asInt
	}

	return total, nil
}

func GetTotalMaxJoltagePart2() (int, error) {
	banks, err := utils.ReadLines("day3/input.txt")
	if err != nil {
		return 0, err
	}

	total := 0
	for _, bank := range banks {
		joltageAsStr := ""
		start := 0
		for i := 0; i < 12; i++ {
			maxIndex := getMaxIndex(bank, start, len(bank)-(12-i)+1)
			start = maxIndex + 1
			joltageAsStr += string(bank[maxIndex])
		}

		joltageAsInt, err := strconv.Atoi(joltageAsStr)
		if err != nil {
			return 0, err
		}
		total += joltageAsInt
	}

	return total, nil
}

func getMaxIndex(bank string, start, end int) int {
	sub := bank[start:end]
	maxIdx := 0
	max := sub[0]

	for i := 1; i < len(sub); i++ {
		if sub[i] > max {
			max = sub[i]
			maxIdx = i
		}
	}

	return start + maxIdx
}
