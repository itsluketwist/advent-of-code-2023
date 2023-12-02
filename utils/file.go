package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToArray(day int, name string, test bool) ([]string, error) {
	var path string
	if test {
		path = fmt.Sprintf("%s.txt", name)
	} else {
		path = fmt.Sprintf("day%02d/%s.txt", day, name)
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
