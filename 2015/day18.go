package main

import (
	"fmt"
	"strings"
)

func Day18(input string, steps int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]string, len(lines))

	for i, line := range lines {
		mat[i] = strings.Split(line,`''`)
	}

	totalOn := 0
	for range steps {
		newMat := mat
		newMatOn := 0

		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				onCount := 0
				if i != 0 {
					if j != 0 {
						if mat[i-1][j-1] == "#" {
							onCount++
						}
					}
					if mat[i-1][j] == "#" {
						onCount++
					}
					if j != len(mat[0]) -1 {
						if mat[i-1][j+1] == "#" {
							onCount++
						}
					}
				}

				if j != 0 {
					if mat[i][j-1] == "#" {
						onCount++
					}
				}

				if j != len(mat[0]) -1 {
					if mat[i][j+1] == "#" {
						onCount++
					}
				}

				if i != len(mat) -1 {
					if j != 0 {
						if mat[i+1][j-1] == "#" {
							onCount++
						}
					}
					if mat[i+1][j] == "#" {
						onCount++
					}
					if j != len(mat[0]) - 1 {
						if mat[i+1][j+1] == "#" {
							onCount++
						}
					}
				}
				if onCount == 3 && mat[i][j] == "." {
					newMat[i][j] = "#"
					newMatOn++
				} else if onCount == 2 && mat[i][j] == "#" {
					newMat[i][j] = "."
				} else if onCount == 3 && mat[i][j] == "#" {
					newMat[i][j] = "."
				}
			}
		}
		mat = newMat
		for _, v := range mat {
			fmt.Println(v)
		}
		fmt.Println(totalOn)
		totalOn = newMatOn
	}
	for _, v := range mat {
		fmt.Println(v)
	}
	fmt.Println(totalOn)
}
