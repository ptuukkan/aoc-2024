package day08

import (
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Antenna struct {
	Frequency rune
	Y         int
	X         int
}

type Point struct {
	Y int
	X int
}

func newAntenna(f rune, y, x int) Antenna {
	return Antenna{Frequency: f, Y: y, X: x}
}

func newPoint(y, x int) Point {
	return Point{Y: y, X: x}
}

func parseInput(input string) ([]Antenna, int) {
	var antennas []Antenna
	lines := utils.SplitNewLines(input)

	for y, line := range lines {
		for x, c := range line {
			if c != '.' {
				antennas = append(antennas, newAntenna(c, y, x))
			}
		}
	}

	return antennas, len(lines)
}

func Part1(input string) string {
	antennas, length := parseInput(input)
	antinodes := make(map[Point]bool)

	for i, antenna_a := range antennas {
		for k, antenna_b := range antennas {
			if i == k || antenna_a.Frequency != antenna_b.Frequency {
				continue
			}
			x := antenna_a.X + (antenna_a.X - antenna_b.X)
			y := antenna_a.Y + (antenna_a.Y - antenna_b.Y)

			if x >= 0 && x < length && y >= 0 && y < length {
				antinode := newPoint(y, x)
				antinodes[antinode] = true
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func Part2(input string) string {
	antennas, length := parseInput(input)
	antinodes := make(map[Point]bool)

	for i, antenna_a := range antennas {
		for k, antenna_b := range antennas {
			if i == k || antenna_a.Frequency != antenna_b.Frequency {
				continue
			}
			antinodes[newPoint(antenna_a.Y, antenna_a.X)] = true
			antinodes[newPoint(antenna_b.Y, antenna_b.X)] = true

			delta_x := antenna_a.X - antenna_b.X
			delta_y := antenna_a.Y - antenna_b.Y
			x := antenna_a.X
			y := antenna_a.Y
			for {
				x += delta_x
				y += delta_y
				if x < 0 || x >= length || y < 0 || y >= length {
					break
				}
				antinodes[newPoint(y, x)] = true
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}
