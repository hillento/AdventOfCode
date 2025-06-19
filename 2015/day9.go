package main

import (
	"fmt"
	"github.com/etnz/permute"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day9(input string) {
	r := make(map[[2]string]int)
	places := []string{}
	shortest := math.MaxUint16
	longest := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " ")
		dist, _ := strconv.Atoi(parts[4])
		r[[2]string{parts[0], parts[2]}] = dist
		if !slices.Contains(places, parts[0]) {
			places = append(places, parts[0])
		}
		if !slices.Contains(places, parts[2]) {
			places = append(places, parts[2])
		}
	}
	perms := permute.Permutations(places)

	for _, perm := range perms {
		sum := 0
		for i := 0; i < len(perm)-1; i++ {
			val1 := [2]string{perm[i], perm[i+1]}
			val2 := [2]string{perm[i+1], perm[i]}
			v1, ok1 := r[val1]
			v2, ok2 := r[val2]
			if ok1 {
				sum += v1
			}
			if ok2 {
				sum += v2
			}
		}
		if sum < shortest {
			shortest = sum
		}
		if sum > longest {
			longest = sum
		}
	}
	fmt.Println("Day9 Part 1 - Shortest route: ", shortest)
	fmt.Println("Day9 Part 2 - Longest route: ", longest)
}
