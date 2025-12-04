package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadInputAsMatrix(path string) ([][]rune, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, r := range line {
			row = append(row, r)
		}

		matrix = append(matrix, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return matrix, nil
}

func ReadSingleLineInputSeparatedByCommas(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return strings.Split(line, ","), nil
}
