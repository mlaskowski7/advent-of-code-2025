package day1

import (
	"strconv"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetPasswordPart1() (int, error) {
	instructions, err := utils.ReadLines("day1/input.txt")
	if err != nil {
		return 0, err
	}

	password := 0
	curr := 50
	for _, instruction := range instructions {
		diff, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return 0, err
		}

		if instruction[0] == 'R' {
			curr += diff
		} else {
			curr -= diff
		}

		curr = ((curr % 100) + 100) % 100
		if curr == 0 {
			password++
		}
	}

	return password, nil
}

func GetPasswordPart2() (int, error) {
	instructions, err := utils.ReadLines("day1/input.txt")
	if err != nil {
		return 0, err
	}

	password := 0
	curr := 50
	for _, instruction := range instructions {
		diff, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return 0, err
		}

		if diff > 100 {
			password += diff / 100
			diff %= 100
		}

		if instruction[0] == 'R' {
			curr += diff
			if curr > 99 {
				password++
				curr %= 100
			}
		} else {
			wasZero := curr == 0
			curr -= diff
			if !wasZero && curr <= 0 {
				password++
			}

			if curr < 0 {
				curr += 100
			}
		}
	}

	return password, nil
}
