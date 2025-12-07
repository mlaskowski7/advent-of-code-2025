package day6

import (
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

	G := make([][]rune, len(lines))
	for i, line := range lines {
		G[i] = []rune(line)
	}

	R := len(G)
	C := len(G[0])

	p1 := 0
	startC := 0

	for cc := 0; cc <= C; cc++ {
		isBlank := true
		if cc < C {
			for r := 0; r < R; r++ {
				if G[r][cc] != ' ' {
					isBlank = false
					break
				}
			}
		}

		if isBlank {
			op := G[R-1][startC]

			p1Score := 0
			if op == '*' {
				p1Score = 1
			}

			for r := 0; r < R-1; r++ {
				p1N := 0
				for c := startC; c < cc; c++ {
					if G[r][c] != ' ' {
						digit := int(G[r][c] - '0')
						p1N = p1N*10 + digit
					}
				}
				if op == '*' {
					p1Score *= p1N
				} else {
					p1Score += p1N
				}
			}

			p1 += p1Score
			startC = cc + 1
		}
	}

	return p1, nil
}

func GetCalculationsSumPart2() (int, error) {
	lines, err := utils.ReadLines("day6/input.txt")
	if err != nil {
		return 0, err
	}

	if len(lines) == 0 {
		return 0, nil
	}

	G := make([][]rune, len(lines))
	for i, line := range lines {
		G[i] = []rune(line)
	}

	R := len(G)
	C := len(G[0])

	p2 := 0
	startC := 0

	for cc := 0; cc <= C; cc++ {
		isBlank := true
		if cc < C {
			for r := 0; r < R; r++ {
				if G[r][cc] != ' ' {
					isBlank = false
					break
				}
			}
		}

		if isBlank {
			op := G[R-1][startC]

			score := 0
			if op == '*' {
				score = 1
			}

			for c := cc - 1; c >= startC; c-- {
				n := 0
				for r := 0; r < R-1; r++ {
					if G[r][c] != ' ' {
						digit := int(G[r][c] - '0')
						n = n*10 + digit
					}
				}
				if op == '+' {
					score += n
				} else {
					score *= n
				}
			}

			p2 += score
			startC = cc + 1
		}
	}

	return p2, nil
}
