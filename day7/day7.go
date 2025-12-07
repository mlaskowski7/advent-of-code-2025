package day7

import (
	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

const splitter = '^'
const start = 'S'

func GetTotalCountOfSplits() (int, error) {
	matrix, err := utils.ReadInputAsMatrix("day7/input.txt")
	if err != nil {
		return 0, err
	}

	R := len(matrix)
	C := len(matrix[0])

	sr, sc := 0, 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matrix[r][c] == start {
				sr, sc = r, c
			}
		}
	}

	p1 := 0
	Q := []struct{ r, c int }{{sr, sc}}
	seen := make(map[struct{ r, c int }]bool)

	for len(Q) > 0 {
		r, c := Q[0].r, Q[0].c
		Q = Q[1:]

		key := struct{ r, c int }{r, c}
		if seen[key] {
			continue
		}
		seen[key] = true

		if r+1 == R {
			continue
		}

		if matrix[r+1][c] == splitter {
			Q = append(Q, struct{ r, c int }{r + 1, c - 1})
			Q = append(Q, struct{ r, c int }{r + 1, c + 1})
			p1++
		} else {
			Q = append(Q, struct{ r, c int }{r + 1, c})
		}
	}

	return p1, nil
}

func GetTotalTimelines() (int, error) {
	matrix, err := utils.ReadInputAsMatrix("day7/input.txt")
	if err != nil {
		return 0, err
	}

	R := len(matrix)
	C := len(matrix[0])

	sr, sc := 0, 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matrix[r][c] == start {
				sr, sc = r, c
			}
		}
	}
	memo := make(map[struct{ r, c int }]int)
	var scoreFunc func(int, int) int
	scoreFunc = func(r, c int) int {
		if r+1 == R {
			return 1
		}

		key := struct{ r, c int }{r, c}
		if val, ok := memo[key]; ok {
			return val
		}

		var result int
		if matrix[r+1][c] == splitter {
			result = scoreFunc(r+1, c-1) + scoreFunc(r+1, c+1)
		} else {
			result = scoreFunc(r+1, c)
		}

		memo[key] = result
		return result
	}

	p2 := scoreFunc(sr, sc)
	return p2, nil
}
