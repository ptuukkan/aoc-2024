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

func anotherFindBestPath(paths [][]utils.Point, costs map[[2]rune]string) string {
	// minCost := math.MaxInt
	jep := ""
	for _, path := range paths {
		cost := ""

		start := 'A'
		for _, c := range path {
			target := dirpad[c.Y][c.X]
			elem, ok := costs[[2]rune{start, target}]
			if !ok {
				fmt.Printf("No path from %s to %s\n", string(start), string(target))
			}
			fmt.Printf("b: from %s to %s: %s\n", string(start), string(target), elem)

			cost += elem
			start = target
		}
		if len(jep) == 0 || len(cost) < len(jep) {
			jep = cost
		}
	}
	return jep
}

func findBestPath(paths [][]utils.Point) []utils.Point {
	minCost := math.MaxInt
	minCostIndex := -1
	for pathIndex, path := range paths {
		code := mapPath(path)
		// fmt.Println(code)
		cost := 0
		start := findRune(dirpad, 'A')
		for _, c := range code {
			target := findRune(dirpad, c)
			nextPaths := findPaths(start, target, dirpad)
			pathCosts := make([]int, len(nextPaths))
			for n, p := range nextPaths {
				a := mapPath(p)
				// fmt.Printf("From %s to %s - path: %s\n", string(dirpad[start.Y][start.X]), string(c), a)
				pathCosts[n] = len(a)
			}

			cost += slices.Min(pathCosts)
			start = target
		}
		if cost < minCost {
			minCost = cost
			minCostIndex = pathIndex
		}
	}
	return paths[minCostIndex]
	// return minCost
}

func createNextCostMap(moveMap map[[2]rune]string, prevCostMap map[[2]rune]uint64) map[[2]rune]uint64 {
	nextCostMap := make(map[[2]rune]uint64)

	for key := range prevCostMap {
		moves := moveMap[key]

		cost := uint64(0)
		start := 'A'
		for _, m := range moves {
			target := m
			segmentCost := prevCostMap[[2]rune{start, target}]
			cost += segmentCost
			start = m
		}
		nextCostMap[key] = cost

	}
	return nextCostMap
}

func createNextMoveMap(moveMap map[[2]rune]string) map[[2]rune]string {
	nextMoveMap := make(map[[2]rune]string)

	for key, value := range moveMap {

		newMoves := ""
		start := 'A'
		for _, m := range value {
			target := m
			segment := moveMap[[2]rune{start, target}]
			newMoves += segment
			start = m
		}
		nextMoveMap[key] = newMoves

	}
	return nextMoveMap
}

func createInitialCostMap(moveMap map[[2]rune]string) map[[2]rune]uint64 {
	costMap := make(map[[2]rune]uint64)

	for key, value := range moveMap {
		costMap[key] = uint64(len(value))
	}
	return costMap
}

func createCostMap(moveMap map[[2]rune]string, level int) map[[2]rune]uint64 {
	costMap := createInitialCostMap(moveMap)

	for i := 0; i < level; i++ {
		costMap = createNextCostMap(moveMap, costMap)
	}

	return costMap
}

func createMoveMap() map[[2]rune]string {
	moveMap := make(map[[2]rune]string)

	for fy, fline := range dirpad {
		for fx, fr := range fline {
			if fr == 'X' {
				continue
			}
			start := utils.NewPoint(fy, fx)

			for ty, tline := range dirpad {
				for tx, tr := range tline {
					if tr == 'X' {
						continue
					}
					target := utils.NewPoint(ty, tx)
					paths := findPaths(start, target, dirpad)
					bestPath := findBestPath(paths)
					moveMap[[2]rune{dirpad[fy][fx], dirpad[ty][tx]}] = mapPath(bestPath)
				}
			}
		}
	}
	return moveMap
}

func createMoveMaps(moveMap map[[2]rune]string, level int) map[[2]rune]string {

	for i := 0; i < level; i++ {
		moveMap = createNextMoveMap(moveMap)
	}

	return moveMap
}

func newVertex(p utils.Point, cost int, trail []utils.Point) Vertex {
	newTrail := make([]utils.Point, len(trail))
	copy(newTrail, trail)
	newTrail = append(newTrail, p)
	return Vertex{Position: p, Cost: cost, trail: newTrail}
}

func calculateCost(code string, costMap map[[2]rune]uint64) uint64 {
	cost := uint64(0)
	start := 'A'
	for _, c := range code {
		target := c
		segmentCost := costMap[[2]rune{start, target}]
		cost += segmentCost
		start = target
	}
	return cost
}

