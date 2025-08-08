package sixteen

import (
	"fmt"
	"strings"
)

// Keypad for part1
var keypad1 = [][]string{
	{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"},
}

// Better keypad
var keypad2 = [][]string{
	{"0", "0", "1", "0", "0"}, {"0", "2", "3", "4", "0"}, {"5", "6", "7", "8", "9"}, {"0", "A", "B", "C", "0"}, {"0", "0", "D", "0", "0"},
}

func Day2(input string, part int) {
	keypadMap := make(map[[2]int]string)
	var keypad [][]string
	var code []string
	var current [2]int

	if part == 1 {
		keypad = keypad1
		current = [2]int{2, 2} //5
	} else {
		keypad = keypad2
		current = [2]int{3, 0} //5
	}

	//Maps the keys to the unique coordinate of the keypad
	for i, v := range keypad {
		for j, s := range v {
			if s != "0" {
				keypadMap[[2]int{i, j}] = s
			}
		}
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, v := range lines {
		for _, r := range v {
			//changes the coordinate accordingly if the new one would hit a key
			switch r {
			case 'U':
				_, prs := keypadMap[[2]int{current[0] - 1, current[1]}]
				if prs {
					current[0]--
				}
			case 'R':
				_, prs := keypadMap[[2]int{current[0], current[1] + 1}]
				if prs {
					current[1]++
				}
			case 'D':
				_, prs := keypadMap[[2]int{current[0] + 1, current[1]}]
				if prs {
					current[0]++
				}
			case 'L':
				_, prs := keypadMap[[2]int{current[0], current[1] - 1}]
				if prs {
					current[1]--
				}
			}
		}
		//once the key is found, add it the code
		code = append(code, keypadMap[current])
	}
	fmt.Printf("Day 2 part %d - The bathroom code is %v\n", part, code)
}

