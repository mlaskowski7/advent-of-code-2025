package day2

import (
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetInvalidIDsSum() (int, error) {
	ranges, err := utils.ReadSingleLineInputSeparatedByCommas("day2/input.txt")
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, r := range ranges {
		split := strings.Split(r, "-")
		start, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		for i := start; i <= end; i++ {
			if isSilly(i) {
				sum += i
			}
		}
	}

	return sum, nil
}

func isSilly(num int) bool {
	numAsStr := strconv.Itoa(num)
	half := len(numAsStr) / 2
	return numAsStr[:half] == numAsStr[half:]
}
