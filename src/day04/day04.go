package day04

import (
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

var lines []string
var length int

var directions = []Point{
	newPoint(0, 1),
	newPoint(1, 1),
	newPoint(1, 0),
	newPoint(1, -1),
	newPoint(0, -1),
	newPoint(-1, 1),
	newPoint(-1, -1),
	newPoint(-1, 0),
}

func findAllIndexes(s string, letter rune) []int {
	var indexes []int
	for i, r := range s {
		if r == letter {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func getCharAtPoint(p Point) string {
	return string(lines[p.Y][p.X])
}

func collect(p Point, dir Point) int {
	s := ""

	for i := 0; i < 4; i++ {
		if p.X < 0 || p.X >= length || p.Y < 0 || p.Y >= length {
			return 0
		}
		s += getCharAtPoint(p)
		p = p.Add(dir)
	}
	if s == "XMAS" {
		return 1
	}
	return 0
}

func findXmas(point Point) int {
	result := 0
	for _, dir := range directions {
		result += collect(point, dir)
	}
	return result
}

func findX_Mas(point Point) int {
	if point.X == 0 || point.X == length-1 || point.Y == 0 || point.Y == length-1 {
		return 0
	}
	corners := make([]Point, 4)
	corners[0] = newPoint(point.Y-1, point.X-1)
	corners[1] = newPoint(point.Y-1, point.X+1)
	corners[2] = newPoint(point.Y+1, point.X+1)
	corners[3] = newPoint(point.Y+1, point.X-1)

	mas_count := 0
	for i, corner := range corners {
		if getCharAtPoint(corner) == "M" && getCharAtPoint(corners[(i+2)%4]) == "S" {
			mas_count++
		}
	}
	if mas_count == 2 {
		return 1
	}
	return 0
}

func Part1(input string) string {
	lines = utils.SplitNewLines(input)
	length = len(lines)
	var points []Point

	for y, line := range lines {
		indexes := findAllIndexes(line, 'X')
		for _, index := range indexes {
			points = append(points, newPoint(y, index))
		}
	}

	xmas_count := 0
	for _, point := range points {
		xmas_count += findXmas(point)
	}

	return strconv.Itoa(xmas_count)
}

func Part2(input string) string {
	lines = utils.SplitNewLines(input)
	length = len(lines)
	var points []Point

	for y, line := range lines {
		indexes := findAllIndexes(line, 'A')
		for _, index := range indexes {
			points = append(points, newPoint(y, index))
		}
	}

	xmas_count := 0
	for _, point := range points {
		xmas_count += findX_Mas(point)
	}

	return strconv.Itoa(xmas_count)
}
