package day2

import (
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type IsSillyDelegate func(num int) bool

func GetInvalidIDsSumPart1() (int, error) {
	return GetInvalidIDsSum(isSillyPart1)
}

func GetInvalidIDsSumPart2() (int, error) {
	return GetInvalidIDsSum(isSillyPart2)
}

func GetInvalidIDsSum(isSilly IsSillyDelegate) (int, error) {
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

func isSillyPart1(num int) bool {
	numAsStr := strconv.Itoa(num)
	half := len(numAsStr) / 2
	return numAsStr[:half] == numAsStr[half:]
}

func isSillyPart2(num int) bool {
	numAsStr := strconv.Itoa(num)

	strLen := len(numAsStr)
	for i := 1; i <= strLen/2; i++ {
		if strLen%i != 0 {
			continue
		}

		left := numAsStr[:i]
		isMatch := true
		for j := i; j < strLen; j += i {
			if numAsStr[j:j+i] != left {
				isMatch = false
				break
			}
		}

		if isMatch {
			return true
		}
	}

	return false
}
