package day22

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

func DAY22() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start)
	fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())
}

func partOne(lines []string) int64 {
	ans := int64(0)

	for _, line := range lines {
		val, _ := strconv.ParseInt(line, 10, 64)
		ans += findSecret(val, 2000)
	}

	return ans
}

type Seq struct {
	seq [4]int
}

func partTwo(lines []string) int64 {
	ans := int64(0)

	freq := make(map[Seq]int, 0)
	for _, line := range lines {
		val, _ := strconv.ParseInt(line, 10, 64)
		diffArr, values := getDiffArray(val, 2000)
		localFreq := make(map[Seq]int, 0)
		for i:=0; i<len(diffArr)-4; i++ {
			var cur [4]int
			copy(cur[:], diffArr[i:i+4])
			if _, ok := localFreq[Seq{seq: cur}]; !ok {
				localFreq[Seq{seq: cur}] = values[i+3]
			}
		}
		// fmt.Println(localFreq[Seq{seq: [4]int{-2, 1, -1, 3}}])
		for key, val := range localFreq {
			freq[key] += val
		}
	}


	for _, val := range freq {
		if int64(val) > ans {
			ans = int64(val)
			// fmt.Printf("Key : %+v, Val : %+v\n", key, val)
		}
	}

	return ans
}


func findSecret(val int64, k int) int64 {
	ans := val
	for i:=0; i<k; i++ {
		ans = mixPrune(ans, ans * 64)
		ans = mixPrune(ans, ans / 32)
		ans = mixPrune(ans, ans * 2048)

	}

	return ans
}

func mixPrune(ans int64, val int64) int64 {
	ans = ans ^ val
	ans %= 16777216
	return ans
}

func getDiffArray(val int64, k int) ([]int, []int) {
	prev := val
	ans := make([]int, 0)
	banana := make([]int, 0)
	for i:=0; i<k; i++ {
		val = findSecret(val, 1)
		ans = append(ans, int(val%10) - int(prev%10))
		banana = append(banana, int(val%10))
		prev = val
	}
	return ans, banana
}