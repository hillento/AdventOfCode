package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Since there is not change in input, or requirements, only 1 function
func Day11(input string) {

	stones := strings.Split(strings.TrimSpace(input), " ")

	pOne := setup(stones, 25)
	pTwo := setup(stones, 75)

	fmt.Println("Day 11.1: ", pOne)
	fmt.Println("Day 11.2: ", pTwo)
}

// just kicks everything off based on input and blink count
func setup(stones []string, blinks int) int {

	stoneCount := 0
	//Set up a cache so we don't have to recalculate if we know the result
	cache := make(map[string][]int)
	//Focus on 1 stone at a time
	for _, stone := range stones {
		stoneCount += getCount(stone, cache, blinks)
	}
	return stoneCount
}

func getCount(stone string, cache map[string][]int, blinks int) int {
	stoneCount := 0

	//if this stone has been found previously
	if _, ok := cache[stone]; ok {
		//and calculated how many stone there will be after a blink
		if cache[stone][blinks-1] != 0 {
			//just return what has already been calculated
			return cache[stone][blinks-1]
		}
	} else {
		//if the stone isn't found, setup a new array for up to 75 splits
		cache[stone] = make([]int, 75)
	}

	//on the last blnik, cache the value
	if blinks == 1 {
		cache[stone][blinks-1] = len(blink(stone))
		return len(blink(stone))
	}

	//make this function call itself reducing the number of blinks left until 1
	//stoneCount will compound with each run and the calls will expand range with each split
	for _, s := range blink(stone) {
		stoneCount += getCount(s, cache, blinks-1)
	}

	//cache the final result and return the count
	cache[stone][blinks-1] = stoneCount
	return stoneCount
}

// follow the rules outlined, nothing fancy
func blink(stone string) []string {
	newArr := []string{}
	if stone == "0" {
		newArr = append(newArr, "1")
	} else if len(stone)%2 == 0 {
		mid := len(stone) / 2
		str1 := stone[:mid]
		str2 := stone[mid:]

		str1 = strings.TrimLeft(str1, "0")
		str2 = strings.TrimLeft(str2, "0")

		if str1 == "" {
			str1 = "0"
		}
		if str2 == "" {
			str2 = "0"
		}
		newArr = append(newArr, str1)
		newArr = append(newArr, str2)
	} else {
		n, _ := strconv.Atoi(stone)
		n *= 2024
		s := strconv.Itoa(n)
		newArr = append(newArr, s)
	}
	return newArr
}
