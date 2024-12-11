package day11

import (
	"fmt"
	"strconv"
	"strings"
)

var memo map[[2]int]int

func parseInput(input string) []int {
	var stones []int
	fields := strings.Fields(input)
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("error converting number")
			return stones
		}
		stones = append(stones, num)
	}
	return stones
}

func blink(num int, blinksLeft int) int {
	blinksLeft--
    
	if result, ok := memo[[2]int{num, blinksLeft}]; ok {
		return result
	}

	if num == 0 {
		newNum := 1
		if blinksLeft > 0 {
			result := blink(newNum, blinksLeft)
			memo[[2]int{num, blinksLeft}] = result
			return result
		}
		return 1
	}

	numStr := strconv.Itoa(num)

	if len(numStr)%2 == 0 {
		split_1, err_1 := strconv.Atoi(numStr[:len(numStr)/2])
		split_2, err_2 := strconv.Atoi(numStr[len(numStr)/2:])
		if err_1 != nil || err_2 != nil {
			fmt.Println("error converting number")
			return 0
		}
		if blinksLeft > 0 {
			result := blink(split_1, blinksLeft) + blink(split_2, blinksLeft)
			memo[[2]int{num, blinksLeft}] = result
			return result
		}
		return 2

	}

	newNum := num * 2024

	if blinksLeft > 0 {
		result := blink(newNum, blinksLeft)
		memo[[2]int{num, blinksLeft}] = result
		return result
	}
	return 1
}

func Part1(input string) string {
	stones := parseInput(input)

	memo = make(map[[2]int]int)

	sum := 0
	for _, stone := range stones {
		sum += blink(stone, 25)
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	stones := parseInput(input)

	memo = make(map[[2]int]int)

	sum := 0
	for _, stone := range stones {
		sum += blink(stone, 75)
	}

	return strconv.Itoa(sum)
}
