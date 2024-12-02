package main

import (
	"bufio"
	"os"
	"strings"
)

func readLines() ([]string, error) {
	lines := make([]string, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

func abs(x int) int {
	if x > 0 {
		return int(x)
	}
	return int(-x)
}
