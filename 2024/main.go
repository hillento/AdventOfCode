package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getInput(day int) string {

	token := os.Getenv("AoCSession")
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	fileName := fmt.Sprintf("Day%d.txt", day)
	filePath := fmt.Sprintf("./inputs/%s", fileName)

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
	day := 11
	s := getInput(day)
	Day11(s)
}
