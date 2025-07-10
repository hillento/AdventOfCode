package main

import (
	"fmt"
	"strings"
)

func Day18(input string, steps int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))

	ON := '#'
	OFF := '.'

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	totalOn := 0
	for range 1 {
		newMat := mat
		newMatOn := 0

		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				fmt.Printf("Coordinate: (%d, %d) is a %s\n", i, j, string(mat[i][j]))
				if j != 0 {
					fmt.Printf("Coordinate: (%d, %d) is a %s\n", i, j-1, string(mat[i][j-1]))
				}
				onCount := 0

				if i != 0 {
					if j != 0 && mat[i-1][j-1] == ON {
						fmt.Println(i-1, j-1)
						onCount++
					}
					if mat[i-1][j] == ON {
						fmt.Println(i-1, j)
						onCount++
					}
					if j != len(mat[0])-1 && mat[i-1][j+1] == ON {
						fmt.Println(i-1, j+1)
						onCount++
					}
				}

				if j != 0 && mat[i][j-1] == ON {
					fmt.Println(i, j-1)
					onCount++
				}

				if j != len(mat[0])-1 && mat[i][j+1] == ON {
					fmt.Println(i, j+1)
					onCount++
				}

				if i != len(mat)-1 {
					if j != 0 && mat[i+1][j-1] == ON {
						fmt.Println(i+1, j-1)
						onCount++
					}
					if mat[i+1][j] == ON {
						fmt.Println(i+1, j)
						onCount++
					}
					if j != len(mat[0])-1 && mat[i+1][j+1] == ON {
						fmt.Println(i+1, j+1)
						onCount++
					}
				}
				if onCount == 3 && mat[i][j] == OFF {
					newMat[i][j] = '#'
					newMatOn++
				} else if onCount == 2 && mat[i][j] == ON {
					newMat[i][j] = '#'
				} else if onCount == 3 && mat[i][j] == ON {
					newMat[i][j] = '#'
				} else {
					newMat[i][j] = '.'
				}
				fmt.Println("Neighbors on: ", onCount)
			}
		}
		mat = newMat
		// for _, v := range mat {
		// 	fmt.Println(string(v))
		// }
		// fmt.Println(totalOn)
		totalOn = newMatOn
	}
	for _, v := range mat {
		fmt.Println(string(v))
	}
	fmt.Println(totalOn)
}
