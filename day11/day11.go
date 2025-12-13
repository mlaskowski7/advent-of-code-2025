package day11

import (
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type AdjecencyList map[string][]string

func GetPathsCount() (int, error) {
	lines, err := utils.ReadLines("day11/input.txt")
	if err != nil {
		return 0, err
	}

	graph := buildAdjecencyList(lines)
	pathsCount := dfs(graph, "you")
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
