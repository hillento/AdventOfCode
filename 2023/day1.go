package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day1_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reg := regexp.MustCompile("[0-9]")
	var digits []string
	sum := 0

	for _, v := range lines {
		digits = reg.FindAllString(v, -1)
		lastIndex := len(digits) - 1
		first, _ := strconv.Atoi(digits[0])
		last, _ := strconv.Atoi(digits[lastIndex])

		sum += (first * 10) + last
	}

	fmt.Printf("Day 1.1 Answer: %v\n", sum)
}

func Day1_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reg := regexp.MustCompile("[0-9]")
	var digits []string
	sum := 0
	//This is a little quirly but sometimes numbers can overrun eachotehr eg eightwo. Both need to be counted.
	//First and last letters are preserved to ensure all numbers get proper conversion
	subs := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7v",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for _, v := range lines {
		str := v
		for old, new := range subs {
			str = fmt.Sprintf(strings.Replace(str, old, new, -1))
		}
		digits = reg.FindAllString(str, -1)
		lastIndex := len(digits) - 1
		first, _ := strconv.Atoi(digits[0])
		last, _ := strconv.Atoi(digits[lastIndex])
		sum += (first * 10) + last
	}

	fmt.Printf("Day 1.2 Answer: %v\n", sum)
}
