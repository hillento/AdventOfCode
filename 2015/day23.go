package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	Register    string
	Instruction string
	Jump        int
}

func Day23(input string, part int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	regA := 0
	regB := 0

	if part == 2 {
		regA = 1
	}

	for i := 0; i < len(lines); {
		instruction := strings.Split(lines[i], " ")
		registerFlag := strings.Trim(instruction[1], ",")
		var reg int
		var jump int
		if registerFlag == "a" {
			reg = regA
		} else if registerFlag == "b" {
			reg = regB
		}
		cmd := instruction[0]

		switch {
		case cmd == "hlf":
			reg /= 2
		case cmd == "tpl":
			reg *= 3
		case cmd == "inc":
			reg++
		case cmd == "jmp":
			jump, _ = strconv.Atoi(strings.Trim(instruction[1], "+-"))
			if strings.Contains(instruction[1], "-") {
				i -= jump
				continue
			} else {
				i += jump
				continue
			}
		case cmd == "jie" && reg%2 == 0:
			jump, _ = strconv.Atoi(strings.Trim(instruction[2], "+"))
			i += jump
			continue
		case cmd == "jio" && reg == 1:
			jump, _ = strconv.Atoi(strings.Trim(instruction[2], "+"))
			i += jump
			continue
		}

		if registerFlag == "a" {
			regA = reg
		} else if registerFlag == "b" {
			regB = reg
		}
		i++

	}
	fmt.Printf("Day 22 Part %d - Register A: %d Register B: %d\n", part, regA, regB)
}
