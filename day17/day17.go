package day17

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DAY17() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	// start = time.Now()
	// result = partTwo(lines)
	// duration = time.Since(start)
	// fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())
}

var reg = [3]int64{0, 0, 0}
var instPointer int
var output []int64

type Inst struct {
	opcode  int
	operand int
}

func partOne(lines []string) int64 {
	ans := int64(0)

	instSet := make([]Inst, 0)
	output = make([]int64, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "Register A") {
			reg[0] = getRegVal(line)
		} else if strings.HasPrefix(line, "Register B") {
			reg[1] = getRegVal(line)
		} else if strings.HasPrefix(line, "Register C") {
			reg[2] = getRegVal(line)
		} else if strings.HasPrefix(line, "Program") {
			val := getInstVal(line)
			for i := 0; i < len(val)-1; i += 2 {
				instSet = append(instSet, Inst{
					opcode:  int(val[i]),
					operand: int(val[i+1]),
				})
			}
		}
	}

	instPointer = 0
	for instPointer < len(instSet) {
		handleInst(instSet[instPointer])
		instPointer++
	}

	for i := range output {
		fmt.Print(output[i])
		if i < len(output) - 1 {
			fmt.Print(",")
		}
	}
	fmt.Println()

	return ans
}

func getRegVal(line string) int64 {
	val, _ := strconv.ParseInt(line[12:], 10, 64)
	return val
}

func getInstVal(line string) []int64 {
	values, _ := utils.StringSliceToInt(strings.Split(line[9:], ","))
	return values
}

func handleInst(inst Inst) {
	switch inst.opcode {
	case 0:
		num := reg[0]
		den := int64(1 << getComboOperand(inst.operand))
		reg[0] = num / den
	case 1:
		reg[1] = (reg[1] ^ int64(inst.operand))
	case 2:
		reg[1] = getComboOperand(inst.operand) % 8
	case 3:
		if reg[0] == 0 {
			return
		}

		instPointer = inst.operand - 1
	case 4:
		reg[1] = reg[1] ^ reg[2]
	case 5:
	 	output = append(output, getComboOperand(inst.operand) % 8)
	case 6:
		num := reg[0]
		den := int64(1 << getComboOperand(inst.operand))
		reg[1] = num / den
	case 7:
		num := reg[0]
		den := int64(1 << getComboOperand(inst.operand))
		reg[2] = num / den
	default:
		fmt.Printf("Unknown opcode : %d\n", inst.opcode)
	}
}

func getComboOperand(operand int) int64 {
	if operand <= 3 {
		return int64(operand)
	} else if operand <= 6 {
		return reg[operand-4]
	} else {
		fmt.Println("invalid operand 7")
		return -1
	}
}
