package fifteen

import (
	"fmt"
)

func Day1(input string) {
	floor := 0
	basementEntered := false
	for i, r := range input {

		if r == '(' {
			floor += 1
		}
		if r == ')' {
			floor -= 1
			if basementEntered == false && floor == -1 {
				basementEntered = true
				fmt.Printf("Day1_2 Solution: %d\n", i+1)
			}
		}
	}
	fmt.Printf("Day1_1 Solution: %d\n", floor)
}
