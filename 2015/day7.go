package fifteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day7(input string) {
	wireToRule := map[string]string{}

	for _, inst := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(inst, " -> ")
		wireToRule[parts[1]] = parts[0]
	}

	//Start at a and work back to save time
	v := findValue(wireToRule, "a", map[string]int{})

	fmt.Printf("Day 7.1 - Value of wire a: %d\n", v)

	wireToRule["b"] = strconv.Itoa(v)
	v = findValue(wireToRule, "a", map[string]int{})
	fmt.Printf("Day 7.2 - Value of wire a: %d\n", v)
}

// use a recursive memoized structire
func findValue(inst map[string]string, wire string, memo map[string]int) int {
	//if we already have the value just return it
	if memoVal, ok := memo[wire]; ok {
		return memoVal
	}

	//if it's a number just return that value
	if v, err := strconv.Atoi(wire); err == nil {
		return v
	}

	rule := inst[wire]
	ruleParts := strings.Split(rule, " ")

	var result int
	switch {
	case len(ruleParts) == 1: //this is for an assignment without logic
		result = findValue(inst, ruleParts[0], memo)
	case ruleParts[0] == "NOT":
		initial := findValue(inst, ruleParts[1], memo)
		result = (math.MaxUint16) ^ initial
	case ruleParts[1] == "AND":
		result = findValue(inst, ruleParts[0], memo) & findValue(inst, ruleParts[2], memo)
	case ruleParts[1] == "OR":
		result = findValue(inst, ruleParts[0], memo) | findValue(inst, ruleParts[2], memo)
	case ruleParts[1] == "LSHIFT":
		result = findValue(inst, ruleParts[0], memo) << findValue(inst, ruleParts[2], memo)
	case ruleParts[1] == "RSHIFT":
		result = findValue(inst, ruleParts[0], memo) >> findValue(inst, ruleParts[2], memo)
	}

	memo[wire] = result
	return result
}
