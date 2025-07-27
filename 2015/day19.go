package fifteen

//Goal is 576

import (
	"fmt"
	"strings"
	"unicode"
)

func Day19_1(input string) {
	rules, startingM := getRules(input)
	newMolecules := map[string]bool{}

	for k, v := range rules {
		found := strings.Count(startingM, k)
		for _, replace := range v {
			if found > 0 {
				for i := range found {
					newMole := strings.TrimSpace(replaceNthSubstring(startingM, k, replace, i+1))
					newMolecules[newMole] = true
				}
			}
		}
	}
	fmt.Printf("Day 19 - Part 1: The number of unique molecules: %d\n", len(newMolecules))
}

//I was not successful at brute forcing part 2.
//In the end this was solved by reading the reddit and learning about _unambiguous language_
//https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/
//Doubt I could have come up with a reliable solution on my own

func Day19_2(input string) {
	_, molecule := getRules(input)
	elements := getElements(molecule)

	totalCount := 0
	parenCount := 0
	commaCount := 0
	fmt.Println(elements)
	for e, c := range elements {
		totalCount += c
		if e == "Rn" || e == "Ar" {
			parenCount += c
		}
		if e == "Y" {
			commaCount += c
		}
	}

	solution := totalCount - parenCount - (2 * commaCount) - 1

	fmt.Printf("Day 19 Part 2 - The least number of steps to reduce reporduce the molecule is: %d\n", solution)
}

func getRules(input string) (rulesMap map[string][]string, startMolecule string) {
	rulesMap = map[string][]string{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	molecule := lines[len(lines)-1]
	rules := lines[:len(lines)-2]

	for _, rule := range rules {
		ruleParts := strings.Split(rule, " => ")
		find := strings.TrimSpace(ruleParts[0])
		replace := strings.TrimSpace(ruleParts[1])
		rulesMap[find] = append(rulesMap[find], replace)
	}
	return rulesMap, molecule
}

func replaceNthSubstring(s, old, newString string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + newString + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func getElements(input string) map[string]int {
	elements := map[string]int{}
	l := 0
	for i := input; i != ""; i = i[l:] {
		l = strings.IndexFunc(i[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(i)
		}
		_, ok := elements[i[:l]]
		if ok {
			elements[i[:l]]++
		} else {
			elements[i[:l]] = 1
		}
	}
	return elements
}
