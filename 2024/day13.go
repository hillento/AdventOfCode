package main

import (
	"fmt"
	"strings"
)

//Button A 3 Coins
//Button B 1 Coin

func Day13_1(input string) {
	machines := strings.Split(input, "\n\n")
	sum := 0

	for _, machine := range machines {
		machine = strings.TrimSpace(machine)
		var ax, ay, bx, by, px, py int
		lines := strings.Split(machine, "\n")
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)

		//Carmer's rule
		dt := ax*by - ay*bx

		if dt == 0 {
			continue
		}

		a := (px*by - py*bx) / dt
		b := (ax*py - ay*px) / dt

		if a >= 0 && b >= 0 && ax*a+bx*b == px && ay*a+by*b == py {
			cost := a*3 + b
			sum += cost
		}
	}
	fmt.Println("Day 13: ", sum)
}

func Day13_2(input string) {
	machines := strings.Split(input, "\n\n")
	sum := 0

	for _, machine := range machines {
		machine = strings.TrimSpace(machine)
		var ax, ay, bx, by, px, py int
		lines := strings.Split(machine, "\n")
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)

		px += 10000000000000
		py += 10000000000000

		//Carmer's rule
		dt := ax*by - ay*bx

		if dt == 0 {
			continue
		}

		a := (px*by - py*bx) / dt
		b := (ax*py - ay*px) / dt

		if a >= 0 && b >= 0 && ax*a+bx*b == px && ay*a+by*b == py {
			cost := a*3 + b
			sum += cost
		}
	}
	fmt.Println("Day 13: ", sum)
}
