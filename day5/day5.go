package day5

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func DAY5() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : %d\n", partOne(lines))
	fmt.Printf("Part 2 : %d\n", partTwo(lines))
}

var adj map[string][]string

func partOne(lines []string) int {
	ans := 0
	isEdge := true
	adj = make(map[string][]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			isEdge = false
			continue
		}

		if isEdge {
			x, y, _ := strings.Cut(line, "|")
			addEdge(x, y)
		} else {
			page := strings.Split(line, ",")
			if ok, val := isPageGood(page); ok {
                ans += val
            }
		}
	}

	return ans
}

func addEdge(x string, y string) {
	adj[y] = append(adj[y], x)
}

func isPageGood(page []string) (bool, int) {
    isGood := true
	for i, u := range page {
		if neigh, ok := adj[u]; ok {
			for _, v := range neigh {
				// v should not come after u
				if ind := slices.Index(page[i+1:], v); ind != -1 {
                    page[i], page[i+1+ind] = page[i+1+ind], page[i]
                    isGood = false
				}
			}
		}
	}

    val, _ := strconv.Atoi(page[len(page)/2])
	return isGood, val
}

func partTwo(lines []string) int {
	ans := 0
	isEdge := true
	adj = make(map[string][]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			isEdge = false
			continue
		}

		if isEdge {
			x, y, _ := strings.Cut(line, "|")
			addEdge(x, y)
		} else {
			page := strings.Split(line, ",")
			if ok, val := isPageGood(page); !ok {
                ans += val
            }
		}
	}

	return ans
}
