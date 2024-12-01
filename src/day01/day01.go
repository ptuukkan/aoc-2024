package day01

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func createLists(input string) ([]int, []int) {
	lines := utils.SplitNewLines(input)
	right := make([]int, len(lines))
	left := make([]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])
		right[i] = a
		left[i] = b
	}

	return right, left
}

func Part1(input string) string {
	right, left := createLists(input)

	slices.Sort(right)
	slices.Sort(left)

	diffs := make([]int, len(right))

	for i, a := range right {
		b := left[i]
		diff := a - b
		if diff <= 0 {
			diff = -diff
		}
		diffs[i] = diff
	}

	sum := 0

	for _, diff := range diffs {
		sum += diff
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	right, left := createLists(input)

	appearances := make(map[int]int)

	for _, n := range left {
		if _, ok := appearances[n]; ok {
			appearances[n]++
		} else {
			appearances[n] = 1
		}
	}

	for i, n := range right {
		if val, ok := appearances[n]; ok {
			right[i] = n * val
		} else {
			right[i] = 0
		}
	}

	sum := 0

	for _, n := range right {
		sum += n
	}

	return strconv.Itoa(sum)
}
