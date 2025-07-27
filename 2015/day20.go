package fifteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Not a fan of brute forcing this but can't really think of an elegant way of doing it right now
func Day20(input string, part int) {
	goal, _ := strconv.Atoi(strings.TrimSpace(input))
	var house int

	for i := 1; i < math.MaxInt64; i++ {
		presents := 0
		for _, f := range getFactors(i) {
			if part == 1 {
				presents += f * 10
			} else if part == 2 && i/f < 50 {
				presents += f * 11
			}
		}
		if presents >= goal {
			house = i
			break
		}
	}
	fmt.Printf("Day 20 part %d - House Number: %d\n", part, int(house))

}

func getFactors(house int) []int {
	factors := []int{}
	sqrt := int(math.Sqrt(float64(house)))
	for i := 1; i <= sqrt; i++ {
		if house%i == 0 {
			factors = append(factors, i, house/i)
		}
	}
	return factors
}
