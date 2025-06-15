package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Lights struct {
	grid map[Coordinate]int
}

func newLights() *Lights {
	var l Lights
	l.grid = make(map[Coordinate]int)
	l.turnLights(Coordinate{x: 0, y: 0}, Coordinate{999, 999}, "off")
	return &l
}

func Day6(input string, part int) {
	lights := newLights()
	on := 0
	lines := strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(input, "through ", ""), "turn ", "")), "\n")

	for _, l := range lines {
		pieces := strings.Split(l, " ")
		command := pieces[0]
		start := strings.Split(pieces[1], ",")
		end := strings.Split(pieces[2], ",")

		sx, _ := strconv.Atoi(start[0])
		sy, _ := strconv.Atoi(start[1])
		ex, _ := strconv.Atoi(end[0])
		ey, _ := strconv.Atoi(end[1])

		if part == 1 {
			lights.turnLights(Coordinate{sx, sy}, Coordinate{ex, ey}, command)
		} else {
			lights.lightsBrightness(Coordinate{sx, sy}, Coordinate{ex, ey}, command)
		}

	}

	for _, v := range lights.grid {
		on += v
	}

	fmt.Printf("Day 6 - Lights on: %d\n", on)
}

func (l Lights) lightsBrightness(start Coordinate, end Coordinate, command string) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			curr := Coordinate{x, y}
			switch command {
			case "on":
				l.grid[curr] += 1
			case "off":
				if l.grid[curr] > 0 {
					l.grid[curr] -= 1
				}
			case "toggle":
				l.grid[curr] += 2
			}
		}
	}
}

func (l Lights) turnLights(start Coordinate, end Coordinate, command string) {

	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			curr := Coordinate{x, y}
			switch command {
			case "on":
				l.grid[curr] = 1
			case "off":
				l.grid[curr] = 0
			case "toggle":
				if l.grid[curr] == 0 {
					l.grid[curr] = 1
				} else {
					l.grid[curr] = 0
				}
			}
		}
	}
}
