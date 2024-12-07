package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func parseInput(input string) []Equation {
	lines := utils.SplitNewLines(input)
	equations := make([]Equation, len(lines))

	for i, line := range lines {
		split := strings.Split(line, ": ")
		value, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("error converting number")
		}
		fields := strings.Fields(split[1])
		numbers := make([]int, len(fields))

		for k, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println("error converting number")
			}
			numbers[k] = num
		}
		equations[i].Numbers = numbers
		equations[i].TestValue = value
	}
	return equations
}

func evaluate(currentValue int, numbers []int, op string, targetValue int, operators []string) bool {
	var newValue int
	if op == "*" {
		newValue = currentValue * numbers[0]
	} else if op == "+" {
		newValue = currentValue + numbers[0]
	} else {
		str_a := strconv.Itoa(currentValue)
		str_b := strconv.Itoa(numbers[0])
		num, err := strconv.Atoi(str_a + str_b)
		if err != nil {
			fmt.Println("error converting number")
		}
		newValue = num
	}

	if len(numbers) == 1 {
		return newValue == targetValue
	}

	for _, o := range operators {
		if evaluate(newValue, numbers[1:], o, targetValue, operators) {
			return true
		}
	}
	return false
}

func Part1(input string) string {
	equations := parseInput(input)
	sum := 0

	operators := []string{"+", "*"}
	for _, equation := range equations {
		for _, op := range operators {
			if evaluate(equation.Numbers[0], equation.Numbers[1:], op, equation.TestValue, operators) {
				sum += equation.TestValue
				break
			}
		}
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	equations := parseInput(input)
	sum := 0
	operators := []string{"+", "*", "||"}
	for _, equation := range equations {
		for _, op := range operators {
			if evaluate(equation.Numbers[0], equation.Numbers[1:], op, equation.TestValue, operators) {
				sum += equation.TestValue
				break
			}
		}
	}
	return strconv.Itoa(sum)
}
