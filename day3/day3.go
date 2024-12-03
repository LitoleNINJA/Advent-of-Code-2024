package day3

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func DAY3() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	var ans int = 0
	for _, line := range lines {
		ans += findMuls(line)
	}
}

func findMuls(line string) int {
	ans := 0
	pos := 0
	indArr := make([]int, 0)
	for {
		ind := strings.Index(line[pos:], "mul(")
		if ind == -1 {
			break
		}

		indArr = append(indArr, ind+pos+3)
		pos += ind + 4
	}

	for i, ind := range indArr {
		nextInd := strings.Index(line[ind:], ")")
		if nextInd == -1 {
			break
		}
		
		if i < len(indArr)-1 && ind+nextInd > indArr[i+1] {
			continue
		}
		
		ans 
	}
	// fmt.Println(indArr)
	return ans
}

func calcMul(line string) int {

	fmt.Println(line)
	return 0
}
