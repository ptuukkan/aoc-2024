package day17

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func parseInput(input string) ([]int, []int) {
	split := strings.Split(input, "\n\n")
	registerInputs := utils.SplitNewLines(split[0])

	registers := make([]int, 3)
	var instructions []int

	for i, r := range registerInputs {
		spl := strings.Split(r, " ")
		num, err := strconv.Atoi(spl[2])
		if err != nil {
			fmt.Println("error converting numbers")
		}
		registers[i] = num
	}

	a := strings.Split(strings.Trim(split[1], "\n"), " ")
	for _, instr := range strings.Split(a[1], ",") {
		num, err := strconv.Atoi(instr)
		if err != nil {
			fmt.Println("error converting numbers")
		}
		instructions = append(instructions, num)
	}

	return registers, instructions
}

func readComboValue(registers []int, combo int) int {
	switch combo {
	case 0, 1, 2, 3:
		return combo
	default:
		return registers[combo-4]
	}
}

func divHelper(registers []int, numerator int, operand int) int {
	denominator := readComboValue(registers, operand)
	denominator = int(math.Pow(2, float64(denominator)))

	return numerator / denominator
}

func adv(registers []int, operand int) {
	numerator := registers[0]
	registers[0] = divHelper(registers, numerator, operand)
}

func bxl(registers []int, operand int) {
	result := registers[1] ^ operand
	registers[1] = result
}

func bst(registers []int, operand int) {
	op := readComboValue(registers, operand)
	result := op % 8
	registers[1] = result
}

func jnz(registers []int) bool {
	if registers[0] == 0 {
		return false
	}
	return true
}

func bxc(registers []int, operand int) {
	result := registers[1] ^ registers[2]
	registers[1] = result
}

func out(registers []int, operand int) int {
	op := readComboValue(registers, operand)

	result := op % 8
	return result
}

func bdv(registers []int, operand int) {
	numerator := registers[0]
	registers[1] = divHelper(registers, numerator, operand)
}
func cdv(registers []int, operand int) {
	numerator := registers[0]
	registers[2] = divHelper(registers, numerator, operand)
}

func runProgram(registers []int, instructions []int) string {
	pointer := 0
	var output []string

	for pointer < len(instructions) {
		switch instructions[pointer] {
		case 0:
			adv(registers, instructions[pointer+1])
		case 1:
			bxl(registers, instructions[pointer+1])
		case 2:
			bst(registers, instructions[pointer+1])
		case 3:
			if jnz(registers) {
				pointer = instructions[pointer+1] - 2
			}
		case 4:
			bxc(registers, instructions[pointer+1])
		case 5:
			num := out(registers, instructions[pointer+1])
			output = append(output, strconv.Itoa(num))
		case 6:
			bdv(registers, instructions[pointer+1])
		case 7:
			cdv(registers, instructions[pointer+1])
		}
		pointer += 2
	}

	return strings.Join(output, ",")
}

func Part1(input string) string {
	registers, instructions := parseInput(input)

	result := runProgram(registers, instructions)
	return result
}

func findResult(instructions, reversed []int, result int) (bool, int) {
	if len(reversed) == 0 {
		return true, result
	}
	result <<= 3
	goal := reversed[0]

	for i := 0; i < 8; i++ {
		registers := []int{result | i, 0, 0}
		output := runProgram(registers, instructions)
		if strings.HasPrefix(output, strconv.Itoa(goal)) {
			if success, res := findResult(instructions, reversed[1:], result|i); success {
				return success, res
			}
		}
	}
	return false, result
}

func Part2(input string) string {
	_, instructions := parseInput(input)
	reversed := make([]int, len(instructions))
	copy(reversed, instructions)
	slices.Reverse(reversed)

	success, result := findResult(instructions, reversed, 0)
	if success {
		return fmt.Sprint(result)
	}

	return "error"
}
