package fifteen

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	paper := 0
	ribbon := 0

	for _, r := range lines {
		nums := strings.Split(r, "x")
		var dmns []int
		var faces []int
		l, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])
		h, _ := strconv.Atoi(nums[2])

		dmns = append(dmns, l)
		dmns = append(dmns, w)
		dmns = append(dmns, h)
		sort.Ints(dmns)

		faces = append(faces, l*w)
		faces = append(faces, l*h)
		faces = append(faces, h*w)
		sort.Ints(faces)

		paper += (faces[0] * 3) + (faces[1] * 2) + (faces[2] * 2)
		ribbon += (dmns[0] + dmns[0] + dmns[1] + dmns[1]) + (dmns[0] * dmns[1] * dmns[2])
	}
	fmt.Printf("2015 Day 2 Paper Needed: %d\n", paper)
	fmt.Printf("2015 Day 2 Ribbon Needed: %d\n", ribbon)
}
