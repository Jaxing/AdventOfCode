package helpers

import (
	"os"
	"bufio"
)

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
