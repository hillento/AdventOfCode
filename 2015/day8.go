package main

import (
	"fmt"
	"strings"
)

func Day8_1(input string) {
	var code, chars int

	for _, line := range strings.Split(input, "\n") {
		code += len(line)

		for i := 1; i < len(line)-1; i++ {
			if line[i] == '\\' {
				next := line[i+1]
				if next == '\\' || next == '"' {
					i++
				} else if next == 'x' {
					i += 3
				}
			}
			chars++
		}
	}
	fmt.Printf("Day8.1 - Total characters minues character in memory: %d\n", code-chars)
}

func Day8_2(input string) {
	var encodeLen, originLen int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		originLen += len(line)
		encodeLen += 2 //quotes at each end

		for i := 0; i < len(line); i++ {
			if line[i] == '\\' || line[i] == '"' {
				encodeLen += 2 //each of these will add a \ to the encoded plus their original character
			} else {
				encodeLen++
			}
		}
	}
	fmt.Printf("Day 8.2 - Encoded character minus string characters: %d\n", encodeLen-originLen)
}

