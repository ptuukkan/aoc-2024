package day02

import (
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func normalize(number int) int {
	if number > 0 {
		return 1
	}
	if number < 0 {
		return -1
	}
	return 0
}

func parseInput(input string) [][]int {
	lines := utils.SplitNewLines(input)
	reports := make([][]int, len(lines))

	for i, line := range lines {
		levels := strings.Fields(line)
		var levelNumbers []int
		for _, n := range levels {
			num, _ := strconv.Atoi(n)
			levelNumbers = append(levelNumbers, num)
		}
		reports[i] = levelNumbers
	}

	return reports
}

func checkReport(report []int) (bool, int) {
	length := len(report)
	var dir *int = nil
	for i, level := range report {
		if i+1 == length {
			break
		}
		diff := level - report[i+1]
		if dir == nil {
			dir = new(int)
			*dir = normalize(diff)
		}
		if utils.Abs(diff) > 3 || utils.Abs(diff) < 1 || normalize(diff) != *dir {
			return false, i
		}
	}
	return true, 0
}

func Part1(input string) string {
	reports := parseInput(input)

	var safeReports = 0

	for _, report := range reports {
		safe, _ := checkReport(report)
		if safe {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports)
}

func Part2(input string) string {
	reports := parseInput(input)

	var safeReports = 0

	for _, report := range reports {
		safe, _ := checkReport(report)
		if !safe {
			for i, _ := range report {
				dampened := append([]int{}, report[:i]...)
				dampened = append(dampened, report[i+1:]...)
				safe, _ := checkReport(dampened)
				if safe {
					safeReports++
					break
				}
			}
		}
		if safe {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports)
}
