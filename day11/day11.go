package day11

import (
	"strings"
	"sync"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type AdjecencyList map[string][]string
type State struct {
	Node        string
	VisitedMask int
}

var stateCache map[State]int
var once sync.Once

var requiredNodes = map[string]int{
	"dac": 1,
	"fft": 2,
}

func GetPathsCount() (int, error) {
	lines, err := utils.ReadLines("day11/input.txt")
	if err != nil {
		return 0, err
	}

	graph := buildAdjecencyList(lines)
	pathsCount := dfs(graph, "you")
	return pathsCount, nil
}

func GetPathsCountPart2() (int, error) {
	lines, err := utils.ReadLines("day11/input.txt")
	if err != nil {
		return 0, err
	}

	once.Do(func() {
		stateCache = make(map[State]int)
	})

	graph := buildAdjecencyList(lines)

	pathsCount := dfsPart2(graph, "svr", 0, 3)
	return pathsCount, nil
}
func dfs(graph AdjecencyList, curr string) int {
	if curr == "out" {
		return 1
	}

	count := 0
	children := graph[curr]
	for _, child := range children {
		count += dfs(graph, child)
	}

	return count
}

func dfsPart2(graph AdjecencyList, curr string, visitedMask int, requiredMask int) int {
	state := State{Node: curr, VisitedMask: visitedMask}
	if result, ok := stateCache[state]; ok {
		return result
	}

	if curr == "out" {
		if visitedMask == requiredMask {
			return 1
		}
		return 0
	}

	nextVisitedMask := visitedMask
	if bit, ok := requiredNodes[curr]; ok {
		nextVisitedMask |= bit
	}

	count := 0
	children := graph[curr]
	for _, child := range children {
		count += dfsPart2(graph, child, nextVisitedMask, requiredMask)
	}

	stateCache[state] = count
	return count
}

func buildAdjecencyList(lines []string) AdjecencyList {
	graph := make(AdjecencyList)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		root := split[0]
		children := strings.Split(split[1], " ")
		graph[root] = append(graph[root], children...)
	}

	return graph
}
