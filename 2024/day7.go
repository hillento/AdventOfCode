package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	for i := 0; i < len(lines); i++ {
		splits := strings.Split(lines[i], ":")
		candidate, _ := strconv.Atoi(splits[0])
		values := strings.Split(strings.TrimSpace(splits[1]), " ")
		nums := []int{}

		for _, v := range values {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}

		chain := [][]int{}
		chain = append(chain, []int{nums[0]})

		for j := 1; j < len(nums); j++ {
			l := len(chain) - 1
			link := []int{}
			for k := 0; k < len(chain[l]); k++ {
				add := nums[j] + chain[l][k]
				mul := nums[j] * chain[l][k]
				link = append(link, add)
				link = append(link, mul)
			}
			chain = append(chain, link)
		}
		finLink := len(chain) - 1
		for _, link := range chain[finLink] {
			if link == candidate {
				output += candidate
				break
			}
		}
	}
	fmt.Println("Day 7.1: ", output)
}

func Day7_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0

	for i := 0; i < len(lines); i++ {
		splits := strings.Split(lines[i], ":")
		candidate, _ := strconv.Atoi(splits[0])
		values := strings.Split(strings.TrimSpace(splits[1]), " ")
		nums := []int{}

		for _, v := range values {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}

		chain := [][]int{}
		chain = append(chain, []int{nums[0]})

		for j := 1; j < len(nums); j++ {
			l := len(chain) - 1
			link := []int{}
			for k := 0; k < len(chain[l]); k++ {
				add := nums[j] + chain[l][k]
				mul := nums[j] * chain[l][k]
				str1 := strconv.Itoa(chain[l][k])
				str2 := strconv.Itoa(nums[j])
				str := str1 + str2
				con, _ := strconv.Atoi(str)
				link = append(link, add)
				link = append(link, mul)
				link = append(link, con)
			}
			chain = append(chain, link)
		}
		finLink := len(chain) - 1
		for _, link := range chain[finLink] {
			if link == candidate {
				output += candidate
				break
			}
		}
	}
	fmt.Println("Day 7.2: ", output)
}
