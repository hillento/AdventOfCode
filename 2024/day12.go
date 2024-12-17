package main

import (
	"fmt"
	"strings"
)

func Day12_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	directions := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	cost := 0

	visited := make(map[[2]int]bool)

	var gardenMapper func(row int, col int, plant rune) int
	gardenMapper = func(row int, col int, plant rune) int {
		area := 0
		perimeter := 0
		q := [][2]int{{row, col}}
		visited[[2]int{row, col}] = true

		for len(q) > 0 {
			currPlant := q[0]
			q = q[1:]
			area++

			for _, dir := range directions {
				newRow := currPlant[0] + dir[0]
				newCol := currPlant[1] + dir[1]

				if newRow < 0 || newCol < 0 || newRow >= len(mat) || newCol >= len(mat[newRow]) {
					perimeter++
					continue
				}

				if mat[newRow][newCol] != plant {
					perimeter++
					continue
				}

				if !visited[[2]int{newRow, newCol}] {
					q = append(q, [2]int{newRow, newCol})
					visited[[2]int{newRow, newCol}] = true
				}
			}
		}
		return area * perimeter
	}
	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if !visited[[2]int{row, col}] {
				cost += gardenMapper(row, col, mat[row][col])
			}
		}
	}
	fmt.Println("Day 12.1: ", cost)
}

func Day12_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	output := 0

	directions := [][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	outerCorners := []int{
		0, 0, 0, 1, 0, 0, 1,
		2, 0, 1, 0, 2,
		1, 2, 2, 4,
	}

	checkInnerCorners := [][][]int{
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
		{{1, -1}, {1, 1}},
		{{-1, 1}, {1, 1}},
		{{1, 1}},
		{{-1, -1}, {-1, 1}},
		{},
		{{-1, 1}},

		{},
		{{-1, -1}, {1, -1}},
		{{1, -1}},
		{},
		{},

		{{-1, -1}},
		{},
		{},
		{},
	}

	visited := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if visited[[2]int{row, col}] {
				continue
			}
			plant := grid[row][col]
			area := 0
			corners := 0

			queue := [][2]int{{row, col}}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				currentRow, currentCol := current[0], current[1]
				if visited[[2]int{currentRow, currentCol}] {
					continue
				}
				visited[[2]int{currentRow, currentCol}] = true

				area++
				cornerType := 0

				for i, dir := range directions {
					newRow := currentRow + dir[0]
					newCol := currentCol + dir[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						cornerType += (1 << i)
					} else if grid[newRow][newCol] != plant {
						cornerType += (1 << i)
					} else if !visited[[2]int{newRow, newCol}] {
						queue = append(queue, [2]int{newRow, newCol})
					}
				}

				outerCornerCount := outerCorners[cornerType]
				innerCornerCount := 0
				for _, corner := range checkInnerCorners[cornerType] {
					newRow := currentRow + corner[0]
					newCol := currentCol + corner[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						continue
					} else if grid[newRow][newCol] != plant {
						innerCornerCount++
					}
				}
				corners += outerCornerCount + innerCornerCount
			}

			price := area * corners
			output += price
		}
	}

	fmt.Println("Output Day 12 Part 2", output)
}
