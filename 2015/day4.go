package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

func Day4(input string, part int) {
	prefixZeroes := 5 + (part - 1)
	for i := 0; i < math.MaxInt32; i++ {
		combined := fmt.Sprintf("%s%d", strings.TrimSpace(input), i)
		hashed := fmt.Sprintf("%x", md5.Sum([]byte(combined)))
		if strings.HasPrefix(hashed, strings.Repeat("0", prefixZeroes)) {
			fmt.Printf("Day4 - Input: %s, Hex: %s, prefix: %d\n", input, hashed, i)
			break
		}
	}
}
