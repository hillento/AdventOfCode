package sixteen

import (
	"fmt"
	"strconv"
	"strings"
)

func Day3(input string, part int) {
	possible := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")

	if part == 1 {
		for _, v := range lines {
			nums := strings.Fields(v)
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			n3, _ := strconv.Atoi(nums[2])

			possible += validTri(n1, n2, n3)
		}
	}
	if part == 2 {
		for i := 0; i < len(lines)-2; i += 3 {
			nums1 := strings.Fields(lines[i])
			nums2 := strings.Fields(lines[i+1])
			nums3 := strings.Fields(lines[i+2])

			n1t1, _ := strconv.Atoi(nums1[0])
			n2t1, _ := strconv.Atoi(nums2[0])
			n3t1, _ := strconv.Atoi(nums3[0])
			possible += validTri(n1t1, n2t1, n3t1)

			n1t2, _ := strconv.Atoi(nums1[1])
			n2t2, _ := strconv.Atoi(nums2[1])
			n3t2, _ := strconv.Atoi(nums3[1])
			possible += validTri(n1t2, n2t2, n3t2)

			n1t3, _ := strconv.Atoi(nums1[2])
			n2t3, _ := strconv.Atoi(nums2[2])
			n3t3, _ := strconv.Atoi(nums3[2])
			possible += validTri(n1t3, n2t3, n3t3)
		}
	}
	fmt.Printf("Day 3 Part %d - Number of valid triangles: %d\n", part, possible)
}

func validTri(n1, n2, n3 int) int {
	if n1+n2 > n3 && n2+n3 > n1 && n1+n3 > n2 {
		return 1
	}
	return 0
}
