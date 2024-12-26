package day25

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func parseThing(input []string) []int {
	thing := make([]int, 5)

	for _, line := range input[1:] {
		for i, r := range line {
			if r == '#' {
				thing[i]++
			}
		}
	}

	return thing
}

func parseInput(input string) ([][]int, [][]int) {
	locks := [][]int{}
	keys := [][]int{}
	split := strings.Split(input, "\n\n")

	for _, splt := range split {
		spl := utils.SplitNewLines(splt)
		if spl[0] == "#####" {
			locks = append(locks, parseThing(spl))
		} else {
			slices.Reverse(spl)
			keys = append(keys, parseThing(spl))
		}
	}

	return locks, keys
}

func fits(lock []int, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func Part1(input string) string {
	locks, keys := parseInput(input)

	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			if fits(lock, key) {
				result++
			}
		}
	}

	return fmt.Sprint(result)
}

func Part2(input string) string {
	return ""
}
