package day21

import (
	"fmt"
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

// type MemoKey struct {
// 	code  string
// 	level int
// }
//
// var memo map[MemoKey]string

func findSequence(s rune, code string, level int, limit int) string {
	pad := dirpad
	if level == 1 {
		pad = keypad
	}
	// if level > 1 {
	// 	if elem, ok := memo[MemoKey{code: code, level: level}]; ok {
	// 		fmt.Println("memo")
	// 		return elem
	// 	}
	// }

	result := ""

	start := findRune(pad, s)
	for _, key := range code {
		target := findRune(pad, key)
		paths := findPaths(start, target, pad)
		var mapped []string
		for _, path := range paths {
			code := mapPath(path)
			if level < limit {
				mapped = append(mapped, findSequence('A', code, level+1, limit))
			} else {
				mapped = append(mapped, code)
			}
		}
		slices.SortFunc(mapped, func(a, b string) int {
			return len(a) - len(b)
		})
		result += mapped[0]
		start = target
	}

	// if level > 1 {
	// 	memo[MemoKey{code: code, level: level}] = result
	// }
	return result
}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	// memo = make(map[MemoKey]string)

	result := 0
	for _, line := range lines {
		seq := findSequence('A', line, 1, 3)
		fmt.Println(seq)
		fmt.Println(len(seq))
		length := len(seq)
		num, err := strconv.Atoi(line[:3])
		if err != nil {
			fmt.Println("error converting number")
		}
		result += num * length
	}

	return strconv.Itoa(result)
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	// memo = make(map[MemoKey]string)
	//
	result := 0
	prev := make([]int, len(lines))
	for i := 2; i < 10; i++ {
		for l, line := range lines {
			seq := findSequence('A', line, 1, i)
			length := len(seq)
			num, err := strconv.Atoi(line[:3])
			if err != nil {
				fmt.Println("error converting number")
			}
			result += num * length
			fmt.Printf("%s: %d - delta: %d\n", line, length, length-prev[l])
			prev[l] = length
		}
	}

	// for i := 2; i < 10; i++ {
	// 	result := findSequence('A', "<", 2, i)
	// 	fmt.Printf("%d - %s - len: %d - %s\n", i, "<", len(result), result)
	// 	result = findSequence('A', "v", 2, i)
	// 	fmt.Printf("%d - %s - len: %d - %s\n", i, "v", len(result), result)
	// 	result = findSequence('A', "^", 2, i)
	// 	fmt.Printf("%d - %s - len: %d - %s\n", i, "^", len(result), result)
	// 	result = findSequence('A', ">", 2, i)
	// 	fmt.Printf("%d - %s - len: %d - %s\n", i, ">", len(result), result)
	// 	fmt.Println()
	//
	// }
	return ""
}
