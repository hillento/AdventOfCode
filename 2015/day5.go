package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Day5_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	nice := 0

	for _, l := range lines {
		m1, _ := regexp.MatchString(`[aeiou].*[aeiou].*[aeiou]`, l) //3 vowels
		m2, _ := regexp.MatchString(`(ab|cd|pq|xy)`, l)             //does not contain substrings
		m3 := false                                                 //double letter

		for i := range l {
			if i == len(l)-1 {
				break
			}
			if l[i] == l[i+1] {
				m3 = true
				break
			}
		}

		if m1 && !m2 && m3 {
			nice++
		}
	}
	fmt.Printf("Day 5.1 - Nice Count: %d\n", nice)
}

func Day5_2(input string) {

	//a function with a return is just easier to break from than nested loops
	checkRule1 := func(l string) bool {
		for i := 0; i < len(l)-2; i++ {
			candidate := l[i : i+2]
			for j := i + 2; j < len(l)-1; j++ {
				if l[j:j+2] == candidate {
					return true
				}
			}
		}
		return false
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	nice := 0

	for _, l := range lines {
		rule1 := checkRule1(l)
		rule2 := false
		for i := 0; i < len(l)-2; i++ {
			if !rule1 {
				break
			}
			if l[i] == l[i+2] {
				rule2 = true
				break
			}
		}
		if rule1 && rule2 {
			nice += 1
		}
	}
	fmt.Printf("Day 5.2 - Nice Count: %d\n", nice)
}
