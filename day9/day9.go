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
	points, err := GetPoints(lines)
	if err != nil {
		return 0, err
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

func GetLargestRectangleAreaPart2() (int, error) {
	lines, err := utils.ReadLines("day9/input.txt")
	if err != nil {
		return 0, err
	}

	points, err := GetPoints(lines)
	if err != nil {
		return 0, err
	}

	border := GetBorder(points)
	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			t1 := points[i]
			t2 := points[j]

			x0, x1 := t1.X, t2.X
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			y0, y1 := t1.Y, t2.Y
			if y0 > y1 {
				y0, y1 = y1, y0
			}

			allGreen := true

			for x, ys := range border.VLines {
				yStart, yEnd := ys[0], ys[1]
				if x > x0 && x < x1 && rangeIntersection(yStart, yEnd+1, y0+1, y1) {
					allGreen = false
					break
				}
			}

			if allGreen {
				for y, xs := range border.HLines {
					xStart, xEnd := xs[0], xs[1]
					if y > y0 && y < y1 && rangeIntersection(xStart, xEnd+1, x0+1, x1) {
						allGreen = false
						break
					}
				}
			}

			if allGreen {
				area := (x1 - x0 + 1) * (y1 - y0 + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}

type Border struct {
	VLines map[int][2]int
	HLines map[int][2]int
}

func GetBorder(points []Coords) *Border {
	vlines := make(map[int][2]int)
	hlines := make(map[int][2]int)

	n := len(points)
	for i := 0; i < n; i++ {
		t1 := points[i]
		t2 := points[(i+1)%n]

		if t1.X == t2.X {
			y0, y1 := t1.Y, t2.Y
			if y0 > y1 {
				y0, y1 = y1, y0
			}
			vlines[t1.X] = [2]int{y0, y1}
		} else if t1.Y == t2.Y {
			x0, x1 := t1.X, t2.X
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			hlines[t1.Y] = [2]int{x0, x1}
		}
	}

	return &Border{VLines: vlines, HLines: hlines}
}

func rangeIntersection(start1, end1, start2, end2 int) bool {
	return start1 < end2 && start2 < end1
}

func GetPoints(lines []string) ([]Coords, error) {
	points := make([]Coords, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		points = append(points, Coords{X: x, Y: y})
	}

	return points, nil
}

func GetArea(left, right Coords) int {
	width := int(math.Abs(float64(left.X-right.X))) + 1
	height := int(math.Abs(float64(left.Y-right.Y))) + 1
	return width * height
}
