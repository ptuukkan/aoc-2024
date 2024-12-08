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

func (p Point) Add(o Point) Point {
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

var trail map[Point]map[int]bool
var loopPositions = 0

func charAtPoint(lines []string, p Point) string {
	return string(lines[p.Y][p.X])
}

func walk(lines []string, guard Point, di int) {
	_, ok := trail[guard]
	if !ok {
		trail[guard] = make(map[int]bool)
	}

	nextPosition := guard.Add(directions[di])
	if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= len(lines) || nextPosition.Y >= len(lines) {
		return
	}
	nextTile := charAtPoint(lines, nextPosition)

	if nextTile == "#" {
		nextDi := (di + 1) % 4
		nextDirection := directions[nextDi]
		nextPosition = guard.Add(nextDirection)
		walk(lines, nextPosition, nextDi)
	} else {
		walk(lines, nextPosition, di)
	}
}

func populateLine(lines []string, guard Point, di int) {
	if guard.X < 0 || guard.Y < 0 || guard.X >= len(lines) || guard.Y >= len(lines) {
		return
	}

	if charAtPoint(lines, guard) == "#" {
		return
	}

	elem, ok := trail[guard]
	if ok {
		elem[di] = true
	} else {
		trail[guard] = make(map[int]bool)
		trail[guard][di] = true
	}

	oppositeDir := directions[(di+2)%4]
	nextPosition := guard.Add(oppositeDir)

	populateLine(lines, nextPosition, di)
}

func walkPart2(lines []string, guard Point, di int) {
	elem, ok := trail[guard]
	if ok {
		elem[di] = true
		if _, ok := elem[(di+1)%4]; ok {
			loopPositions++
		}
	} else {
		trail[guard] = make(map[int]bool)
		trail[guard][di] = true
	}

	nextPosition := guard.Add(directions[di])
	if nextPosition.X < 0 || nextPosition.Y < 0 || nextPosition.X >= len(lines) || nextPosition.Y >= len(lines) {
		populateLine(lines, guard, di)
		return
	}
	nextTile := charAtPoint(lines, nextPosition)

	if nextTile == "#" {
		populateLine(lines, guard, di)
		nextDi := (di + 1) % 4
		nextDirection := directions[nextDi]
		nextPosition = guard.Add(nextDirection)
		walkPart2(lines, nextPosition, nextDi)
	} else {
		walkPart2(lines, nextPosition, di)
	}

}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	guard, err := findGuard(lines)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	trail = make(map[Point]map[int]bool)

	walk(lines, guard, 0)

	return strconv.Itoa(len(trail))
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	guard, err := findGuard(lines)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	trail = make(map[Point]map[int]bool)

	walkPart2(lines, guard, 0)

	return strconv.Itoa(loopPositions)
}
