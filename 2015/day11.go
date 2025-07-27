package fifteen

import (
	"fmt"
	"regexp"
	"strings"
)

// a: 97
// b: 122
func Day11(input string, part int) {
	candidate := []rune(strings.TrimSpace(input))
	count := 0
	for {
		i := len(candidate) - 1
		if candidate[i] == 122 {
			for candidate[i] == 122 {
				i--
			}
		}
		candidate[i] += 1
		if checkString(string(candidate)) {
			count++
			if part == count {
				break
			}
			candidate[i] += 1
		}
		for i < len(candidate)-1 {
			i++
			candidate[i] = 97
		}
	}
	fmt.Printf("Day 11  part %d- New Password is: %s\n", part, string(candidate))
}

func checkString(input string) bool {
	r1 := false
	r2 := false
	r3 := false

	r2reg := regexp.MustCompile(`i|o|l`)
	r2 = !r2reg.MatchString(input)
	repeat := 0
	for i := 0; i < len(strings.TrimSpace(input))-1; i++ {
		if input[i] == input[i+1] {
			repeat++
			i++
		}
		if repeat == 2 {
			r3 = true
			break
		}
	}
	for i := 0; i < len(strings.TrimSpace(input))-2; i++ {
		if input[i]+1 == input[i+1] && input[i]+2 == input[i+2] {
			r1 = true
			break
		}
	}
	return (r1 && r2 && r3)
}
