package main

import (
	"strings"
)

func Day8_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	antennas := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := range mat {
		for col := range mat[row] {
			if mat[row][col] != '.' {
				antennas[mat[row][col]] = append(antennas[mat[row][col]], [2]int{row, col})
			}
		}
	}

	for _, pos := range antennas {
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				pos1 := pos[i]
				pos2 := pos[j]

				dRow := pos2[0] - pos1[0]
				dCol := pos2[1] - pos1[1]

				antiNode1 := [2]int{pos1[0] - dRow, pos1[1] - dCol}
				antiNode2 := [2]int{pos2[0] + dRow, pos2[1] + dCol}

				if antiNode1[0] >= 0 && antiNode1[0] < len(mat) && antiNode1[1] >= 0 && antiNode1[1] < len(mat[0]) {
					antiNodes[antiNode1] = true
				}
				if antiNode2[0] >= 0 && antiNode2[0] < len(mat) && antiNode2[1] >= 0 && antiNode2[1] < len(mat[0]) {
					antiNodes[antiNode2] = true
				}
			}
		}
	}
	println("Day8.1: ", len(antiNodes))
}

func Day8_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mat := make([][]rune, len(lines))

	for i, line := range lines {
		mat[i] = []rune(line)
	}

	antennas := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := range mat {
		for col := range mat[row] {
			if mat[row][col] != '.' {
				antennas[mat[row][col]] = append(antennas[mat[row][col]], [2]int{row, col})
			}
		}
	}

	for _, pos := range antennas {
		for i := 0; i < len(pos); i++ {
			for j := i + 1; j < len(pos); j++ {
				pos1 := pos[i]
				pos2 := pos[j]

				dRow := pos2[0] - pos1[0]
				dCol := pos2[1] - pos1[1]

				antiNode1Row := pos1[0]
				antiNode1Col := pos1[1]

				for antiNode1Row >= 0 && antiNode1Row < len(mat) && antiNode1Col >= 0 && antiNode1Col < len(mat) {
					antiNodes[[2]int{antiNode1Row, antiNode1Col}] = true

					antiNode1Row -= dRow
					antiNode1Col -= dCol
				}

				antiNode2Row := pos2[0]
				antiNode2Col := pos2[1]

				for antiNode2Row >= 0 && antiNode2Row < len(mat[0]) && antiNode2Col >= 0 && antiNode2Col < len(mat[0]) {
					antiNodes[[2]int{antiNode2Row, antiNode2Col}] = true

					antiNode2Row += dRow
					antiNode2Col += dCol
				}
			}
		}
	}

	println("Day8.2: ", len(antiNodes))
}