func calculateMoves(code string, moveMap map[[2]rune]string) string {
	moves := ""
	start := 'A'
	for _, c := range code {
		target := c
		segment := moveMap[[2]rune{start, target}]
		moves += segment
		start = target
	}
	return moves
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

func getMapCost(costMap map[[2]rune]uint64, start, target utils.Point) uint64 {
	possiblePaths := findPaths(start, target, keypad)
	minCost := uint64(math.MaxUint64)
	for _, path := range possiblePaths {
		code := mapPath(path)
		cost := calculateCost(code, costMap)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func getMoves(moveMap map[[2]rune]string, start, target utils.Point) string {
	possiblePaths := findPaths(start, target, keypad)
	moves := ""
	for _, path := range possiblePaths {
		code := mapPath(path)
		cost := calculateMoves(code, moveMap)
		if moves == "" || len(cost) < len(moves) {
			moves = cost
		}
	}
	return moves
}

func Part1(input string) string {
	lines := utils.SplitNewLines(input)
	moveMap := createMoveMap()
	costMap := createCostMap(moveMap, 1)

	result := uint64(0)
	for _, line := range lines {
		start := findRune(keypad, 'A')
		totalCost := uint64(0)
		for _, c := range line {
			target := findRune(keypad, c)
			totalCost += getMapCost(costMap, start, target)
			start = target
		}

		num, err := strconv.Atoi(line[:3])
		if err != nil {
			fmt.Println("error")
		}
		result += totalCost * uint64(num)
	}

	return fmt.Sprint(result)
}

func Part2(input string) string {
	lines := utils.SplitNewLines(input)
	iterations := 4
	moveMap := createMoveMap()
	moveMap = createMoveMaps(moveMap, iterations-2)
	fmt.Println("Cost map -----------------------")
	for key, value := range moveMap {
		fmt.Printf("From %s to %s: %s\n", string(key[0]), string(key[1]), value)
	}
	fmt.Println("End cost map -----------------------")

	seq := findSequence('A', "029A", 1, iterations)
	fmt.Println(seq)
	fmt.Println(len(seq))

	result := uint64(0)
	for _, line := range lines {
		start := findRune(keypad, 'A')
		moves := ""
		for _, c := range line {
			target := findRune(keypad, c)
			moves += getMoves(moveMap, start, target)
			start = target
		}
		fmt.Println(moves)
		fmt.Println(len(moves))

		if seq == moves && len(seq) == len(moves) {
			fmt.Println("equal")
		}
	}

	return fmt.Sprint(result)
}

// func Part2(input string) string {
// 	lines := utils.SplitNewLines(input)
// 	iterations := 4
// 	moveMap := createMoveMap()
// 	costMap := createCostMap(moveMap, iterations-2)
//
// 	seq := findSequence('A', "029A", 1, iterations)
// 	fmt.Println(len(seq))
// 	result := uint64(0)
// 	for _, line := range lines {
// 		start := findRune(keypad, 'A')
// 		totalCost := uint64(0)
// 		for _, c := range line {
// 			target := findRune(keypad, c)
// 			totalCost += getMapCost(costMap, start, target)
// 			start = target
// 		}
//
// 		num, err := strconv.Atoi(line[:3])
// 		if err != nil {
// 			fmt.Println("error")
// 		}
// 		fmt.Println(totalCost)
// 		result += uint64(totalCost) * uint64(num)
// 	}
// 	//
// 	return fmt.Sprint(result)
// lines := utils.SplitNewLines(input)
// moveMap := createMoveMap()
// costMap := createCostMap(moveMap, 1)
// fmt.Println("Cost map -----------------------")
// for key, value := range costMap {
// 	fmt.Printf("From %s to %s: %s\n", string(key[0]), string(key[1]), value)
// }
// fmt.Println("End cost map -----------------------")
//
// jep := calculateCost("<A^A>^^AvvvA", costMap)
// fmt.Println(jep)
// // fmt.Println(calculateCost("<", costMap))
//
// result := 0
// for _, line := range lines {
// 	codeCost := 0
// 	start := findRune(keypad, 'A')
// 	for _, x := range line {
// 		target := findRune(keypad, x)
// 		paths := findPaths(start, target, keypad)
// 		minCost := math.MaxInt
// 		for _, p := range paths {
// 			code := mapPath(p)
// 			cost := calculateCost(code, costMap)
// 			if cost < minCost {
// 				minCost = cost
// 			}
// 		}
// 		codeCost += minCost
// 		start = target
// 	}
// 	result += codeCost
// }
// for key, value := range routes {
// 	fmt.Printf("From %s to %s - %d\n", string(key[0]), string(key[1]), value)
// }
//
// for i := 2; i < 5; i++ {
//
// }
// lines := utils.SplitNewLines(input)
// // memo = make(map[MemoKey]string)
// //
// result := 0
// prev := make([]int, len(lines))
// for i := 2; i < 3; i++ {
// 	for l, line := range lines {
// 		seq := findSequence('A', line, 1, i)
// 		length := len(seq)
// 		num, err := strconv.Atoi(line[:3])
// 		if err != nil {
// 			fmt.Println("error converting number")
// 		}
// 		result += num * length
// 		fmt.Printf("%s: %d - delta: %d\n", line, length, length-prev[l])
// 		prev[l] = length
// 	}
// }
//
// 	for i := 2; i < 5; i++ {
// 		result := findSequence('A', "<", 2, i)
// 		fmt.Printf("%d - from %s to %s - len: %d - %s\n", i, "A", "<", len(result), result)
// 		result = findSequence('A', "v", 2, i)
// 		fmt.Printf("%d - from %s to %s - len: %d - %s\n", i, "A", "v", len(result), result)
// 		result = findSequence('A', "^", 2, i)
// 		fmt.Printf("%d - from %s to %s - len: %d - %s\n", i, "A", "^", len(result), result)
// 		result = findSequence('A', ">", 2, i)
// 		fmt.Printf("%d - from %s to %s - len: %d - %s\n", i, "A", ">", len(result), result)
// 		fmt.Println()
// 	}
// 	return fmt.Sprint(result)
// }
