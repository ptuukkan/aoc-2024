package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0

	results := re.FindAllStringSubmatch(input, -1)
	for _, result := range results {
		a, err_a := strconv.Atoi(result[1])
		b, err_b := strconv.Atoi(result[2])
		if err_a != nil || err_b != nil {
			fmt.Println("error converting numbers")
		}
		sum += a * b
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	input = strings.ReplaceAll(input, "\n", "")

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	disabledRe := regexp.MustCompile(`don't\(\)(.*?)(do\(\)|$)`)

	input = disabledRe.ReplaceAllString(input, "")

	sum := 0

	results := re.FindAllStringSubmatch(input, -1)
	for _, result := range results {
		a, err_a := strconv.Atoi(result[1])
		b, err_b := strconv.Atoi(result[2])
		if err_a != nil || err_b != nil {
			fmt.Println("error converting numbers")
		}
		sum += a * b
	}
	return strconv.Itoa(sum)
}
