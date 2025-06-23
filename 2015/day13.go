package main

import (
	"fmt"
	"github.com/etnz/permute"
	"slices"
	"strconv"
	"strings"
)

func Day13_1(input string, part int) {
	happiest := 0
	rules := make(map[[2]string]int)
	people := []string{}
	if part == 2 {
		people = append(people, "Me")
	}
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, v := range lines {
		var happiness int
		rule := strings.Split(strings.TrimSpace(v), " ")
		if !slices.Contains(people, rule[0]) {
			people = append(people, strings.Trim(rule[0], "."))
		}
		points, _ := strconv.Atoi(rule[3])

		if rule[2] == "gain" {
			happiness = points
		} else {
			happiness = points * -1
		}
		rules[[2]string{rule[0], strings.Trim(rule[(len(rule)-1)], ".")}] = happiness
	}

	perms := permute.Permutations(people)

	for _, perm := range perms {
		sum := 0
		for i := 0; i < len(perm); i++ {
			var seats1 [2]string
			var seats2 [2]string
			if i == len(perm)-1 {
				seats1 = [2]string{perm[i], perm[0]}
				seats2 = [2]string{perm[0], perm[i]}
			} else {
				seats1 = [2]string{perm[i], perm[i+1]}
				seats2 = [2]string{perm[i+1], perm[i]}
			}
			happy1, err1 := rules[seats1]
			happy2, err2 := rules[seats2]

			//for part 2
			//if the rules isn't in the dictionary they are sitting by me and it's 0 change for both directions
			if err1 == false && err2 == false {
				sum += 0
			} else {
				sum += happy1 + happy2
			}
		}
		if sum > happiest {
			happiest = sum
		}
	}
	fmt.Printf("Day 13 Part 1 - Happiest value: %d\n", happiest)
}
