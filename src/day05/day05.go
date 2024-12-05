package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func parsePageNumbers(input string) [][]int {
	split := strings.Split(input, "\n\n")
	lines := utils.SplitNewLines(split[1])

	pageNumbers := make([][]int, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, ",")
		pageNumbers[i] = make([]int, len(numbers))
		for k, n := range numbers {
			num, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println("failed to convert numbers")
			}
			pageNumbers[i][k] = num

		}
	}

	return pageNumbers
}

func parseRules(input string) [][]int {
	split := strings.Split(input, "\n\n")
	lines := utils.SplitNewLines(split[0])

	rules := make([][]int, len(lines))

	for i, line := range lines {
		rules[i] = make([]int, 2)
		rule := strings.Split(line, "|")
		for k, r := range rule {
			num, err := strconv.Atoi(r)
			if err != nil {
				fmt.Println("error converting number")
			}
			rules[i][k] = num
		}

	}
	return rules
}

func isSafeManual(manual []int, rules [][]int) bool {
	if len(manual) == 1 {
		return true
	}
	var pagesAfter []int
	currentPage := manual[0]

	for _, rule := range rules {
		if currentPage == rule[0] {
			pagesAfter = append(pagesAfter, rule[1])
		}
	}

	for _, page := range manual[1:] {
		if !slices.Contains(pagesAfter, page) {
			return false
		}
	}

	return isSafeManual(manual[1:], rules)
}

func reorderManual(manual []int, rules [][]int) []int {
	pagesAfter := make(map[int][]int)
	for _, page := range manual {
		pagesAfter[page] = []int{}
		for _, rule := range rules {
			if page == rule[0] && slices.Contains(manual, rule[1]) {
				pagesAfter[page] = append(pagesAfter[page], rule[1])
			}
		}
	}

	keys := make([]int, 0, len(pagesAfter))
	for k := range pagesAfter {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return len(pagesAfter[b]) - len(pagesAfter[a])
	})

	return keys
}

func Part1(input string) string {
	rules := parseRules(input)
	manuals := parsePageNumbers(input)

	sum := 0
	for _, manual := range manuals {
		if isSafeManual(manual, rules) {
			sum += manual[len(manual)/2]
		}
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	rules := parseRules(input)
	manuals := parsePageNumbers(input)

	manuals = slices.DeleteFunc(manuals, func(manual []int) bool {
		return isSafeManual(manual, rules)
	})

	sum := 0
	for _, manual := range manuals {
		manual = reorderManual(manual, rules)
		sum += manual[len(manual)/2]
	}

	return strconv.Itoa(sum)
}
