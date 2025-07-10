package main

import (
	"fmt"
	// "regexp"
	"strings"
)

func Day19(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// molecule := lines[len(lines)-1]
	rules := lines[:len(lines)-2]

	for _, rule := range rules {
		ruleParts := strings.Split(rule, " => ")
		find := strings.TrimSpace(ruleParts[0])
		replace := strings.TrimSpace(ruleParts[1])

		fmt.Println(find, replace)
		// r := regexp.MustCompile(find)
	}
}
