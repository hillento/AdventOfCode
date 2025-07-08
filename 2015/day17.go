package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// GLOBALS FOR PART 2
var MINCOUNT int = 1000
var COUNTS = map[int]int{}

func Day17(input string, part int) {
	nums := []int{}
	for _, l := range strings.Split(strings.TrimSpace(input), "\n") {
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if part == 1 {
		storageCount := Part1(nums, 150)
		fmt.Printf("Day 17 Part 1 - The number of ways the containers can store 150 liters of eggnog is: %d\n", storageCount)
	}
	if part == 2 {
		Part2(nums, 150, 0)
		fmt.Printf("Day 17 Part 2 - The number of ways the minimum number of containers can store 150 liters of eggnog is: %d\n", COUNTS[MINCOUNT])
	}

}

func Part1(containers []int, goal int) int {
	//If the goal ends up negative this isn't a valid combo
	if goal < 0 {
		return 0
	}
	//if the goal is exactly 0 then a working combo has been found, add 1
	if goal == 0 {
		return 1
	}
	//if we have run out of containers and there is still eggnog, it's not valid
	if len(containers) == 0 && goal > 0 {
		return 0
	}

	//There are 2 options, use the first container or don't. the recursive method will ensure all options are checked
	//Both calls remove the first container
	//First one assumes the first container in the current set was used
	//Second call assumes the first container in the current set was discarded
	return Part1(containers[1:], goal-containers[0]) + Part1(containers[1:], goal)

}

func Part2(containers []int, goal int, count int) int {
	if goal < 0 {
		return 0
	}
	//The combo was valid, check if the count has already been found.
	//If the count has already been found, increment it, if not add it with count of 1
	if goal == 0 {
		MINCOUNT = min(count, MINCOUNT)
		_, ok := COUNTS[count]
		if ok {
			COUNTS[count] += 1
		} else {
			COUNTS[count] = 1
		}
		return 1
	}
	if len(containers) == 0 && goal > 0 {
		return 0
	}

	return Part2(containers[1:], goal-containers[0], count+1) + Part2(containers[1:], goal, count)
}
