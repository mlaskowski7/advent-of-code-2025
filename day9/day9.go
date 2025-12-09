package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type Coords struct {
	X int
	Y int
}

func GetLargestRectangleArea() (int, error) {
	lines, err := utils.ReadLines("day9/input.txt")
	if err != nil {
		return 0, err
	}

	n := len(lines)

	maxArea := 0
	points := make([]Coords, 0, n)
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		points = append(points, Coords{X: x, Y: y})
	}

	for i, point := range points {
		for j := i + 1; j < n; j++ {
			area := GetArea(point, points[j])
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func GetArea(left, right Coords) int {
	width := int(math.Abs(float64(left.X-right.X))) + 1
	height := int(math.Abs(float64(left.Y-right.Y))) + 1
	return width * height
}
