package day14

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Robot struct {
	Position utils.Point
	Velocity utils.Point
}

func mod(num int, modulo int) int {
	result := num % modulo
	if result < 0 {
		result += modulo
	}
	return result
}

func convertNumber(str string) (int, int) {
	split := strings.Split(str, ",")
	x, err_x := strconv.Atoi(split[0])
	y, err_y := strconv.Atoi(split[1])

	if err_x != nil || err_y != nil {
		fmt.Println("error converting numbers")
		return 0, 0
	}
	return x, y
}

func parseInput(input string) []*Robot {
	lines := utils.SplitNewLines(input)
	robots := make([]*Robot, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		px, py := convertNumber(strings.ReplaceAll(split[0], "p=", ""))
		vx, vy := convertNumber(strings.ReplaceAll(split[1], "v=", ""))
		p := utils.NewPoint(py, px)
		v := utils.NewPoint(vy, vx)
		robots[i] = &Robot{Position: p, Velocity: v}
	}

	return robots
}

func countQuadrants(positions []utils.Point, height, width int) []int {
	quadrants := make([]int, 4)
	halfheight := height / 2
	halfwidth := width / 2

	for _, p := range positions {
		if p.X < halfwidth && p.Y < halfheight {
			quadrants[0]++
		} else if p.X > halfwidth && p.Y < halfheight {
			quadrants[1]++
		} else if p.X < halfwidth && p.Y > halfheight {
			quadrants[2]++
		} else if p.X > halfwidth && p.Y > halfheight {
			quadrants[3]++
		}
	}

	return quadrants
}

func Part1(input string) string {
	height := 103
	width := 101
	robots := parseInput(input)
	finalPositions := make([]utils.Point, len(robots))

	for i, robot := range robots {
		vx := robot.Velocity.X * 100
		vy := robot.Velocity.Y * 100
		tv := utils.NewPoint(vy, vx)
		p := robot.Position.Add(tv)
		finalPositions[i] = utils.NewPoint(mod(p.Y, height), mod(p.X, width))
	}

	quadrants := countQuadrants(finalPositions, height, width)

	product := 1
	for _, q := range quadrants {
		product *= q
	}

	return strconv.Itoa(product)
}

func getLines(robots []*Robot, height, width int) []string {
	lines := make([][]rune, height)
	for h := 0; h < height; h++ {
		lines[h] = make([]rune, width)
		for i := range lines[h] {
			lines[h][i] = '.'
		}
	}

	for _, robot := range robots {
		lines[robot.Position.Y][robot.Position.X] = '#'
	}

	stringLines := make([]string, len(lines))
	for i, row := range lines {
		stringLines[i] = string(row)
	}
	return stringLines
}

func Part2(input string) string {
	height := 103
	width := 101
	robots := parseInput(input)

	for i := 0; i < 20000; i++ {
		lines := getLines(robots, height, width)
		if (i%width) == 7 && (i%height) == 53 {
			for _, line := range lines {
				fmt.Println(line)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		for _, robot := range robots {
			newPosition := robot.Position.Add(robot.Velocity)
			robot.Position = utils.NewPoint(mod(newPosition.Y, height), mod(newPosition.X, width))
		}
	}
	return ""
}
