package fifteen

import (
	"fmt"
)

type Coor struct {
	x int
	y int
}

func Day3(input string) {
	coords := Coor{0, 0}
	count := 1
	m := make(map[Coor]int)
	m[coords] = count

	for _, r := range input {
		switch r {
		case '^':
			coords.y++
		case '>':
			coords.x++
		case 'v':
			coords.y--
		case '<':
			coords.x--
		}
		if m[coords] != 0 {
			continue
		}
		count++
		m[coords] = count
	}
	fmt.Printf("Day3 - Santa visited %d unique houses\n", count)
}

func Day3_2(input string) {

	rCoords := Coor{0, 0}
	sCoords := Coor{0, 0}
	coords := Coor{0, 0}
	count := 1
	m := make(map[Coor]int)
	m[sCoords] = count

	for i, r := range input {
		if i%2 == 0 {
			rCoords = travel(rCoords, r)
			coords = rCoords
		} else {
			sCoords = travel(sCoords, r)
			coords = sCoords
		}
		if m[coords] != 0 {
			continue
		}
		count++
		m[coords] = count
	}
	fmt.Printf("Day3 - Santa and robot visited %d unique houses\n", count)
}

func travel(coords Coor, dir rune) Coor {
	switch dir {
	case '^':
		coords.y++
	case '>':
		coords.x++
	case 'v':
		coords.y--
	case '<':
		coords.x--
	}
	return coords
}
