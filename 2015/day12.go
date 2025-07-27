package fifteen

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

func Day12(input string, part int) {
	var ans int
	if part == 1 {
		ans = Day12_1(input)
	} else {
		ans = Day12_2(input)
	}
	fmt.Printf("Day 12 Part %d - Value:  %d\n", part, ans)
}

func Day12_1(input string) int {
	sum := 0
	re := regexp.MustCompile("[-]?[0-9]+")
	nums := re.FindAllString(input, -1)
	for _, v := range nums {
		n, _ := strconv.Atoi(v)
		sum += n
	}
	return sum
}

func Day12_2(input string) int {
	if !regexp.MustCompile("[{}]").MatchString(input) ||
		!regexp.MustCompile("red").MatchString(input) {
		return Day12_1(input)
	}

	// try to parse into an object if that's
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(input), &obj)
	// not a json object, assume it's an array
	if err != nil {
		// parse into an array
		var arr []interface{}
		err := json.Unmarshal([]byte(input), &arr)
		if err != nil {
			panic(err)
		}

		var arrayTotal int
		for _, v := range arr {
			// marshal each array element into a string, then pass it back into part2
			str, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			arrayTotal += Day12_2(string(str))
		}

		return arrayTotal
	}

	// if any value in the object is "red" this object & its children RETURN ZERO
	for _, v := range obj {
		// have to convert interface into a string first
		str, ok := v.(string)
		if ok && str == "red" {
			return 0
		}
	}

	var total int
	for _, v := range obj {
		str, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		total += Day12_2(string(str))
	}

	return total
}
