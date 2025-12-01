package main

import (
	"fmt"

	"github.com/mlaskowski7/advent-of-code-2025/day1"
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

}
