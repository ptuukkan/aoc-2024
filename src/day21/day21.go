package day21

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

var keypad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}
var dirpad = [][]rune{
	{'X', '^', 'A'},
	{'<', 'v', '>'},
}

func findRune(lines [][]rune, target rune) utils.Point {
	for y, line := range lines {
		for x, r := range line {
			if r == target {
				return utils.NewPoint(y, x)
			}
		}
	}

	fmt.Println("could not find target rune")
	return utils.NewPoint(0, 0)
}

type Vertex struct {
	Position utils.Point
	Cost     int
	trail    []utils.Point
}

func newVertex(p utils.Point, cost int, trail []utils.Point) Vertex {
	newTrail := make([]utils.Point, len(trail))
	copy(newTrail, trail)
	newTrail = append(newTrail, p)
	return Vertex{Position: p, Cost: cost, trail: newTrail}
}

func findPaths(p1, p2 utils.Point, pad [][]rune) [][]utils.Point {
	var queue []Vertex
	var paths [][]utils.Point

	vertex := Vertex{Position: p1, Cost: 0, trail: []utils.Point{p1}}
	queue = append(queue, vertex)

	for len(queue) != 0 {
		vertex = queue[0]
		queue = queue[1:]

		if vertex.Position == p2 {
			paths = append(paths, vertex.trail)
		} else {
			x := vertex.Position.X
			y := vertex.Position.Y
			dx := p2.X - vertex.Position.X
			dy := p2.Y - vertex.Position.Y
			sx := 0
			sy := 0
			if dx > 0 {
				sx = 1
			} else if dx < 0 {
				sx = -1
			}
			if dy > 0 {
				sy = 1
			} else if dy < 0 {
				sy = -1
			}
			if sy != 0 && pad[y+sy][x] != 'X' {
				queue = append(queue, newVertex(utils.NewPoint(y+sy, x), vertex.Cost+1, vertex.trail))
			}
			if sx != 0 && pad[y][x+sx] != 'X' {
				queue = append(queue, newVertex(utils.NewPoint(y, x+sx), vertex.Cost+1, vertex.trail))
			}
		}
	}
	return paths
}

func mapDirection(p1, p2 utils.Point) rune {
	p := p2.Subtract(&p1)
	index := slices.Index(utils.Directions, p)
	if index == -1 {
		fmt.Println(p1)
		fmt.Println(p2)
	}
	dirs := []rune{'^', '>', 'v', '<'}
	return dirs[index]
}

func mapPath(path []utils.Point) string {
	code := ""
	for i, p := range path {
		if i > 0 {
			dir := mapDirection(path[i-1], p)
			code += string(dir)
		}
	}
	code += "A"
	return code
}

func createPathMap(pad [][]rune) map[[2]rune][]string {
	pathMap := make(map[[2]rune][]string)
	for start_y, start_line := range pad {
		for start_x, start_r := range start_line {
			if start_r == 'X' {
				continue
			}
			start := utils.NewPoint(start_y, start_x)
			for end_y, end_line := range pad {
				for end_x, end_r := range end_line {
					if end_r == 'X' {
						continue
					}
					end := utils.NewPoint(end_y, end_x)
					paths := findPaths(start, end, pad)
					for _, path := range paths {
						code := mapPath(path)
						if elem, ok := pathMap[[2]rune{start_r, end_r}]; ok {
							elem = append(elem, code)
							pathMap[[2]rune{start_r, end_r}] = elem
						} else {
							pathMap[[2]rune{start_r, end_r}] = []string{code}
						}
					}
				}
			}
		}
	}

	return pathMap
}

func mapCode(start rune, code string, pathMap map[[2]rune][]string) []string {
	if len(code) == 0 {
		return []string{}
	}
	end := rune(code[0])
	paths := pathMap[[2]rune{start, end}]
	result := []string{}
	for _, path := range paths {
		subPaths := mapCode(end, code[1:], pathMap)
		if len(subPaths) > 0 {
			for _, subPath := range subPaths {
				result = append(result, path+subPath)
			}
		} else {
			result = append(result, path)
		}
	}

	return result
}

func mapCost(start rune, code string, costMap map[[2]rune]int) int {
	totalCost := 0
	for _, end := range code {
		key := [2]rune{start, end}
		cost := costMap[key]
		totalCost += cost
		start = end
	}

	return totalCost
}

func createCostMap(pathMap map[[2]rune][]string) map[[2]rune]int {
	costMap := make(map[[2]rune]int)

	for _, start_line := range dirpad {
		for _, start_r := range start_line {
			if start_r == 'X' {
				continue
			}
			for _, end_line := range dirpad {
				for _, end_r := range end_line {
					if end_r == 'X' {
						continue
					}
					key := [2]rune{start_r, end_r}
					path := pathMap[key]
					costMap[key] = len(path[0])
				}
			}
		}
	}
	return costMap
}
func createSecondCostMap(firstCostMap map[[2]rune]int, pathMap map[[2]rune][]string) map[[2]rune]int {
	costMap := make(map[[2]rune]int)

	for _, start_line := range dirpad {
		for _, start_r := range start_line {
			if start_r == 'X' {
				continue
			}
			for _, end_line := range dirpad {
				for _, end_r := range end_line {
					if end_r == 'X' {
						continue
					}
					key := [2]rune{start_r, end_r}
					paths := pathMap[key]
					minCost := math.MaxInt
					for _, path := range paths {
						cost := mapCost('A', path, firstCostMap)
						if cost < minCost {
							minCost = cost
						}
					}
					costMap[key] = minCost
				}
			}
		}
	}
	return costMap
}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	dirPathMap := createPathMap(dirpad)
	keyPathMap := createPathMap(keypad)
	costMap := createCostMap(dirPathMap)
	secondCostMap := createSecondCostMap(costMap, dirPathMap)

	result := 0
	for _, line := range lines {
		minCost := math.MaxInt
		paths := mapCode('A', line, keyPathMap)
		for _, p := range paths {
			cost := mapCost('A', p, secondCostMap)
			if cost < minCost {
				minCost = cost
			}
		}

		num, err := strconv.Atoi(line[:3])
		if err != nil {
			fmt.Println("error")
		}
		result += minCost * num
	}

	return fmt.Sprint(result)
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	dirPathMap := createPathMap(dirpad)
	keyPathMap := createPathMap(keypad)
	costMap := createCostMap(dirPathMap)
	secondCostMap := createSecondCostMap(costMap, dirPathMap)

	for i := 0; i < 23; i++ {
		secondCostMap = createSecondCostMap(secondCostMap, dirPathMap)
	}

	result := 0
	for _, line := range lines {
		minCost := math.MaxInt
		paths := mapCode('A', line, keyPathMap)
		for _, p := range paths {
			cost := mapCost('A', p, secondCostMap)
			if cost < minCost {
				minCost = cost
			}
		}

		num, err := strconv.Atoi(line[:3])
		if err != nil {
			fmt.Println("error")
		}
		result += minCost * num
	}

	return fmt.Sprint(result)
}
