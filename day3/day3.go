package day3

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func DAY3() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : %d\n", partOne(lines))
	fmt.Printf("Part 2 : %d\n", partTwo(lines))
}

func partOne(lines []string) int {
	var ans int = 0
	for _, line := range lines {
		ans += findMuls(line, false)
	}

	return ans
}

func findMuls(line string, enableCheck bool) int {
	ans := 0
	indArr := make([]int, 0)
	isEnabled := true
	for i := 0; i < len(line); i++ {
		if isEnabled && i < len(line)-4 && line[i:i+4] == "mul(" {
			indArr = append(indArr, i+3)
			i += 3
		}

		if enableCheck && i < len(line)-4 && line[i:i+4] == "do()" {
			isEnabled = true
			i += 3
		} else if enableCheck && i < len(line)-7 && line[i:i+7] == "don't()" {
			isEnabled = false
			i += 6
		}
	}

	for i, ind := range indArr {
		nextInd := strings.Index(line[ind:], ")")
		if nextInd == -1 {
			break
		}

		if i < len(indArr)-1 && ind+nextInd > indArr[i+1] {
			continue
		}

		// fmt.Println(line[ind+1 : ind+nextInd])
		ans += calcMul(line[ind+1 : ind+nextInd])
	}

	return ans
}

func calcMul(line string) int {

	reg, err := regexp.Compile(`(\d{1,3}),(\d{1,3})`)
	if err != nil {
		fmt.Printf("Error matching regex : %v", err)
		return 0
	}

	if reg.MatchString(line) {
		x, y, ok := strings.Cut(line, ",")
		if !ok {
			fmt.Printf(", not found in %s\n", line)
			return 0
		}

		x_int, _ := strconv.Atoi(x)
		y_int, _ := strconv.Atoi(y)

		return x_int * y_int
	}
	return 0
}

func partTwo(lines []string) int {
	var ans int = 0
	for _, line := range lines {
		ans += findMuls(line, true)
	}

	return ans
}
