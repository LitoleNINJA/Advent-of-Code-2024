package day1

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftList []string
var rightList []string

func day1() {
	leftList = make([]string, 0)
	rightList = make([]string, 0)

	err := readFile()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	if len(leftList) != len(rightList) {
		fmt.Println("List size mismatch !")
		os.Exit(1)
	}

	fmt.Printf("Part 1 : %d\n", partOne())
	fmt.Printf("Part 2 : %d\n", partTwo())
}

func partOne() int64 {
	sort.Strings(leftList)
	sort.Strings(rightList)

	ans := int64(0)
	for i := 0; i < len(leftList); i++ {
		leftInt, err := strconv.ParseInt(leftList[i], 10, 64)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}

		rightInt, err := strconv.ParseInt(rightList[i], 10, 64)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}

		ans += int64(utils.Abs(int(leftInt - rightInt)))
	}

	return ans
}

func partTwo() int64 {
	sort.Strings(leftList)

	fre := make(map[string]int64)
	for _, i := range rightList {
		fre[i]++
	}

	ans := int64(0)
	for i := 0; i < len(leftList); i++ {
		leftInt, err := strconv.ParseInt(leftList[i], 10, 64)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}

		if val, ok := fre[leftList[i]]; ok {
			ans += leftInt * val
		}
	}

	return ans
}
func readFile() error {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}
		leftList = append(leftList, parts[0])
		rightList = append(rightList, parts[1])
	}
	return nil
}
