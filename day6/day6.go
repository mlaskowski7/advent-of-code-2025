package day6

import (
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetCalculationsSum() (int, error) {
	lines, err := utils.ReadLines("day6/input.txt")
	if err != nil {
		return 0, err
	}

	if len(lines) == 0 {
		return 0, nil
	}

	operations := lines[len(lines)-1]

	grid := make([][]int, 0)
	for i := 0; i < len(lines)-1; i++ {
		gridRow := []int{}
		split := strings.Split(lines[i], " ")
		for _, elem := range split {
			if elem != "" {
				parsed, err := strconv.Atoi(elem)
				if err != nil {
					return 0, err
				}
				gridRow = append(gridRow, parsed)
			}
		}
		if len(gridRow) > 0 {
			grid = append(grid, gridRow)
		}
	}

	if len(grid) == 0 {
		return 0, nil
	}

	numCols := len(grid[0])

	grandTotal := 0
	opIdx := 0

	for col := 0; col < numCols; col++ {
		isEmpty := true
		numbers := []int{}

		for row := 0; row < len(grid); row++ {
			if col < len(grid[row]) && grid[row][col] != 0 {
				isEmpty = false
				numbers = append(numbers, grid[row][col])
			}
		}

		if isEmpty || len(numbers) == 0 {
			continue
		}

		for opIdx < len(operations) && operations[opIdx] == ' ' {
			opIdx++
		}

		if opIdx >= len(operations) {
			break
		}

		op := operations[opIdx]
		opIdx++

		result := numbers[0]
		for i := 1; i < len(numbers); i++ {
			switch op {
			case '+':
				result += numbers[i]
			case '*':
				result *= numbers[i]
			}
		}

		grandTotal += result
	}

	return grandTotal, nil
}
