package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

func Day25(input string, part int) {
	bits := strings.Split(strings.TrimSpace(input), " ")
	nums := []int{}

	//form the scrap, just start at the highest number, doesn't seem to make a computational difference, it could start at any point and still work.
	currentNum := 27995004
	currentRow := 6
	currentCol := 6
	maxCoor := 11

	//Set in the instructions
	multiplier := 252533
	divisor := 33554393

	for _, v := range bits {
		bit := strings.Trim(v, ",.")
		num, err := strconv.Atoi(bit)
		if err == nil {
			nums = append(nums, num)
		}
	}

	goalRow := nums[0]
	goalCol := nums[1]

	for currentRow != goalRow || currentCol != goalCol {
		if currentRow > 1 {
			currentRow--
			currentCol++
		} else if currentRow == 1 {
			currentCol = 1
			maxCoor++
			currentRow = maxCoor
		}
		currentNum = (currentNum * multiplier) % divisor
	}
	fmt.Printf("Day 25 Part %d - The number at Col: %d, Row: %d is %d\n", part, goalCol, goalRow, currentNum)
}
