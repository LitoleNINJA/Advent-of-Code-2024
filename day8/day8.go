package day8

import (
	"fmt"
	"os"
	"strings"
)

func DAY8() {
	// Read the input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Convert content to grid
	rawGrid := string(content)
	lines := strings.Split(rawGrid, "\n")

	// Create 2D grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Find unique frequencies (characters)
	frequencies := make(map[rune]bool)
	for _, line := range lines {
		for _, char := range line {
			if char != '\n' && char != '.' {
				frequencies[char] = true
			}
		}
	}

	// Find coordinates for each frequency
	freqCoords := make(map[rune][][]int)
	for freq := range frequencies {
		var coords [][]int
		for i, line := range grid {
			for j, char := range line {
				if char == freq {
					coords = append(coords, []int{i, j})
				}
			}
		}
		freqCoords[freq] = coords
	}

	// Get antinodes function
	getAntinodes := func(coords [][]int) map[string]bool {
		antinodes := make(map[string]bool)
		for _, coord1 := range coords {
			for _, coord2 := range coords {
				x1, y1 := coord1[0], coord1[1]
				x2, y2 := coord2[0], coord2[1]

				antiX := 2*x1 - x2
				antiY := 2*y1 - y2

				antinodes[fmt.Sprintf("%d,%d", antiX, antiY)] = true
			}
		}
		return antinodes
	}

	// Collect all valid antinodes
	validAntinodes := make(map[string]bool)
	for _, freqCoordList := range freqCoords {
		antinodes := getAntinodes(freqCoordList)

		for antinode := range antinodes {
			parts := strings.Split(antinode, ",")
			x, y := parseInt(parts[0]), parseInt(parts[1])

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
				validAntinodes[antinode] = true
			}
		}
	}

	// Print the answer
	fmt.Println(len(validAntinodes))
}

// Helper function to parse int safely
func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
