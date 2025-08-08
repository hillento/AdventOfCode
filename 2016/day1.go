package sixteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var dirs = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type Location struct {
	x int
	y int
}

func Day1(input string, part int) {
	tracker := make(map[int]int)
	locations := map[[2]int]bool{
		{0, 0}: true,
	}
	currentDir := 0
	tracker[0] = 0
	tracker[1] = 0

	values := strings.Split(strings.TrimSpace(input), ", ")
out:
	for _, v := range values {
		turn := v[0]
		dist, _ := strconv.Atoi(v[1:])

		if turn == 'R' {
			currentDir = (currentDir + 1) % 4
		} else {
			currentDir = (currentDir + 3) % 4
		}

		for range dist {
			tracker[0] += dirs[currentDir][0]
			tracker[1] += dirs[currentDir][1]

			Loc := [2]int{tracker[0], tracker[1]}
			_, prs := locations[Loc]
			if part == 2 && prs {
				break out
			}
			locations[Loc] = true
		}
	}
	distance := int(math.Abs(float64(tracker[0])) + math.Abs(float64(tracker[1])))
	fmt.Printf("Day 1 part %d - The number of blocks to travel is: %d\n", part, distance)
}
