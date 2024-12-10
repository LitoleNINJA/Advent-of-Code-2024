package day9

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func DAY9() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	// fmt.Printf("Part 1 : %d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start)
	fmt.Printf("Part 2 : %d | Time Taken : %.4f seconds\n", result, duration.Seconds())
}

func partOne(lines []string) int64 {
	ans := int64(0)

	res := diskMap(lines[0])
	// fmt.Println(res)
	compressedRes := compressDiskMap(res)
	// fmt.Println(compressedRes)

	ans = checkSum(compressedRes)
	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)

	res := diskMap(lines[0])
	// fmt.Println(res)
	compressedRes := compressDiskMap2(res)
	fmt.Println(compressedRes)

	// ans = checkSum(compressedRes)
	return ans
}

func diskMap(line string) []string {
	id := 0
	flag := true
	res := make([]string, 0)

	for _, ch := range line {
		if flag {
			val := int(ch - '0')
			for i := 0; i < val; i++ {
				res = append(res, strconv.Itoa(id))
			}
			id++
		} else {
			val := int(ch - '0')
			for i := 0; i < val; i++ {
				res = append(res, ".")
			}
		}
		flag = !flag
	}

	return res
}

func compressDiskMap(diskMap []string) []string {
	ans := diskMap

	l, r := -1, -1
	for i, s := range diskMap {
		if isNumber(s) {
			r = i
		} else if l == -1 && s == "." {
			l = i
		}
	}

	for {
		if l >= r {
			break
		}

		ans[l] = ans[r]
		ans[r] = "."

		for {
			if ans[l] == "." {
				break
			}
			l++
		}
		for {
			if isNumber(ans[r]) {
				break
			}
			r--
		}
	}

	return ans
}

func compressDiskMap2(diskMap []string) [][]string {
	ans := make([][]string, 0)
	
	for i := 0; i < len(diskMap); i++ {
		cur := make([]string, 0)
		cur = append(cur, diskMap[i])
		for {
			if i < len(diskMap)-1 && diskMap[i+1] != diskMap[i] {
				ans = append(ans, cur)
				continue
			}
			cur = append(cur, diskMap[i])
		}
	}

	return ans
}

func checkSum(diskMap []string) int64 {
	ans := int64(0)

	for i, s := range diskMap {
		if s == "." {
			continue
		}

		val, _ := strconv.Atoi(s)
		ans += int64(i * val)
	}

	return ans
}

func isNumber(s string) bool {
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
