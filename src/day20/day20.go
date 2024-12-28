package day20

import (
	"fmt"

	"github.com/ptuukkan/aoc-2024/utils"
)

func findRune(lines []string, target rune) utils.Point {
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

type Node struct {
	Position utils.Point
	Cost     int
	Prev     *Node
}

func newNode(p utils.Point, prev *Node) *Node {
	cost := 0
	if prev != nil {
		cost = prev.Cost + 1
	}
	return &Node{Position: p, Cost: cost, Prev: prev}
}

func dijkstra(maze []string, start utils.Point) map[utils.Point]int {
	visited := make(map[utils.Point]int)
	queue := []*Node{newNode(start, nil)}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node.Position] = node.Cost

		for _, dir := range utils.Directions {
			adj := node.Position.Add(dir)
			if adj.OutOfBounds(len(maze)) {
				continue
			}

			if maze[adj.Y][adj.X] == '#' {
				continue
			}

			if _, ok := visited[adj]; ok {
				continue
			}
			queue = append(queue, newNode(adj, node))
		}
	}

	return visited
}

func pointsWithinRange(maze []string, start utils.Point, cheatRange int) map[utils.Point]int {
	length := len(maze)
	startRow := start.Y - cheatRange
	if startRow < 0 {
		startRow = 0
	}
	endRow := start.Y + cheatRange
	if endRow >= length {
		endRow = length - 1
	}
	startColumn := start.X - cheatRange
	if startColumn < 0 {
		startColumn = 0
	}
	endColumn := start.X + cheatRange
	if endColumn >= length {
		endColumn = length - 1
	}

	result := make(map[utils.Point]int)
	for y := startRow; y <= endRow; y++ {
		for x := startColumn; x <= endColumn; x++ {
			dy := start.Y - y
			dx := start.X - x
			distance := utils.Abs(dy) + utils.Abs(dx)
			if distance <= cheatRange {
				result[utils.NewPoint(y, x)] = distance
			}
		}
	}

	return result
}

func calculateCheats(maze []string, costs, reverseCosts map[utils.Point]int, initialBest int, cheatRange int) map[[2]utils.Point]int {
	result := make(map[[2]utils.Point]int)

	for start, startCost := range costs {
		if startCost >= initialBest {
			continue
		}
		pointsInRange := pointsWithinRange(maze, start, cheatRange)
		for point, length := range pointsInRange {
			if maze[point.Y][point.X] == '#' {
				continue
			}
			cost := startCost + reverseCosts[point] + length
			saves := initialBest - cost
			if saves > 0 {
				result[[2]utils.Point{start, point}] = saves
			}
		}
	}
	return result
}

func Part1(input string) string {
	maze := utils.SplitNewLines(input)
	start := findRune(maze, 'S')
	end := findRune(maze, 'E')
	costs := dijkstra(maze, start)
	reverseCosts := dijkstra(maze, end)

	cheats := calculateCheats(maze, costs, reverseCosts, costs[end], 2)

	result := 0
	for _, value := range cheats {
		fmt.Println(value)
		if value >= 100 {
			result++
		}
	}

	return fmt.Sprint(result)
}

func Part2(input string) string {
	maze := utils.SplitNewLines(input)
	start := findRune(maze, 'S')
	end := findRune(maze, 'E')
	costs := dijkstra(maze, start)
	reverseCosts := dijkstra(maze, end)

	cheats := calculateCheats(maze, costs, reverseCosts, costs[end], 20)

	result := 0
	for _, value := range cheats {
		if value >= 100 {
			result++
		}
	}

	return fmt.Sprint(result)
}
