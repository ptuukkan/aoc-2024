package day06

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Add(o Point) Point {
	return Point{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func newPoint(y, x int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func findGuard(lines []string) (Point, error) {
	for y, line := range lines {
		for x, c := range line {
			if c == '^' {
				return newPoint(y, x), nil
			}
		}
	}

	return newPoint(0, 0), errors.New("No guard")
}

var directions = []Point{
	newPoint(-1, 0),
	newPoint(0, 1),
	newPoint(1, 0),
	newPoint(0, -1),
}

func charAtPoint(lines []string, p Point) string {
	return string(lines[p.Y][p.X])
}

func walk(lines []string, guard Point, di int, trail map[Point]map[int]bool) bool {
	tile, ok := trail[guard]
	if !ok {
		trail[guard] = make(map[int]bool)
		trail[guard][di] = true
	} else if tile[di] {
		return true
	}

	nextPosition := guard.Add(directions[di])
	if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= len(lines) || nextPosition.Y >= len(lines) {
		return false
	}

	nextTile := charAtPoint(lines, nextPosition)
	if nextTile == "#" {
		return walk(lines, guard, (di+1)%4, trail)
	} else {
		return walk(lines, nextPosition, di, trail)
	}
}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	guard, err := findGuard(lines)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	trail := make(map[Point]map[int]bool)

	walk(lines, guard, 0, trail)

	return strconv.Itoa(len(trail))
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	guard, err := findGuard(lines)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	trail := make(map[Point]map[int]bool)
	walk(lines, guard, 0, trail)

	loopPositions := 0
	for tile, _ := range trail {
		if tile == guard {
			continue
		}
		bytes := []byte(lines[tile.Y])
		bytes[tile.X] = '#'
		lines[tile.Y] = string(bytes)
		if walk(lines, guard, 0, make(map[Point]map[int]bool)) {
			loopPositions++
		}
		bytes = []byte(lines[tile.Y])
		bytes[tile.X] = '.'
		lines[tile.Y] = string(bytes)
	}

	return strconv.Itoa(loopPositions)
}
