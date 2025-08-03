package fifteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Group struct {
	count int
	qe    int
}

func Day24(input string, part int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	weights := []int{}
	totalWeight := 0
	var goal int

	for _, v := range lines {
		weight, _ := strconv.Atoi(v)
		totalWeight += weight
		weights = append(weights, weight)
	}

	if part == 1 {
		goal = totalWeight / 3
	} else if part == 2 {
		goal = totalWeight / 4
	}

	minGroup := split(weights, 0, goal, 1, 0)
	fmt.Printf("Day 24 part %d - The packages in the center will have a count of %d and a QE of %d\n", part, minGroup.count, minGroup.qe)
}

func split(weights []int, position int, remainingWeight int, currentQE int, currentCount int) Group {
	if remainingWeight == 0 {
		return Group{currentCount, currentQE}
	} else if remainingWeight < 0 || position == len(weights) {
		return Group{math.MaxInt, math.MaxInt}
	}

	included := split(weights, position+1, remainingWeight-weights[position], currentQE*weights[position], currentCount+1)

	notIncluded := split(weights, position+1, remainingWeight, currentQE, currentCount)

	if included.count < notIncluded.count {
		return included
	} else if notIncluded.count < included.count {
		return notIncluded
	} else if included.qe < notIncluded.qe {
		return included
	} else {
		return notIncluded
	}
}
