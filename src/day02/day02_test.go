package day02

import (
	"testing"

	"github.com/ptuukkan/aoc-2024/utils"
)

var day = "day02"
var part1Expected = "2"
var part2Expected = "4"

func TestPart1(t *testing.T) {
	input, err := utils.ReadFile("../../inputs/" + day + "_test")
	if err != nil {
		t.Fatal("Error getting input")
	}

	actual := Part1(input)

	if actual != part1Expected {
		t.Errorf("Part1: expected=%s, actual=%s", part1Expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadFile("../../inputs/" + day + "_test")
	if err != nil {
		t.Fatal("Error getting input")
	}

	actual := Part2(input)

	if actual != part2Expected {
		t.Errorf("Part2: expected=%s, actual=%s", part2Expected, actual)
	}
}
