package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reg := regexp.MustCompile(`(mul\()[0-9]{1,3},[0-9]{1,3}(\))`)
	valid := []string{}
	sum := 0

	for _, v := range lines {
		valid = reg.FindAllString(v, -1)
		for _, v2 := range valid {
			split := strings.Split(v2, " ")
			rems := "mul()"
			result := strings.Map(func(r rune) rune {
				if strings.ContainsRune(rems, r) {
					return -1
				}
				return r
			}, split[0])
			nums := strings.Split(result, ",")
			digit, _ := strconv.Atoi(nums[0])
			digit2, _ := strconv.Atoi(nums[1])
			sum += digit * digit2
		}
	}
	fmt.Printf("Day 3.1 solution: %d\n", sum)
}

func Day3_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	reg1 := regexp.MustCompile(`(do\(\)|don\'t\(\))`)
	reg2 := regexp.MustCompile(`(mul\()[0-9]{1,3},[0-9]{1,3}(\))`)

	reg := regexp.MustCompile(reg1.String() + "|" + reg2.String())
	valid := []string{}
	enabled := true

	sum := 0

	for _, v := range lines {
		valid = reg.FindAllString(v, -1)
		for _, v2 := range valid {
			split := strings.Split(v2, " ")
			rems := "mul()"
			result := strings.Map(func(r rune) rune {
				if strings.ContainsRune(rems, r) {
					return -1
				}
				return r
			}, split[0])

			if result == "do" {
				enabled = true
			} else if result == "don't" {
				enabled = false
			}
			if enabled == true && result != "do" && result != "don't" {
				nums := strings.Split(result, ",")
				digit, _ := strconv.Atoi(nums[0])
				digit2, _ := strconv.Atoi(nums[1])
				sum += digit * digit2
			}
		}
	}
	fmt.Printf("Day 3.2 solution: %d\n", sum)
}
