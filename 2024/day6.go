package main

import (
	"fmt"
	"strings"
)

func Day6_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for i := range mat {
		for j, w := range mat[i] {
			if string(w) == "^" {
				guardRow = i
				guardCol = j
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visited := make(map[[2]int]bool)

	for {
		visited[[2]int{guardRow, guardCol}] = true
		currentDirection := directions[guardDirection]

		nextGuardRow := guardRow + currentDirection[0]
		nextGuardCol := guardCol + currentDirection[1]

		if nextGuardRow < 0 || nextGuardRow >= len(mat) || nextGuardCol < 0 || nextGuardCol >= len(mat[0]) {
			break
		}

		if mat[nextGuardRow][nextGuardCol] == '#' {
			guardDirection = (guardDirection + 1) % 4
			currentDirection = directions[guardDirection]
			nextGuardRow = guardRow + currentDirection[0]
			nextGuardCol = guardCol + currentDirection[1]

			if nextGuardRow < 0 || nextGuardRow >= len(mat) || nextGuardCol < 0 || nextGuardCol >= len(mat[0]) {
				break
			}
		}
		guardRow = nextGuardRow
		guardCol = nextGuardCol
	}
	fmt.Println(len(visited))
}

func Day6_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))
	loops := 0

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	guardRow := -1
	guardCol := -1
	guardDirection := 0

	for i := range mat {
		for j, w := range mat[i] {
			if string(w) == "^" {
				guardRow = i
				guardCol = j
				break
			}
		}
	}

	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			if mat[row][col] != '.' {
				continue
			}

			mat[row][col] = '#'
			visited := make(map[[3]int]bool)
			currRow := guardRow
			currCol := guardCol
			currDir := guardDirection
			loop := false

			for {
				guardLoc := [3]int{currRow, currCol, currDir}

				if visited[guardLoc] {
					loop = true
					break
				}
				visited[guardLoc] = true

				nextGuardRow := currRow + directions[currDir][0]
				nextGuardCol := currCol + directions[currDir][1]

				if nextGuardRow < 0 || nextGuardRow >= len(mat) || nextGuardCol < 0 || nextGuardCol >= len(mat[0]) {
					break
				}

				if mat[nextGuardRow][nextGuardCol] == '#' {
					currDir = (currDir + 1) % 4
				} else {
					currRow = nextGuardRow
					currCol = nextGuardCol
				}
			}
			if loop {
				loops++
			}
			mat[row][col] = '.'
		}
	}
	fmt.Println(loops)
}
