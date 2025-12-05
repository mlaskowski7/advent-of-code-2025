package day5

import (
	"sort"
	"strconv"

	"github.com/mlaskowski7/advent-of-code-2025/utils"
)

type interval struct {
	start int
	end   int
}

func GetAvailableFreshIDs() (int, error) {
	lines, err := utils.ReadLines("day5/input.txt")
	if err != nil {
		return 0, err
	}

	sep := -1
	for i, l := range lines {
		if l == "" {
			sep = i
			break
		}
	}
	if sep == -1 {
		return 0, nil
	}

	rangeLines := lines[:sep]
	availLines := lines[sep+1:]

	ranges := make([]interval, 0, len(rangeLines))
	for _, ln := range rangeLines {
		if ln == "" {
			continue
		}
		// split on '-'
		dash := -1
		for i, ch := range ln {
			if ch == '-' {
				dash = i
				break
			}
		}
		if dash == -1 {
			continue
		}
		s, err := strconv.Atoi(ln[:dash])
		if err != nil {
			return 0, err
		}
		e, err := strconv.Atoi(ln[dash+1:])
		if err != nil {
			return 0, err
		}
		if e < s {
			s, e = e, s
		}
		ranges = append(ranges, interval{start: s, end: e})
	}

	if len(ranges) == 0 {
		return 0, nil
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].start < ranges[j].start })
	merged := make([]interval, 0, len(ranges))
	for _, iv := range ranges {
		if len(merged) == 0 || iv.start > merged[len(merged)-1].end+1 {
			merged = append(merged, iv)
		} else {
			if iv.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = iv.end
			}
		}
	}

	freshCount := 0
	for _, ln := range availLines {
		if ln == "" {
			continue
		}
		id, err := strconv.Atoi(ln)
		if err != nil {
			return 0, err
		}
		i := sort.Search(len(merged), func(i int) bool { return merged[i].start > id })
		idx := i - 1
		if idx >= 0 && id <= merged[idx].end {
			freshCount++
		}
	}

	return freshCount, nil
}

func GetTotalFreshIDs() (int, error) {
	lines, err := utils.ReadLines("day5/input.txt")
	if err != nil {
		return 0, err
	}

	ranges := make([]interval, 0)
	for _, ln := range lines {
		if ln == "" {
			break
		}
		dash := -1
		for i, ch := range ln {
			if ch == '-' {
				dash = i
				break
			}
		}
		if dash == -1 {
			continue
		}
		s, err := strconv.Atoi(ln[:dash])
		if err != nil {
			return 0, err
		}
		e, err := strconv.Atoi(ln[dash+1:])
		if err != nil {
			return 0, err
		}
		if e < s {
			s, e = e, s
		}
		ranges = append(ranges, interval{start: s, end: e})
	}

	if len(ranges) == 0 {
		return 0, nil
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].start < ranges[j].start })
	merged := make([]interval, 0, len(ranges))
	for _, iv := range ranges {
		if len(merged) == 0 || iv.start > merged[len(merged)-1].end+1 {
			merged = append(merged, iv)
		} else {
			if iv.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = iv.end
			}
		}
	}

	total := 0
	for _, iv := range merged {
		total += iv.end - iv.start + 1
	}

	return total, nil
}
