package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

type Reindeer struct {
	name     string
	velocity int
	burst    int
	rest     int
	points   int
}

func Day14(input string, part int) {
	reindeerPerSec := map[string][]int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, l := range lines {
		distance := 0
		data := strings.Split(l, " ")
		velocity, _ := strconv.Atoi(string(data[3]))
		name := data[0]
		run, _ := strconv.Atoi(string(data[6]))
		rest, _ := strconv.Atoi(string(data[13]))

		remainingRun := run
		remainingRest := rest
		for t := 1; t <= 2503; t++ {
			if remainingRun > 0 {
				distance += velocity
				remainingRun--
			} else {
				remainingRest--
				if remainingRest == 0 {
					remainingRun = run
					remainingRest = rest
				}
			}
			reindeerPerSec[name] = append(reindeerPerSec[name], distance)
		}
	}
	if part == 1 {
		furthest := 0
		var fastest string
		for k, v := range reindeerPerSec {

			if v[2502] > furthest { //0 base index so 2502 is the numer of seconds in the race
				furthest = v[2502]
				fastest = k
			}
		}
		fmt.Printf("Day14 Part 1 - Fastest: %s with Distance: %d km\n", fastest, furthest)
	}

	if part == 2 {
		reindeerScores := map[string]int{}
		for t := 0; t < 2503; t++ {
			var leaders []string
			var leadDist int
			for name, distSlice := range reindeerPerSec {
				if distSlice[t] > leadDist {
					leaders = []string{name}
					leadDist = distSlice[t]
				} else if distSlice[t] == leadDist {
					leaders = append(leaders, name)
				}
			}

			for _, name := range leaders {
				reindeerScores[name]++
			}
		}
		var bestDeer string
		highScore := 0
		for n, v := range reindeerScores {
			if v > highScore {
				highScore = v
				bestDeer = n
			}
		}
		fmt.Printf("Day14 Part 2 - Highest Scoring Deer: %s with score: %d points\n", bestDeer, highScore)
	}
}
