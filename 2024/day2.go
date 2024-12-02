package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	safe := 0

	for _, v := range lines {
		r := strings.Split(v, " ")
		var digits []int

		for i := 0; i < len(r); i++ {
			digit, _ := strconv.Atoi(r[i])
			digits = append(digits, digit)
		}
		if reportIsSafe(digits) {
			safe++
		}
	}
	fmt.Printf("Day 2.1 solution: %d\n", safe)
}

// TODO: Remove the first instance of an error and double check
func Day2_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	safe := 0

	for _, v := range lines {
		r := strings.Split(v, " ")
		var digits []int

		for i := 0; i < len(r); i++ {
			digit, _ := strconv.Atoi(r[i])
			digits = append(digits, digit)
		}
		isSafe := false
		isSafe = reportIsSafe(digits)

		if !isSafe {
			for i := 0; i < len(digits); i++ {
				modifiedDigits := []int{}
				modifiedDigits = append(modifiedDigits, digits[:i]...)
				modifiedDigits = append(modifiedDigits, digits[i+1:]...)
				isSafe = reportIsSafe(modifiedDigits)

				if isSafe {
					break
				}
			}
		}
		if isSafe {
			safe++
		}
	}
	fmt.Printf("Day 2.2 solution: %d\n", safe)
}

func reportIsSafe(report []int) bool {
	isSafe := true
	prevDiff := 0
	for i := 0; i <= len(report)-2; i++ {
		diff := report[i] - report[i+1]
		if diff == 0 || diff > 3 || diff < -3 {
			isSafe = false
			break
		}
		if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
			isSafe = false
			break
		}
		prevDiff = diff
		if i == len(report)-2 {
			isSafe = true
		}
	}
	return isSafe
}
