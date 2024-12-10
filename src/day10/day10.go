package day10

import (
	"errors"
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Point struct {
	Y int
	X int
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

var dirs = []Point{
	newPoint(-1, 0),
	newPoint(0, 1),
	newPoint(1, 0),
	newPoint(0, -1),
}

func getCharAtPoint(lines []string, p Point) (rune, error) {
	if p.X < 0 || p.Y < 0 || p.X >= len(lines) || p.Y >= len(lines) {
		return 'x', errors.New("out of bounds")
	}
	return rune(lines[p.Y][p.X]), nil
}

func canClimb(lines []string, p Point, di int) bool {
	a, err_a := getCharAtPoint(lines, p)
	b, err_b := getCharAtPoint(lines, p.Add(dirs[di]))

	if err_a != nil || err_b != nil {
		return false
	}
	if b-a == 1 {
		return true
	}
	return false
}

func hike(lines []string, p Point, peaks map[Point]bool) int {
	if current, _ := getCharAtPoint(lines, p); current == '9' {
		peaks[p] = true
		return 1
	}

	sum := 0
	for i := 0; i < 4; i++ {
		if canClimb(lines, p, i) {
			sum += hike(lines, p.Add(dirs[i]), peaks)
		}
	}
	return sum
}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	sum := 0
	for y, line := range lines {
		for x, c := range line {
			if c == '0' {
				peaks := make(map[Point]bool)
				hike(lines, newPoint(y, x), peaks)
				sum += len(peaks)
			}
		}
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	sum := 0
	for y, line := range lines {
		for x, c := range line {
			if c == '0' {
				peaks := make(map[Point]bool)
				sum += hike(lines, newPoint(y, x), peaks)
			}
		}
	}
	return strconv.Itoa(sum)
}
