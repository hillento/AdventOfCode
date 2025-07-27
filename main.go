package main

import (
	fifteen "AdventOfCode/fifteen"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	// testInput := "H => HO\nH => OH\nO => HH\n\nHOHOHO"
	// Day19(testInput)

	year := 2015
	day := 22
	i := getInput(year, day)
	fifteen.Day21(i, 2)
}
