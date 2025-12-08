package day8

import (
	"sort"
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type Distance struct {
	b1, b2 int
	dist   int64
}

func parseBoxes() ([][]int, error) {
	lines, err := utils.ReadLines("day8/input.txt")
	if err != nil {
		return nil, err
	}

	boxes := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		boxes[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			boxes[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	return boxes, nil
}

func calculateDistances(boxes [][]int) []Distance {
	distances := []Distance{}
	for b1 := 0; b1 < len(boxes); b1++ {
		for b2 := b1 + 1; b2 < len(boxes); b2++ {
			dx := boxes[b1][0] - boxes[b2][0]
			dy := boxes[b1][1] - boxes[b2][1]
			dz := boxes[b1][2] - boxes[b2][2]
			dist := int64(dx*dx + dy*dy + dz*dz)
			distances = append(distances, Distance{b1, b2, dist})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})
	return distances
}

func GetThreeLargestCircuitsProduct() (int64, error) {
	boxes, err := parseBoxes()
	if err != nil {
		return 0, err
	}

	distances := calculateDistances(boxes)
	circuits := []map[int]bool{}
	mergedIndex := make(map[int]bool)
	p1Limit := 1000

	for i, d := range distances {
		connection := map[int]bool{d.b1: true, d.b2: true}

		merged := false
		for c := 0; c < len(circuits); c++ {
			overlaps := false
			for box := range connection {
				if circuits[c][box] {
					overlaps = true
					break
				}
			}

			if overlaps {
				for box := range connection {
					circuits[c][box] = true
				}

				if i <= p1Limit-1 {
					allInMerged := true
					for box := range connection {
						if !mergedIndex[box] {
							allInMerged = false
							break
						}
					}

					if allInMerged {
						for cc := 0; cc < len(circuits); cc++ {
							if c != cc {
								overlapsCC := false
								for box := range connection {
									if circuits[cc][box] {
										overlapsCC = true
										break
									}
								}

								if overlapsCC {
									for box := range circuits[cc] {
										circuits[c][box] = true
									}
									circuits = append(circuits[:cc], circuits[cc+1:]...)
									break
								}
							}
						}
					}
				}

				merged = true
				break
			}
		}

		if !merged {
			circuits = append(circuits, connection)
		}

		for box := range connection {
			mergedIndex[box] = true
		}

		if i == p1Limit-1 {
			sort.Slice(circuits, func(i, j int) bool {
				return len(circuits[i]) > len(circuits[j])
			})

			if len(circuits) >= 3 {
				return int64(len(circuits[0])) * int64(len(circuits[1])) * int64(len(circuits[2])), nil
			}
		}
	}

	return 0, nil
}

func GetLastConnectionXProduct() (int64, error) {
	boxes, err := parseBoxes()
	if err != nil {
		return 0, err
	}

	distances := calculateDistances(boxes)
	circuits := []map[int]bool{}
	mergedIndex := make(map[int]bool)
	boxIndex := make(map[int]bool)
	for i := 0; i < len(boxes); i++ {
		boxIndex[i] = true
	}

	for i, d := range distances {
		connection := map[int]bool{d.b1: true, d.b2: true}

		merged := false
		for c := 0; c < len(circuits); c++ {
			overlaps := false
			for box := range connection {
				if circuits[c][box] {
					overlaps = true
					break
				}
			}

			if overlaps {
				for box := range connection {
					circuits[c][box] = true
				}

				allInMerged := true
				for box := range connection {
					if !mergedIndex[box] {
						allInMerged = false
						break
					}
				}

				if allInMerged {
					for cc := 0; cc < len(circuits); cc++ {
						if c != cc {
							overlapsCC := false
							for box := range connection {
								if circuits[cc][box] {
									overlapsCC = true
									break
								}
							}

							if overlapsCC {
								for box := range circuits[cc] {
									circuits[c][box] = true
								}
								circuits = append(circuits[:cc], circuits[cc+1:]...)
								break
							}
						}
					}
				}

				merged = true
				break
			}
		}

		if !merged {
			circuits = append(circuits, connection)
		}

		for box := range connection {
			mergedIndex[box] = true
		}

		for box := range connection {
			delete(boxIndex, box)
		}

		if len(boxIndex) == 0 {
			return int64(boxes[d.b1][0]) * int64(boxes[d.b2][0]), nil
		}

		_ = i
	}

	return 0, nil
}
