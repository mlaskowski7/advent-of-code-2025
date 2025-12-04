package day4

import "github.com/mlaskowski7/advent-of-code-2025/utils"

func GetAccessibleRollsCount() (int, error) {
	matrix, err := utils.ReadInputAsMatrix("day4/input.txt")
	if err != nil {
		return 0, err
	}

	rows := len(matrix)
	cols := len(matrix[0])

	counter := 0
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '@' && IsValid(row, col, rows, cols, matrix) {
				counter++
			}
		}
	}

	return counter, nil
}

func IsValid(row, col, rowsLen, colsLen int, matrix [][]rune) bool {
	found := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newRow := row + i
			newCol := col + j
			if newRow < rowsLen && newRow >= 0 && newCol < colsLen && newCol >= 0 && matrix[newRow][newCol] == '@' {
				found++
				if found > 3 {
					return false
				}
			}
		}
	}

	return true
}
