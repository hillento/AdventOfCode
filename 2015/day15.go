package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

func Day15(input string) {
	ings := [][]int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, l := range lines {
		parts := strings.Split(l, " ")
		var capacity, durability, flavor, texture, calories int
		capacity, _ = strconv.Atoi(strings.Trim(parts[2], ","))
		durability, _ = strconv.Atoi(strings.Trim(parts[4], ","))
		flavor, _ = strconv.Atoi(strings.Trim(parts[6], ","))
		texture, _ = strconv.Atoi(strings.Trim(parts[8], ","))
		calories, _ = strconv.Atoi(strings.Trim(parts[10], ","))
		fmt.Println(capacity, durability, flavor, texture, calories)

		ings = append(ings, []int{capacity, durability, flavor, texture, calories})
	}

	highScore := 0
	bestLowCal := 0
	for in1 := 0; in1 < 100; in1++ {
		for in2 := 0; in2+in1 < 100; in2++ {
			for in3 := 0; in3+in2+in1 < 100; in3++ {
				in4 := 100 - in1 - in2 - in3

				cap := ings[0][0]*in1 + ings[1][0]*in2 + ings[2][0]*in3 + ings[3][0]*in4
				dur := ings[0][1]*in1 + ings[1][1]*in2 + ings[2][1]*in3 + ings[3][1]*in4
				fla := ings[0][2]*in1 + ings[1][2]*in2 + ings[2][2]*in3 + ings[3][2]*in4
				tex := ings[0][3]*in1 + ings[1][3]*in2 + ings[2][3]*in3 + ings[3][3]*in4
				cal := ings[0][4]*in1 + ings[1][4]*in2 + ings[2][4]*in3 + ings[3][4]*in4
				if cap < 0 {
					cap = 0
				}
				if dur < 0 {
					dur = 0
				}
				if fla < 0 {
					fla = 0
				}
				if tex < 0 {
					tex = 0
				}

				score := cap * dur * fla * tex
				if cal == 500 && bestLowCal < score {
					bestLowCal = score
				}
				if score > highScore {
					highScore = score
				}
			}
		}
	}
	fmt.Printf("Day15 Part 1 - The best cookie recipe score is: %d\n", highScore)
	fmt.Printf("Day15 Part 2 - The best low cal cookie recipe score is: %d\n", bestLowCal)

}
