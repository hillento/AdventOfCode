package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Trait struct {
	name     string
	quantity int
}

func Day16(input string, part int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	aunts := map[string][]Trait{}
	//part 1 is simple string matching
	rules := []string{
		"children: 3",
		"cats: 7",
		"samoyeds: 2",
		"pomeranians: 3",
		"akitas: 0",
		"vizslas: 0",
		"goldfish: 5",
		"trees: 3",
		"cars: 2",
		"perfumes: 1",
	}
	//Part 2 needs the quantities
	rulesMap := map[string]int{}
	for _, v := range rules {
		split := strings.Split(v, ":")
		item := split[0]
		quantity, _ := strconv.Atoi(strings.TrimSpace(split[1]))
		rulesMap[item] = quantity
	}

	re1 := regexp.MustCompile(`Sue [0-9]{1,}`)
	re2 := regexp.MustCompile(`[a-z]*: [0-9]{1,}`)
	for _, l := range lines {
		aunt := re1.FindString(l)
		traitString := re2.FindAllString(l, -1)

		traits := []Trait{}
		for _, v := range traitString {
			split := strings.Split(v, ":")
			item := split[0]
			quantity, _ := strconv.Atoi(strings.TrimSpace(split[1]))
			traits = append(traits, Trait{item, quantity})
		}

		aunts[aunt] = traits
		if part == 1 {
			matches := 0
			for _, v := range traitString {
				if slices.Contains(rules, v) {
					matches++
				}
			}
			if matches == 3 {
				fmt.Printf("Day 16 Part1 - Aunt %s sent you the gift\n", aunt)
				break
			}
		}
	}

	if part == 2 {
		for k, v := range aunts {
			matches := 0
			for _, t := range v {
				if t.name == "cats" || t.name == "trees" {
					if t.quantity > rulesMap[t.name] {
						matches++
					}
					continue
				} else if t.name == "goldfish" || t.name == "pomeranians" {
					if t.quantity < rulesMap[t.name] {
						matches++
					}
					continue
				} else if t.quantity == rulesMap[t.name] {
					matches++
				}
				if matches == 3 {
					fmt.Printf("Day 16 Part 2 - Aunt %s sent you the gift\n", k)
					break
				}
			}
		}
	}

}
