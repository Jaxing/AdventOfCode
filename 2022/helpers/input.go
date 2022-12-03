package helpers

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

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

func Submit(year int, day int, task int, answer int) (string, error) {
	cookies, _ := ReadFile("/home/jesper/code/AdventOfCode/cookies")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("level", strconv.Itoa(task))
	_ = writer.WriteField("answer", strconv.Itoa(answer))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookies[0]))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err_req := client.Do(req)

	if err_req != nil {
		return "", err_req
	}
	defer resp.Body.Close()
	body, err_resp := ioutil.ReadAll(resp.Body)
	return ParseSubmitMessage(string(body)), err_resp
}

func ParseSubmitMessage(message string) string {
	r, _ := regexp.Compile("<main>\n\t<article><p>(.*)</p></article>\n\t</main>")
	match := r.FindStringSubmatch(message)

	return match[1]
}
