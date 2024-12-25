package day24

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func DAY24() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	/*start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start)
	fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())*/
}

var input map[string]int
var adj map[string][]string
var logic map[string]int // 0 => AND | 1 => OR | 2 => XOR

func partOne(lines []string) int64 {
	ans := int64(0)

	isInputWire := true
	input = make(map[string]int, 0)
	logic = make(map[string]int, 0)
	adj = make(map[string][]string, 0)
	zValues := make([]string, 0)
	for _, line := range  lines {
		if len(line) == 0 {
			isInputWire = false
			continue
		}

		if isInputWire {
			s, val := getInput(line)
			input[s] = val
		} else {
			in1, in2, op, out := getGateValues(line)
			adj[out] = append(adj[out], in1, in2)
			logic[out] = op

			if strings.HasPrefix(out, "z") {
				zValues = append(zValues, out)
			}
		}
	}

	slices.Sort(zValues)
	for i, z := range zValues {
		val := resolve(z)
		// fmt.Printf("%s = %d\n", z, val)
		if val == 1 {
			ans |= (1<<i)
		}
	}

	return ans
}

func getInput(line string) (string, int) {
	s, valStr, _ := strings.Cut(line, ": ")
	val, _ := strconv.Atoi(valStr)

	return s, val
}


func getGateValues(line string) (string, string, int, string) {
	left, right, _ := strings.Cut(line, " -> ")
	values := strings.Fields(left)
	var op int
	if values[1] == "AND" {
		op = 0
	} else if values[1] == "OR" {
		op = 1
	} else if values[1] == "XOR" {
		op = 2
	} else {
		fmt.Println("Invalid op")
	}

	return values[0], values[2], op, right
}

func resolve(u string) int {
	if val, ok := input[u]; ok {
		return val
	}

	if len(adj[u]) != 2 {
		fmt.Println("What happend ?")
		os.Exit(1)
	}

	x := resolve(adj[u][0])
	y := resolve(adj[u][1])
	
	var res int
	if logic[u] == 0 {
		res = x & y
	} else if logic[u] == 1 {
		res = x | y
	} else if logic[u] == 2 {
		res = x ^ y
	} else {
		fmt.Println("What the hell happend ?")
		os.Exit(69)
	}

	input[u] = res
	return res
}