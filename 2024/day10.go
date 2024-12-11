package main

import (
	"fmt"
	"strings"
)

func Day10_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]int, len(lines))

	for i, line := range lines {
		mat[i] = make([]int, len(lines))

		for j, v := range line {
			mat[i][j] = int(v - '0')
		}
	}

	directions := [][2]int{
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //left
		{0, -1}, //down
	}

	//need to define this way so it can call itself
	var trailRun func(row int, col int, height int, trailEnds map[[2]int]bool)

	trailRun = func(row int, col int, height int, trailEnds map[[2]int]bool) {
		if row < 0 || col < 0 || row >= len(mat) || col >= len(mat[0]) || mat[row][col] != height {
			return
		}

		currPos := [2]int{row, col}

		if height == 9 {
			trailEnds[currPos] = true
			return
		}

		for _, d := range directions {
			nextR := row + d[0]
			nextC := col + d[1]
			trailRun(nextR, nextC, height+1, trailEnds)
		}
	}

	score := 0

	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			trailEnds := make(map[[2]int]bool)
			trailRun(row, col, 0, trailEnds)

			score += len(trailEnds)
		}
	}
	fmt.Println("Day 10.1: ", score)
}

func Day10_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]int, len(lines))

	for i, line := range lines {
		mat[i] = make([]int, len(lines))

		for j, v := range line {
			mat[i][j] = int(v - '0')
		}
	}

	directions := [][2]int{
		{-1, 0}, //up
		{0, 1},  //right
		{1, 0},  //left
		{0, -1}, //down
	}

	//need to define this way so it can call itself
	var trailRun func(row int, col int, height int) int

	trailRun = func(row int, col int, height int) int {
		if row < 0 || col < 0 || row >= len(mat) || col >= len(mat[0]) || mat[row][col] != height {
			return 0
		}

		if height == 9 {
			return 1
		}

		trails := 0

		for _, d := range directions {
			nextR := row + d[0]
			nextC := col + d[1]
			trails += trailRun(nextR, nextC, height+1)

		}
		return trails
	}

	rating := 0

	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if mat[row][col] == 0 {
				rating += trailRun(row, col, 0)
			}
		}
	}
	fmt.Println("Day 10.2: ", rating)
}
