package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day9_1(input string) {

	inArr := strings.Split(input, "")
	sum := 0
	compArr := []string{}
	fileId := 0

	for i := 0; i < len(inArr); i++ {
		if i%2 == 0 {
			n, _ := strconv.Atoi(inArr[i])
			id := strconv.Itoa(fileId)
			for range n {
				compArr = append(compArr, id)
			}
			fileId++
		} else {
			n, _ := strconv.Atoi(inArr[i])
			for range n {
				compArr = append(compArr, ".")
			}
		}
	}

	last := len(compArr) - 1

	for i, v := range compArr {
		if v == "." {
			compArr[i], compArr[last] = compArr[last], compArr[i]
			for compArr[last] == "." {
				last--
			}
		}
		if last == i {
			break
		}
	}

	for i, v := range compArr {
		n, _ := strconv.Atoi(v)
		sum += i * n
	}

	println("Day9.1: ", sum)
}

func Day9_2(input string) {

	inArr := strings.Split(input, "")
	sum := 0
	compArr := [][]string{}
	fileId := 0

	for i := 0; i < len(inArr); i++ {
		if i%2 == 0 {
			n, _ := strconv.Atoi(inArr[i])
			id := strconv.Itoa(fileId)
			nArr := []string{}
			for range n {
				nArr = append(nArr, id)
			}
			compArr = append(compArr, nArr)
			fileId++
		} else {
			n, _ := strconv.Atoi(inArr[i])
			nArr := []string{}
			if n == 0 {
				continue
			}
			for range n {
				nArr = append(nArr, ".")
			}
			compArr = append(compArr, nArr)
		}
	}

	last := len(compArr) - 1
	for z := last; z >= 0; z-- {
		if compArr[z][0] != "." {
			for i := range compArr {
				if len(compArr[i]) >= len(compArr[z]) && compArr[i][0] == "." {

					if len(compArr[i]) > len(compArr[z]) {
						free1 := compArr[i][:len(compArr[z])]
						free2 := compArr[i][len(compArr[z]):]
						compArr = append(compArr[:i], compArr[i+1:]...)
						compArr = slices.Insert(compArr, i, free1)
						compArr = slices.Insert(compArr, i+1, free2)
						z++
					}

					compArr[i], compArr[z] = compArr[z], compArr[i]
					break
				}
				if i >= z {
					break
				}
			}
		}
	}

	finArr := []string{}

	for _, row := range compArr {
		for _, el := range row {
			finArr = append(finArr, el)
		}
	}

	for i, v := range finArr {
		if v != "." {
			n, _ := strconv.Atoi(v)
			sum += i * n
		}
	}

	println("Day9.2: ", sum)
}
