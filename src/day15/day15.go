package day15

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

func parseInput(input string) ([][]rune, string) {
	split := strings.Split(input, "\n\n")
	lines := utils.SplitNewLines(split[0])
	moves := strings.ReplaceAll(split[1], "\n", "")

	warehouse := make([][]rune, 0, len(lines))
	for _, line := range lines {
		warehouse = append(warehouse, []rune(line))
	}

	return warehouse, moves
}

func printWarehouse(warehouse [][]rune) {
	for _, line := range warehouse {
		fmt.Println(string(line))
	}
}

func findRobot(warehouse [][]rune) (*utils.Point, error) {
	for y, line := range warehouse {
		for x, r := range line {
			if r == '@' {
				return &utils.Point{Y: y, X: x}, nil
			}
		}
	}
	return nil, errors.New("Could not find robot")
}

func charAtPoint(warehouse [][]rune, p utils.Point) rune {
	return warehouse[p.Y][p.X]
}

func canPush(warehouse [][]rune, point utils.Point, dir int) bool {
	direction := utils.Directions[dir]
	tile := charAtPoint(warehouse, point)
	switch tile {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		return canPush(warehouse, point.Add(direction), dir)
	case '[', ']':
		if dir%2 == 1 {
			return canPush(warehouse, point.Add(direction), dir)
		}

		otherPoint := point.Right()
		if tile == ']' {
			otherPoint = point.Left()
		}

		return canPush(warehouse, point.Add(direction), dir) && canPush(warehouse, otherPoint.Add(direction), dir)
	}
	return false
}

func push(warehouse [][]rune, point utils.Point, dir int) {

	direction := utils.Directions[dir]

	tile := charAtPoint(warehouse, point)

	switch tile {
	case '.':
		return
	case 'O':
		push(warehouse, point.Add(direction), dir)
		warehouse[point.Y][point.X] = '.'
		point = point.Add(direction)
		warehouse[point.Y][point.X] = tile
		return
	case '[', ']':
		if dir%2 == 1 {
			push(warehouse, point.Add(direction), dir)
			warehouse[point.Y][point.X] = '.'
			point = point.Add(direction)
			warehouse[point.Y][point.X] = tile
			return
		}
		otherPoint := point.Right()
		otherChar := ']'
		if tile == ']' {
			otherPoint = point.Left()
			otherChar = '['
		}

		push(warehouse, point.Add(direction), dir)
		push(warehouse, otherPoint.Add(direction), dir)

		warehouse[point.Y][point.X] = '.'
		warehouse[otherPoint.Y][otherPoint.X] = '.'
		point = point.Add(direction)
		otherPoint = otherPoint.Add(direction)
		warehouse[point.Y][point.X] = tile
		warehouse[otherPoint.Y][otherPoint.X] = otherChar
	}
}

func move(warehouse [][]rune, robot *utils.Point, m rune) {
	moves := []rune{'^', '>', 'v', '<'}
	dir := slices.Index(moves, m)

	direction := utils.Directions[dir]
	nextTile := charAtPoint(warehouse, robot.Add(direction))

	switch nextTile {
	case '.':
		warehouse[robot.Y][robot.X] = '.'
		robot.Move(direction)
		warehouse[robot.Y][robot.X] = '@'
		return
	case '#':
		return
	case 'O', '[', ']':
		if canPush(warehouse, robot.Add(direction), dir) {
			push(warehouse, robot.Add(direction), dir)
			warehouse[robot.Y][robot.X] = '.'
			robot.Move(direction)
			warehouse[robot.Y][robot.X] = '@'
		}
		return
	}

}

func Part1(input string) string {
	warehouse, moves := parseInput(input)

	robot, err := findRobot(warehouse)
	if err != nil {
		return err.Error()
	}

	for _, m := range moves {
		move(warehouse, robot, m)
	}

	totalGps := 0

	for top, line := range warehouse {
		for left, r := range line {
			if r == 'O' {
				totalGps += 100*top + left
			}
		}
	}

	return strconv.Itoa(totalGps)
}

func getFat(warehouse [][]rune) {
	for y, line := range warehouse {
		widelane := make([]rune, len(line)*2)
		w := 0
		for _, r := range line {
			switch r {
			case '#':
				widelane[w] = '#'
				widelane[w+1] = '#'
			case 'O':
				widelane[w] = '['
				widelane[w+1] = ']'
			case '.':
				widelane[w] = '.'
				widelane[w+1] = '.'
			case '@':
				widelane[w] = '@'
				widelane[w+1] = '.'
			}
			w += 2
		}
		warehouse[y] = widelane
	}
}

func Part2(input string) string {
	warehouse, moves := parseInput(input)
	getFat(warehouse)

	robot, err := findRobot(warehouse)
	if err != nil {
		return err.Error()
	}

	for _, m := range moves {
		move(warehouse, robot, m)
	}

	totalGps := 0

	for top, line := range warehouse {
		for left, r := range line {
			if r == '[' {
				totalGps += 100*top + left
			}
		}
	}

	return strconv.Itoa(totalGps)
}
