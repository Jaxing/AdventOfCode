package helpers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func DownloadInput(year int, day int) (string, error) {
	cookies, err := ReadFile("/home/jesper/code/AdventOfCode/cookies")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookies[0]))
	resp, _ := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func StringListToIntList(input_string []string) ([]int, error) {
	var input_length = len(input_string)
	var input = make([]int, input_length)

	for i := 0; i < input_length; i++ {
		new_int, err := strconv.Atoi(input_string[i])
		if err != nil {
			return nil, err
		}

		input[i] = new_int
	}

	return input, nil
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	var rows = []string{}

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, nil
}

func ReadLine(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), nil
}
