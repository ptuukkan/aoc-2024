package main

import (
	"fmt"
	"os"

	"github.com/ptuukkan/aoc-2024/src/day01"
	"github.com/ptuukkan/aoc-2024/src/day02"
	"github.com/ptuukkan/aoc-2024/src/day03"
	"github.com/ptuukkan/aoc-2024/src/day04"
	"github.com/ptuukkan/aoc-2024/src/day05"
	"github.com/ptuukkan/aoc-2024/src/day06"
	"github.com/ptuukkan/aoc-2024/src/day07"
	"github.com/ptuukkan/aoc-2024/src/day08"
	"github.com/ptuukkan/aoc-2024/src/day09"
	"github.com/ptuukkan/aoc-2024/src/day10"
	"github.com/ptuukkan/aoc-2024/src/day11"
	"github.com/ptuukkan/aoc-2024/src/day12"
	"github.com/ptuukkan/aoc-2024/src/day13"
	"github.com/ptuukkan/aoc-2024/src/day14"
	"github.com/ptuukkan/aoc-2024/src/day15"
	"github.com/ptuukkan/aoc-2024/src/day16"
	"github.com/ptuukkan/aoc-2024/src/day17"
	"github.com/ptuukkan/aoc-2024/src/day18"
	"github.com/ptuukkan/aoc-2024/src/day19"
	"github.com/ptuukkan/aoc-2024/src/day20"
	"github.com/ptuukkan/aoc-2024/src/day21"
	"github.com/ptuukkan/aoc-2024/src/day22"
	"github.com/ptuukkan/aoc-2024/src/day23"
	"github.com/ptuukkan/aoc-2024/src/day24"
	"github.com/ptuukkan/aoc-2024/src/day25"
	"github.com/ptuukkan/aoc-2024/utils"
)

type HandlerFunc func(input string) string

func createHandlers() map[string]map[string]HandlerFunc {
	handlers := map[string]map[string]HandlerFunc{
		"day01": {
			"part1": day01.Part1,
			"part2": day01.Part2,
		},
		"day02": {
			"part1": day02.Part1,
			"part2": day02.Part2,
		},
		"day03": {
			"part1": day03.Part1,
			"part2": day03.Part2,
		},
		"day04": {
			"part1": day04.Part1,
			"part2": day04.Part2,
		},
		"day05": {
			"part1": day05.Part1,
			"part2": day05.Part2,
		},
		"day06": {
			"part1": day06.Part1,
			"part2": day06.Part2,
		},
		"day07": {
			"part1": day07.Part1,
			"part2": day07.Part2,
		},
		"day08": {
			"part1": day08.Part1,
			"part2": day08.Part2,
		},
		"day09": {
			"part1": day09.Part1,
			"part2": day09.Part2,
		},
		"day10": {
			"part1": day10.Part1,
			"part2": day10.Part2,
		},
		"day11": {
			"part1": day11.Part1,
			"part2": day11.Part2,
		},
		"day12": {
			"part1": day12.Part1,
			"part2": day12.Part2,
		},
		"day13": {
			"part1": day13.Part1,
			"part2": day13.Part2,
		},
		"day14": {
			"part1": day14.Part1,
			"part2": day14.Part2,
		},
		"day15": {
			"part1": day15.Part1,
			"part2": day15.Part2,
		},
		"day16": {
			"part1": day16.Part1,
			"part2": day16.Part2,
		},
		"day17": {
			"part1": day17.Part1,
			"part2": day17.Part2,
		},
		"day18": {
			"part1": day18.Part1,
			"part2": day18.Part2,
		},
		"day19": {
			"part1": day19.Part1,
			"part2": day19.Part2,
		},
		"day20": {
			"part1": day20.Part1,
			"part2": day20.Part2,
		},
		"day21": {
			"part1": day21.Part1,
			"part2": day21.Part2,
		},
		"day22": {
			"part1": day22.Part1,
			"part2": day22.Part2,
		},
		"day23": {
			"part1": day23.Part1,
			"part2": day23.Part2,
		},
		"day24": {
			"part1": day24.Part1,
			"part2": day24.Part2,
		},
		"day25": {
			"part1": day25.Part1,
			"part2": day25.Part2,
		},
	}
	return handlers
}

func main() {
	args := os.Args[1:]
	handlers := createHandlers()
	day := handlers[args[0]]
	if day == nil {
		fmt.Println("Invalid day")
		return
	}
	handler := handlers[args[0]][args[1]]
	if handler == nil {
		fmt.Println("Invalid part")
		return
	}
	input, err := utils.ReadFile("inputs/" + args[0])
	if err != nil {
		return
	}
	fmt.Println(handler(input))
}
