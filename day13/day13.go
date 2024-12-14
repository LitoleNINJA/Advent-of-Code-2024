package day13

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func DAY13() {
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

type Arcade struct {
	buttonA [2]int64
	buttonB [2]int64
	prize   [2]int64
}

func partOne(lines []string) int64 {
	ans := int64(0)
	arcadeMachines := readArcades(lines)

	for _, arcade := range arcadeMachines {
		if moves, ok := isReachable(arcade, false); ok {
			ans += moves
		}
	}

	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)
	arcadeMachines := readArcades(lines)

	for _, arcade := range arcadeMachines {
		if moves, ok := isReachable(arcade, true); ok {
			ans += moves
		}
	}

	return ans
}

func readArcades(lines []string) []Arcade {
	arcadeMachines := make([]Arcade, 0)

	var arcade Arcade
	for _, line := range lines {
		if len(line) == 0 {
			arcadeMachines = append(arcadeMachines, arcade)
			arcade = Arcade{}
		}

		if config, ok := strings.CutPrefix(line, "Button A: "); ok {
			x, y := utils.ReadXY(config, "X+", "Y+", ", ")
			arcade.buttonA = [2]int64{x, y}
		} else if config, ok := strings.CutPrefix(line, "Button B: "); ok {
			x, y := utils.ReadXY(config, "X+", "Y+", ", ")
			arcade.buttonB = [2]int64{x, y}
		} else if config, ok := strings.CutPrefix(line, "Prize: "); ok {
			x, y := utils.ReadXY(config, "X=", "Y=", ", ")
			arcade.prize = [2]int64{x, y}
		}
	}
	arcadeMachines = append(arcadeMachines, arcade)

	return arcadeMachines
}

func isReachable(arcade Arcade, isPart2 bool) (int64, bool) {
	moves := int64(0)
	reachable := false
	add := int64(0)
	if isPart2 {
		add = 10000000000000
	}

	x1, y1 := arcade.buttonA[0], arcade.buttonA[1]
	x2, y2 := arcade.buttonB[0], arcade.buttonB[1]
	c := arcade.prize[0] + add
	d := arcade.prize[1] + add

	denominator := x1*y2 - y1*x2
	if denominator == 0 {
		return moves, false
	}

	a := float64(c*y2-d*x2) / float64(denominator)
	b := float64(d*x1-c*y1) / float64(denominator)

	if float64(int64(a)) == a && float64(int64(b)) == b {
		// fmt.Println(a, b)
		moves = int64(a*3 + b)
		reachable = true
	}

	return moves, reachable
}
