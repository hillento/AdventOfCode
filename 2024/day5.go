package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5(input string) {
	split := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := strings.Split(strings.TrimSpace(split[0]), "\n")
	updates := strings.Split(strings.TrimSpace(split[1]), "\n")
	rejects := [][]string{}
	sum := 0
	sum2 := 0
	rulesMap := map[string][]string{}

	for _, r := range rules {
		rulesSplit := strings.Split(r, "|")
		rulesMap[rulesSplit[1]] = append(rulesMap[rulesSplit[1]], rulesSplit[0])
	}
	for _, u := range updates {
		safe := true
		updatesSplit := strings.Split(u, ",")
		l := len(updatesSplit)
		for i := 0; i < len(updatesSplit)-1; i++ {
			compSlice := updatesSplit[i+1 : l]
			if safePrint(compSlice, rulesMap[updatesSplit[i]]) == false {
				rejects = append(rejects, updatesSplit)
				safe = false
				break
			}
		}
		if safe == true {
			mid, _ := strconv.Atoi(updatesSplit[l/2])
			sum += mid
		}
	}
	fmt.Printf("Day 4.1: %v\n", sum)

	for _, r := range rejects {
		l := len(r)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if !safePrint([]string{r[i], r[j]}, rulesMap[r[i]]) {
					hold := r[i]
					r[i] = r[j]
					r[j] = hold
				}
			}
		}
		mid, _ := strconv.Atoi(r[l/2])
		sum2 += mid
	}
	fmt.Printf("Day 4.2: %v\n", sum2)
}

func safePrint(s1 []string, s2 []string) bool {
	for _, el1 := range s1 {
		if slices.Contains(s2, el1) {
			return false
		}
	}
	return true
}
