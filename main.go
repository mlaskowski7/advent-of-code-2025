package main

import (
	"fmt"

	"github.com/mlaskowski7/advent-of-code-2025/day1"
	"github.com/mlaskowski7/advent-of-code-2025/day2"
)

func main() {
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

	invalidIDsSum, err := day2.GetInvalidIDsSum()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 2 part 1 - %v", err)
	} else {
		fmt.Printf("Day 2, Part 1 answer is %d\n", invalidIDsSum)
	}

}
