package day1

import (
	"strconv"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetPassword() (int, error) {
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

		curr %= 100
		if curr == 0 {
			password++
		}
	}

	return password, nil
}
