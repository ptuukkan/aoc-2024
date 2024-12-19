package day19

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func parseInput(input string) ([]string, []string) {
	split := strings.Split(input, "\n\n")
	towels := strings.Split(split[0], ", ")
	slices.SortFunc(towels, func(a, b string) int {
		return len(b) - len(a)
	})
	designs := utils.SplitNewLines(split[1])

	return towels, designs
}

var memo map[string]int

func possibleCombinations(design string, towels []string) int {
	if result, ok := memo[design]; ok {
		return result
	}
	if len(design) == 0 {
		return 1
	}
	var possibleTowels []string
	for _, t := range towels {
		if strings.HasPrefix(design, t) {
			possibleTowels = append(possibleTowels, t)
		}
	}

	combinations := 0
	for _, t := range possibleTowels {
		combinations += possibleCombinations(design[len(t):], towels)
	}
	memo[design] = combinations
	return combinations
}

func Part1(input string) string {
	towels, designs := parseInput(input)

	memo = make(map[string]int)
	possibleDesigns := 0
	for _, design := range designs {
		if possibleCombinations(design, towels) > 0 {
			fmt.Println(possibleDesigns)
			possibleDesigns++
		}
	}

	return strconv.Itoa(possibleDesigns)
}

func Part2(input string) string {
	towels, designs := parseInput(input)

	memo = make(map[string]int)
	combinations := 0
	for _, design := range designs {
		combinations += possibleCombinations(design, towels)
	}

	return strconv.Itoa(combinations)
}
