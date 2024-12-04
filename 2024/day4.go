package main

import (
	"fmt"
	"strings"
)

func Day4_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m := [][]string{}
	count := 0

	for _, v := range lines {
		count += strings.Count(v, "XMAS")
		count += strings.Count(v, "SAMX")
		split := strings.Split(v, "")
		m = append(m, split)
	}
	t := transpose(m)
	for _, v := range t {
		count += strings.Count(strings.Join(v, ""), "XMAS")
		count += strings.Count(strings.Join(v, ""), "SAMX")
	}
	d := getDiagonals(m)
	for _, v := range d {
		count += strings.Count(strings.Join(v, ""), "XMAS")
		count += strings.Count(strings.Join(v, ""), "SAMX")
	}

	r := reverseRows(m)
	rd := getDiagonals(r)
	for _, v := range rd {
		count += strings.Count(strings.Join(v, ""), "XMAS")
		count += strings.Count(strings.Join(v, ""), "SAMX")
	}
	fmt.Printf("Day 4.1: %d\n", count)
}

func Day4_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	count := 0

	for i := 0; i < len(lines) -2; i++ {
		for j := 0; j < len(lines[i]) -2; j++ {
			if string(lines[i][j]) == "M" && string(lines[i][j+2]) == "M" && string(lines[i+1][j+1]) == "A"  && string(lines[i+2][j]) == "S" && string(lines[i+2][j+2]) == "S" {
				count++
			}
			if string(lines[i][j]) == "S" && string(lines[i][j+2]) == "S" && string(lines[i+1][j+1]) == "A"  && string(lines[i+2][j]) == "M" && string(lines[i+2][j+2]) == "M" {
				count++
			}
			if string(lines[i][j]) == "M" && string(lines[i][j+2]) == "S" && string(lines[i+1][j+1]) == "A"  && string(lines[i+2][j]) == "M" && string(lines[i+2][j+2]) == "S" {
				count++
			}
			if string(lines[i][j]) == "S" && string(lines[i][j+2]) == "M" && string(lines[i+1][j+1]) == "A"  && string(lines[i+2][j]) == "S" && string(lines[i+2][j+2]) == "M" {
				count++
			}
		}
	}
	fmt.Printf("Day 4.2: %d\n", count)
}

func transpose(mat [][]string) [][]string {
	transpose := make([][]string, len(mat[0]))
	for i := range transpose {
		transpose[i] = make([]string, len(mat))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			transpose[j][i] = mat[i][j]
		}
	}
	return transpose
}

func reverseRows(mat [][]string) [][]string {
	revs := [][]string{}
	for i := range mat {
		r := []string{}
		for j := range mat[i] {
		r = append(r, mat[i][len(mat)-1-j])
		}
		revs = append(revs, r)
	}
	return revs
}

func getDiagonals(mat [][]string) [][]string {
	diags := [][]string{}
	//diagonal left to righ going from the corner down
	for t := 0; t < len(mat); t++ {
		d := []string{}
		for i := 0; i < len(mat[0])-t; i++ {
			d = append(d, mat[i+t][i])
		}
		diags = append(diags, d)
	}
	//diagonal left to right going from corner + 1 to the right 
	for t := 1; t < len(mat[0]); t++ {
		d := [] string{}
		for i := 0; i < len(mat) - t; i++ {
			d = append(d, mat[i][i+t])
		}
		diags = append(diags, d)
	}
	return diags
}
