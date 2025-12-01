package main

import (
	"fmt"

	"github.com/mlaskowski7/advent-of-code-2025/day1"
)

func main() {
	pass, err := day1.GetPassword()
	if err != nil {
		fmt.Printf("An error occured while getting solution for day 1 task 1 - %v", err)
	} else {
		fmt.Printf("Day 1, Part 1 answer is %d\n", pass)
	}

}
