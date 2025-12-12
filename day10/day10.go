package day10

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

func GetFewestButtonPreses() (int, error) {
	lines, err := utils.ReadLines("day10/input.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for i, line := range lines {
		presses, err := getFewestButtonPresesFromLine(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i+1, err)
		}
		total += presses
	}
	return total, nil
}

func getFewestButtonPresesFromLine(line string) (int, error) {
	indicatorLightsRegex := regexp.MustCompile(`\[([.#]+)\]`)
	buttonsRegex := regexp.MustCompile(`\(([^)]*)\)`)

	indicatorLightsMatch := indicatorLightsRegex.FindStringSubmatch(line)
	if len(indicatorLightsMatch) < 2 {
		return 0, fmt.Errorf("invalid bracket pattern")
	}

	lightsPattern := indicatorLightsMatch[1]
	n := len(lightsPattern)
	target := 0
	for idx, ch := range lightsPattern {
		if ch == '#' {
			target |= 1 << idx
		}
	}

	var buttons []int
	for _, m := range buttonsRegex.FindAllStringSubmatch(line, -1) {
		if len(m) < 2 {
			continue
		}
		content := strings.TrimSpace(m[1])
		if content == "" {
			buttons = append(buttons, 0)
			continue
		}

		parts := strings.Split(content, ",")
		mask := 0
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}

			v, err := strconv.Atoi(p)
			if err != nil {
				return 0, fmt.Errorf("invalid button index '%s'", p)
			}
			if v < 0 || v >= n {
				return 0, fmt.Errorf("button index out of range: %d", v)
			}
			mask |= 1 << v
		}
		buttons = append(buttons, mask)
	}

	maxState := 1 << n
	dist := make([]int, maxState)
	for i := range dist {
		dist[i] = -1
	}
	q := list.New()
	start := 0
	dist[start] = 0
	q.PushBack(start)
	for q.Len() > 0 {
		curElem := q.Front()
		q.Remove(curElem)
		cur := curElem.Value.(int)
		if cur == target {
			return dist[cur], nil
		}
		for _, b := range buttons {
			nxt := cur ^ b
			if dist[nxt] == -1 {
				dist[nxt] = dist[cur] + 1
				q.PushBack(nxt)
			}
		}
	}

	return 0, fmt.Errorf("didnt found a result")
}
