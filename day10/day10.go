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

func GetFewestButtonPresesPart2() (int, error) {
	lines, err := utils.ReadLines("day10/input.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for i, line := range lines {
		presses, err := solveJoltageLine(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i+1, err)
		}
		total += presses
	}
	return total, nil
}

type Machine struct {
	buttons  []uint32
	joltages []int32
}

type SubsetXOR struct {
	subset []int
	xor    uint32
}

func solveJoltageLine(line string) (int, error) {
	buttonsRe := regexp.MustCompile(`\(([^)]*)\)`)
	joltagesRe := regexp.MustCompile(`\{([^}]*)\}`)

	var buttons []uint32
	for _, m := range buttonsRe.FindAllStringSubmatch(line, -1) {
		if len(m) < 2 {
			continue
		}
		content := strings.TrimSpace(m[1])
		mask := uint32(0)
		if content != "" {
			for _, s := range strings.Split(content, ",") {
				s = strings.TrimSpace(s)
				if s != "" {
					v, err := strconv.Atoi(s)
					if err != nil {
						return 0, fmt.Errorf("invalid button index '%s'", s)
					}
					mask |= 1 << uint32(v)
				}
			}
		}
		buttons = append(buttons, mask)
	}

	joltagesMatch := joltagesRe.FindStringSubmatch(line)
	if len(joltagesMatch) < 2 {
		return 0, fmt.Errorf("invalid joltages pattern")
	}
	var joltages []int32
	for _, s := range strings.Split(joltagesMatch[1], ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			v, err := strconv.Atoi(s)
			if err != nil {
				return 0, fmt.Errorf("invalid joltage '%s'", s)
			}
			joltages = append(joltages, int32(v))
		}
	}

	var subsetXors []SubsetXOR
	numButtons := len(buttons)
	for mask := 0; mask < (1 << numButtons); mask++ {
		var subset []int
		xorVal := uint32(0)
		for i := 0; i < numButtons; i++ {
			if (mask & (1 << i)) != 0 {
				subset = append(subset, i)
				xorVal ^= buttons[i]
			}
		}
		subsetXors = append(subsetXors, SubsetXOR{subset, xorVal})
	}

	machine := &Machine{buttons, joltages}
	result, found := machine.solveRecursive(subsetXors, joltages)
	if found {
		return result, nil
	}
	return 0, fmt.Errorf("no solution found")
}

func (m *Machine) solveRecursive(subsetXors []SubsetXOR, joltages []int32) (int, bool) {
	allZero := true
	for _, j := range joltages {
		if j != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0, true
	}

	parity := uint32(0)
	for i, j := range joltages {
		if j%2 != 0 {
			parity |= 1 << uint32(i)
		}
	}

	var best *int = nil
	for _, su := range subsetXors {
		if su.xor != parity {
			continue
		}

		newJoltages := make([]int32, len(joltages))
		valid := true
		for i := 0; i < len(joltages); i++ {
			count := int32(0)
			for _, btnIdx := range su.subset {
				if (m.buttons[btnIdx] & (1 << uint32(i))) != 0 {
					count++
				}
			}
			diff := joltages[i] - count
			if diff < 0 {
				valid = false
				break
			}
			newJoltages[i] = diff / 2
		}
		if !valid {
			continue
		}

		if recResult, found := m.solveRecursive(subsetXors, newJoltages); found {
			pressesSoFar := len(su.subset) + 2*recResult
			if best == nil || pressesSoFar < *best {
				best = &pressesSoFar
			}
		}
	}

	if best != nil {
		return *best, true
	}
	return 0, false
}
