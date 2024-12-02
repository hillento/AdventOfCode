package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0

	var left []int
	var right []int

	for _, s := range lines {
		parts := strings.Fields(s)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		d := left[i] - right[i]

		if d < 0 {

			d *= -1
		}
		sum += d
	}

	fmt.Printf("Day 1_1 answer: %v\n", sum)
}

func Day1_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	similarity := 0

	var left []int
	var right []int

	for _, s := range lines {
		parts := strings.Fields(s)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	sort.Ints(left)
	sort.Ints(right)

	rightMap := make(map[int]int)

	for _, v := range right {
		rightMap[v]++
	}
	for _, v := range left {
		similarity += v * rightMap[v]
	}
	fmt.Printf("Day 1_2 answer: %v\n", similarity)
}
