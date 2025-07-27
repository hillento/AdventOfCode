package fifteen

import (
	"fmt"
	"strings"
)

func Day18(input string, steps int, part int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	ON := '#'
	OFF := '.'

	mat := make([][]rune, len(lines))
	for i, line := range lines {
		mat[i] = []rune(line)
	}

	if part == 2 {
		mat[0][0] = '#'
		mat[len(lines[0])-1][0] = '#'
		mat[len(lines[0])-1][len(lines[0])-1] = '#'
		mat[0][len(lines[0])-1] = '#'
	}

	totalOn := 0
	for range steps {
		newMat := make([][]rune, len(lines))
		for i, line := range lines {
			newMat[i] = []rune(line)
		}
		newMatOn := 0

		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				cell := mat[i][j]
				onCount := 0

				if i != 0 {
					if j != 0 && mat[i-1][j-1] == ON {
						onCount++
					}
					if mat[i-1][j] == ON {
						onCount++
					}
					if j != len(mat[0])-1 && mat[i-1][j+1] == ON {
						onCount++
					}
				}

				if j != 0 && mat[i][j-1] == ON {
					onCount++
				}

				if j != len(mat[0])-1 && mat[i][j+1] == ON {
					onCount++
				}

				if i != len(mat)-1 {
					if j != 0 && mat[i+1][j-1] == ON {
						onCount++
					}
					if mat[i+1][j] == ON {
						onCount++
					}
					if j != len(mat[0])-1 && mat[i+1][j+1] == ON {
						onCount++
					}
				}
				if part == 2 && ((i == 0 && j == 0) || (i == len(mat)-1 && j == 0) || (i == 0 && j == len(mat)-1) || (i == len(mat)-1 && j == len(mat)-1)) {
					newMat[i][j] = ON
					newMatOn++
				} else if onCount == 3 && cell == OFF {
					newMat[i][j] = ON
					newMatOn++
				} else if cell == ON && (onCount == 3 || onCount == 2) {
					newMat[i][j] = ON
					newMatOn++
				} else if cell == ON {
					newMat[i][j] = OFF
				} else {
					newMat[i][j] = mat[i][j]
					if cell == ON {
						newMatOn++
					}
				}
			}
		}
		mat = newMat
		totalOn = newMatOn
	}
	fmt.Printf("Day 18 part %d - Number of lights on after 100 iterations: %d\n", part, totalOn)
}
