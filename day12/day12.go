package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

const (
	EMPTY  = '.'
	FILLED = '#'
)

type Coordinate struct {
	x, y int
}

type Present struct {
	id    int
	shape []string
}

type Piece struct {
	area       int
	id         int
	shape      []string
	placements [][]Coordinate
}

func GetRegionsCount() (int, error) {
	lines, err := utils.ReadLines("day12/input.txt")
	if err != nil {
		return 0, err
	}

	presentRegex := regexp.MustCompile(`^\d+:`)

	input := strings.Join(lines, "\n")
	sections := strings.Split(input, "\n\n")

	presents := []Present{}
	regions := []string{}

	for _, section := range sections {
		if matched := presentRegex.MatchString(section); matched {
			lines := strings.Split(section, "\n")
			idStr := strings.TrimSuffix(lines[0], ":")
			id, _ := strconv.Atoi(idStr)
			shape := lines[1:]
			presents = append(presents, Present{id: id, shape: shape})
		} else {
			regionLines := strings.Split(section, "\n")
			for _, line := range regionLines {
				if strings.TrimSpace(line) != "" {
					regions = append(regions, line)
				}
			}
		}
	}

	answer := 0
	for i, region := range regions {
		fmt.Printf("Region %d of %d\n", i+1, len(regions))

		parts := strings.Split(region, ":")
		dimensions := strings.TrimSpace(parts[0])
		countsStr := strings.TrimSpace(parts[1])

		dimParts := strings.Split(dimensions, "x")
		width, _ := strconv.Atoi(dimParts[0])
		height, _ := strconv.Atoi(dimParts[1])

		countStrs := strings.Fields(countsStr)
		thePresents := []Present{}
		for idx, countStr := range countStrs {
			count, _ := strconv.Atoi(countStr)
			for n := 0; n < count; n++ {
				uniqueID := (n+1)*10 + idx
				thePresents = append(thePresents, Present{
					id:    uniqueID,
					shape: presents[idx].shape,
				})
			}
		}

		underTheTree := make([][]rune, height)
		for y := 0; y < height; y++ {
			row := make([]rune, width)
			for x := 0; x < width; x++ {
				row[x] = EMPTY
			}
			underTheTree[y] = row
		}

		if tryToFit(thePresents, underTheTree) {
			answer++
		}
	}

	return answer, nil
}

func tryToFit(thePresents []Present, underTheTree [][]rune) bool {
	if len(thePresents) == 0 {
		return true
	}

	pieces := []Piece{}
	for _, present := range thePresents {
		piece := shapeToPiece(present.shape, present.id, underTheTree)
		pieces = append(pieces, piece)
	}

	sortPiecesByPlacements(pieces)

	presentArea := 0
	for _, piece := range pieces {
		presentArea += piece.area
	}
	treeArea := len(underTheTree) * len(underTheTree[0])

	if presentArea > treeArea {
		fmt.Println("Short cut!")
		return false
	}

	for _, placement := range pieces[0].placements {
		if canPlace(placement, underTheTree) {
			place(placement, underTheTree, pieces[0].id)

			if placeNext(pieces, 1, underTheTree) {
				unplace(placement, underTheTree)
				return true
			}

			unplace(placement, underTheTree)
		}
	}

	return false
}

func placeNext(pieces []Piece, pieceIndex int, board [][]rune) bool {
	if pieceIndex >= len(pieces) {
		return true
	}

	for _, placement := range pieces[pieceIndex].placements {
		if canPlace(placement, board) {
			place(placement, board, pieces[pieceIndex].id)
			if placeNext(pieces, pieceIndex+1, board) {
				unplace(placement, board)
				return true
			}
			unplace(placement, board)
		}
	}

	return false
}

func canPlace(placement []Coordinate, board [][]rune) bool {
	for _, coord := range placement {
		if board[coord.y][coord.x] != EMPTY {
			return false
		}
	}
	return true
}

func place(placement []Coordinate, board [][]rune, id int) {
	for _, coord := range placement {
		board[coord.y][coord.x] = rune('0' + (id % 10))
	}
}

func unplace(placement []Coordinate, board [][]rune) {
	for _, coord := range placement {
		board[coord.y][coord.x] = EMPTY
	}
}

func rotate(shape []string) []string {
	if len(shape) == 0 {
		return shape
	}
	rotated := make([]string, len(shape[0]))
	for x := 0; x < len(shape[0]); x++ {
		row := ""
		for y := len(shape) - 1; y >= 0; y-- {
			row += string(shape[y][x])
		}
		rotated[x] = row
	}
	return rotated
}

func shapeToCoordinates(shape []string) []Coordinate {
	coordinates := []Coordinate{}
	for y, row := range shape {
		for x, space := range row {
			if space == FILLED {
				coordinates = append(coordinates, Coordinate{x: x, y: y})
			}
		}
	}
	return coordinates
}

func shapeToPiece(shape []string, id int, board [][]rune) Piece {
	area := 0
	for _, row := range shape {
		for _, ch := range row {
			if ch == FILLED {
				area++
			}
		}
	}

	placements := [][]Coordinate{}
	variations := make(map[string]bool)

	addPlacements := func(s []string) {
		variation := strings.Join(s, "\n")
		if !variations[variation] {
			variations[variation] = true
			placements = append(placements, shapeToPlacements(s, board)...)
		}
	}

	current := make([]string, len(shape))
	copy(current, shape)

	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			addPlacements(current)
			current = rotate(current)
		}
		reversed := make([]string, len(current))
		for i := 0; i < len(current); i++ {
			reversed[i] = current[len(current)-1-i]
		}
		current = reversed
	}

	return Piece{
		area:       area,
		id:         id,
		shape:      shape,
		placements: placements,
	}
}

func shapeToPlacements(shape []string, board [][]rune) [][]Coordinate {
	placements := [][]Coordinate{}
	coordinates := shapeToCoordinates(shape)

	if len(coordinates) == 0 {
		return placements
	}

	maxX, maxY := 0, 0
	for _, coord := range coordinates {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}
	pieceWidth := maxX + 1
	pieceHeight := maxY + 1

	boardWidth := len(board[0])
	boardHeight := len(board)

	for y := 0; y <= boardHeight-pieceHeight; y++ {
		for x := 0; x <= boardWidth-pieceWidth; x++ {
			placement := []Coordinate{}
			valid := true
			for _, coord := range coordinates {
				newX := x + coord.x
				newY := y + coord.y
				if newX >= len(board[newY]) {
					valid = false
					break
				}
				placement = append(placement, Coordinate{x: newX, y: newY})
			}
			if valid {
				placements = append(placements, placement)
			}
		}
	}

	return placements
}

func sortPiecesByPlacements(pieces []Piece) {
	for i := 0; i < len(pieces)-1; i++ {
		for j := 0; j < len(pieces)-i-1; j++ {
			if len(pieces[j].placements) > len(pieces[j+1].placements) {
				pieces[j], pieces[j+1] = pieces[j+1], pieces[j]
			}
		}
	}
}
