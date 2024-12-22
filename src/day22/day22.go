package day22

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func sequenceToString(sequence []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sequence)), ","), "[]")
}

func parseInput(input string) []int {
	lines := utils.SplitNewLines(input)
	initialNumbers := make([]int, len(lines))

	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting numbers")
			continue
		}
		initialNumbers[i] = num
	}
	return initialNumbers
}

func calculateNextSecret(num int) int {
	num = mix(num, num*64)
	num = prune(num)
	num = mix(num, num/32)
	num = prune(num)
	num = mix(num, num*2048)
	num = prune(num)
	return num
}

func mix(a, b int) int {
	return a ^ b
}

func prune(num int) int {
	return num % 16777216
}

func windowed(slice []int, size int) [][]int {
	windows := make([][]int, size)
	var result [][]int

	for i, num := range slice {
		index := i % size
		windows[index] = make([]int, 0, size)
		for k := 0; k < size; k++ {
			if windows[k] != nil {
				windows[k] = append(windows[k], num)
			}
			if len(windows[k]) == size {
				result = append(result, windows[k])
			}
		}
	}
	return result
}

func Part1(input string) string {
	initialNumbers := parseInput(input)
	length := 2000

	result := 0
	for _, num := range initialNumbers {
		for i := 0; i < length; i++ {
			num = calculateNextSecret(num)
		}
		result += num
	}

	return fmt.Sprint(result)
}

func Part2(input string) string {
	initialNumbers := parseInput(input)
	length := 2001

	secretNumbers := make([][]int, len(initialNumbers))
	for n, num := range initialNumbers {
		numbers := make([]int, length)
		numbers[0] = num
		for i := 1; i < length; i++ {
			num = calculateNextSecret(num)
			numbers[i] = num
		}
		secretNumbers[n] = numbers
	}

	prices := make([][]int, len(secretNumbers))
	for s, secretNumber := range secretNumbers {
		price := make([]int, length)
		for i, num := range secretNumber {
			price[i] = num % 10
		}
		prices[s] = price
	}

	deltas := make([][]int, len(prices))
	for i, price := range prices {
		delta := make([]int, length)
		for k, p := range price {
			if k == 0 {
				delta[k] = 0
			} else {
				delta[k] = p - price[k-1]
			}
		}
		deltas[i] = delta
	}

	windows := make([]map[string]int, len(deltas))
	for i, delta := range deltas {
		windows[i] = make(map[string]int)
		for k, w := range windowed(delta, 4) {
			key := sequenceToString(w)
			if _, ok := windows[i][key]; !ok {
				windows[i][key] = prices[i][k+3]
			}
		}
	}

	shared := make(map[string]int)
	for _, window := range windows {
		for key, value := range window {
			if _, ok := shared[key]; ok {
				shared[key] += value
			} else {
				shared[key] = value
			}
		}
	}

	var keys []string
	for key := range shared {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return shared[b] - shared[a]
	})

	return fmt.Sprint(shared[keys[0]])
}
