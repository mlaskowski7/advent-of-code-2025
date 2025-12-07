package main

import (
	"fmt"

	"github.com/mlaskowski7/advent-of-code-2025/day1"
	"github.com/mlaskowski7/advent-of-code-2025/day2"
	"github.com/mlaskowski7/advent-of-code-2025/day3"
	"github.com/mlaskowski7/advent-of-code-2025/day4"
	"github.com/mlaskowski7/advent-of-code-2025/day5"
	"github.com/mlaskowski7/advent-of-code-2025/day6"
	"github.com/mlaskowski7/advent-of-code-2025/day7"
)

func main() {
	// day 1
	pass, err := day1.GetPasswordPart1()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 1 part 1 - %v", err)
	} else {
		fmt.Printf("Day 1, Part 1 answer is %d\n", pass)
	}
	pass, err = day1.GetPasswordPart2()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 1 part 2 - %v", err)
	} else {
		fmt.Printf("Day 1, Part 2 answer is %d\n", pass)
	}

	// day 2
	invalidIDsSum, err := day2.GetInvalidIDsSumPart1()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 2 part 1 - %v", err)
	} else {
		fmt.Printf("Day 2, Part 1 answer is %d\n", invalidIDsSum)
	}
	invalidIDsSum, err = day2.GetInvalidIDsSumPart2()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 2 part 2 - %v", err)
	} else {
		fmt.Printf("Day 2, Part 2 answer is %d\n", invalidIDsSum)
	}

	// day 3
	totalMaxJoltage, err := day3.GetTotalMaxJoltage()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 3 part 1 - %v", err)
	} else {
		fmt.Printf("Day 3, Part 1 answer is %d\n", totalMaxJoltage)
	}
	totalMaxJoltagePart2, err := day3.GetTotalMaxJoltagePart2()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 3 part 2 - %v", err)
	} else {
		fmt.Printf("Day 3, Part 2 answer is %d\n", totalMaxJoltagePart2)
	}

	// day 4
	accessibleRollsCount, err := day4.GetAccessibleRollsCount()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 4 part 1 - %v", err)
	} else {
		fmt.Printf("Day 4, Part 1 answer is %d\n", accessibleRollsCount)
	}
	accessibleRollsCountPart2, err := day4.GetAccessibleRollsCountPart2()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 4 part 2 - %v", err)
	} else {
		fmt.Printf("Day 4, Part 2 answer is %d\n", accessibleRollsCountPart2)
	}

	// day 5
	availableFreshIDsCount, err := day5.GetAvailableFreshIDs()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 5 part 1 - %v", err)
	} else {
		fmt.Printf("Day 5, Part 1 answer is %d\n", availableFreshIDsCount)
	}
	totalFreshIDs, err := day5.GetTotalFreshIDs()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 5 part 2 - %v", err)
	} else {
		fmt.Printf("Day 5, Part 2 answer is %d\n", totalFreshIDs)
	}

	// day 6
	calculationsSum, err := day6.GetCalculationsSum()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 6 part 1 - %v\n", err)
	} else {
		fmt.Printf("Day 6, Part 1 answer is %d\n", calculationsSum)
	}
	calculationsSumPart2, err := day6.GetCalculationsSumPart2()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 6 part 2 - %v\n", err)
	} else {
		fmt.Printf("Day 6, Part 2 answer is %d\n", calculationsSumPart2)
	}

	// day 7
	splitsCount, err := day7.GetTotalCountOfSplits()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 7 part 1 - %v\n", err)
	} else {
		fmt.Printf("Day 7, Part 1 answer is %d\n", splitsCount)
	}
}
