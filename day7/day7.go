package day7

import "github.com/mlaskowski7/advent-of-code-2025/utils"

const splitter = '^'
const start = 'S'

func GetTotalCountOfSplits() (int, error) {
	matrix, err := utils.ReadInputAsMatrix("day7/input.txt")
	if err != nil {
		return 0, err
	}

	startIndex := 0
	for col := range matrix[0] {
		if matrix[0][col] == start {
			startIndex = col
			break
		}
	}

	visited := make(map[string]bool)
	return countSplits(0, startIndex, matrix, visited), nil
}

func countSplits(row, col int, matrix [][]rune, visited map[string]bool) int {
	rows := len(matrix)
	cols := len(matrix[0])

	if col >= cols || col < 0 {
		return 0
	}
	if row >= rows || row < 0 {
		return 0
	}

	for row < rows && matrix[row][col] != splitter {
		row++
	}

	if row >= rows {
		return 0
	}

	key := string(rune(row)) + "," + string(rune(col))
	if visited[key] {
		return 0
	}
	visited[key] = true

	return 1 + countSplits(row+1, col-1, matrix, visited) + countSplits(row+1, col+1, matrix, visited)
}
