package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var Dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func ReadLines() ([]string, error) {
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

func Abs(x int) int {
	if x > 0 {
		return int(x)
	}
	return int(-x)
}

func StringSliceToInt(values []string) ([]int64, error) {
	result := make([]int64, 0, len(values))

	for _, v := range values {
		// Skip empty strings
		if v == "" {
			continue
		}

		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

func InRange(lines []string, i int, j int) bool {
	return (i >= 0 && i < len(lines) && j >= 0 && j < len(lines[0]))
}

func DigitCount(n int64) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func StringSliceToByte(values []string) [][]byte {
	result := make([][]byte, 0, len(values))

	for _, v := range values {
		result = append(result, []byte(v))
	}

	return result
}
