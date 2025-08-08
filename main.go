package main

import (
	// fifteen "AdventOfCode/fifteen"
	sixteen "AdventOfCode/sixteen"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Part 2 131 is too high

func getInput(year, day int) string {

	token := os.Getenv("AoCSession")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	fileName := fmt.Sprintf("Day%d.txt", day)
	filePath := fmt.Sprintf("./%d/inputs/%s", year, fileName)

	// if the file already exists just return the file content
	input, err := os.ReadFile(filePath)
	if err == nil {
		return string(input)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Add("cookie", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	_ = os.WriteFile(filePath, body, 0664)

	return string(body)
}

func main() {

	// testInput := "aaaaa-bbb-z-y-x-123[abxyz]\na-b-c-d-e-f-g-h-987[abcde]\nnot-a-real-room-404[oarel]\ntotally-real-room-200[decoy]"

	// i := testInput
	year := 2016
	day := 4
	i := getInput(year, day)
	sixteen.Day4(i, 1)
	sixteen.Day4(i, 2)
}
