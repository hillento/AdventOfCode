package main

import (
	"fmt"
	"strings"
)

func Day10(input string, itr int) {
	in := strings.TrimSpace(input)

	for i := 0; i < itr; i++ {
		var count int
		var seq strings.Builder

		for i := 0; i < len(in); i++ {
			if i == len(in)-1 || in[i] != in[i+1] {
				// add to seen
				seq.WriteString(fmt.Sprintf("%d%s", count+1, string(in[i])))
				count = 0
			} else {
				// build up running count
				count++
			}

		}
		in = seq.String()
	}

	fmt.Printf("Day 10 - %d LookSay transformations: %d\n", itr, len(in))
}
