package day14

import (
	"aoc2024/utils"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DAY14() {
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

type Robot struct {
	pos [2]int
	vel [2]int
}

func partOne(lines []string) int64 {
	ans := int64(1)

	robots := make([]Robot, 0)
	for _, line := range lines {
		var robot Robot
		values := strings.Fields(line)
		x, y := utils.ReadXY(values[0], "p=", "", ",")
		robot.pos = [2]int{int(x), int(y)}
		x, y = utils.ReadXY(values[1], "v=", "", ",")
		robot.vel = [2]int{int(x), int(y)}

		robots = append(robots, robot)
	}

	quad := quadCount(robots, 101, 103)
	for _, val := range quad {
		ans *= int64(val)
	}
	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(1)

	robots := make([]Robot, 0)
	for _, line := range lines {
		var robot Robot
		values := strings.Fields(line)
		x, y := utils.ReadXY(values[0], "p=", "", ",")
		robot.pos = [2]int{int(x), int(y)}
		x, y = utils.ReadXY(values[1], "v=", "", ",")
		robot.vel = [2]int{int(x), int(y)}

		robots = append(robots, robot)
	}

	// for i := 101; i <= 10000; i++ {
		var grid [101][103]byte
		for i := range grid {
			for j := range grid[i] {
				grid[i][j] = '.'
			}
		}
		for _, robot := range robots {
			x, y := findPos(robot, 101, 103, 7709)
			grid[x][y] = '#'
		}

		if err := saveGridAsPNG(grid, 7709); err != nil {
			fmt.Printf("Error saving frame %d: %v\n", 7709, err)
		}

	// }
	return ans
}

func quadCount(robots []Robot, n, m int) [4]int {
	var ans [4]int
	for _, robot := range robots {
		x, y := findPos(robot, n, m, 100)

		if x < n/2 && y < m/2 {
			ans[0]++
		} else if x < n/2 && y > m/2 {
			ans[1]++
		} else if x > n/2 && y < m/2 {
			ans[2]++
		} else if x > n/2 && y > m/2 {
			ans[3]++
		}
	}

	return ans
}

func findPos(robot Robot, n, m, time int) (int, int) {
	x, y := robot.pos[0], robot.pos[1]
	vx, vy := robot.vel[0], robot.vel[1]

	for i := 0; i < time; i++ {
		x = utils.Mod(x+vx, n)
		y = utils.Mod(y+vy, m)
	}

	return x, y
}

func saveGridAsPNG(grid [101][103]byte, frameNum int) error {
	outDir := "frames"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	scale := 10
	img := image.NewRGBA(image.Rect(0, 0, 103*scale, 101*scale))

	for i := range grid {
		for j := range grid[i] {
			c := color.White
			if grid[i][j] == '#' {
				c = color.Black
			}
			for si := 0; si < scale; si++ {
				for sj := 0; sj < scale; sj++ {
					img.Set(j*scale+sj, i*scale+si, c)
				}
			}
		}
	}

	filename := filepath.Join(outDir, fmt.Sprintf("frame_%05d.png", frameNum))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}