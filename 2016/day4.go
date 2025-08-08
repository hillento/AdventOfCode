package sixteen

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// 361089 is too low

func Day4(input string, part int) {

	sum := 0
	rooms := map[string]int{}
	
	for line := range strings.Lines(input) {
		letterMap := map[string]int{}
		rule := line[len(line) - 7: len(line) -2]
		sectorSlice := []string{}
		seq := line[:len(line) -7]
		part2Seq := []string{}
		for _, v := range seq {
			if unicode.IsLetter(v) {
				_, ok := letterMap[string(v)]
				if ok {
					letterMap[string(v)]++
				} else {
					letterMap[string(v)] = 1
				}
			}
			if unicode.IsDigit(v) {
				sectorSlice = append(sectorSlice, string(v))
			}
				
		}
		sector, _ := strconv.Atoi(strings.Join(sectorSlice,""))
		for _, v := range seq {
			if v == '-' {
				part2Seq = append(part2Seq," ")
			} else if unicode.IsLetter(v) {
				part2Seq = append(part2Seq, string(rune((int(v) - 97 + sector) % 26) + 97 ))
			}
		}
		letters := make([]string,0,len(letterMap))
		for k := range letterMap {
			letters = append(letters, string(k))
		}

		sort.Slice(letters, func(i, j int) bool {
			if letterMap[letters[i]] != letterMap[letters[j]] {
				return letterMap[letters[i]] > letterMap[letters[j]]
			}
			return letters[i] < letters[j]
		})
		if strings.Join(letters[:5], "") == rule {
			num, _ := strconv.Atoi(strings.Join(sectorSlice,""))
			sum += num
			if part == 2 {
				rooms[strings.Join(part2Seq,"")] = sector
			}
		}
	}
	if part == 1{
		fmt.Printf("Day 4 Part %d - The sum of the sectors for real rooms is: %d\n",part,sum)
	} else {
		for k, v := range rooms {
			if strings.Contains(k, "north") {
				fmt.Printf("Day 4 Part %d - The sector ID of the room containing the north pole objects is : %d\n",part,v)
			}
		}
	}
}
