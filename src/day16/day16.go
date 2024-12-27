package day16

import (
	"fmt"
	"slices"
	"strconv"

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

type Vertex struct {
	Position  utils.Point
	Direction utils.Point
	Cost      int
	Prev      *Vertex
}

func newVertex(p utils.Point, d utils.Point, cost int, prev Vertex) Vertex {
	return Vertex{Position: p, Direction: d, Cost: cost, Prev: &prev}
}

func queueAdj(maze []string, vertex Vertex, queue *[]Vertex, visited *[]Vertex) {
	for _, dir := range utils.Directions {
		adj := vertex.Position.Add(dir)
		if maze[adj.Y][adj.X] == '#' {
			continue
		}
		cost := vertex.Cost + 1
		if dir != vertex.Direction {
			cost += 1000
		}

		if slices.ContainsFunc(*queue, func(v Vertex) bool {
			return v.Position == adj && v.Cost <= cost
		}) {
			continue
		}

		if slices.ContainsFunc(*visited, func(v Vertex) bool {
			return v.Position == adj && v.Cost <= cost
		}) {
			continue
		}

		*queue = append(*queue, newVertex(adj, dir, cost, vertex))
	}
}

func findBestPath(maze []string) *Vertex {
	start := findRune(maze, 'S')
	end := findRune(maze, 'E')

	var queue []Vertex
	var visited []Vertex

	vertex := Vertex{Position: start, Direction: utils.Directions[1], Cost: 0, Prev: nil}

	for {
		visited = append(visited, vertex)
		queueAdj(maze, vertex, &queue, &visited)
		if len(queue) == 0 || vertex.Position == end {
			break
		}
		slices.SortFunc(queue, func(a, b Vertex) int {
			return a.Cost - b.Cost
		})

		vertex = queue[0]
		queue = queue[1:]
	}

	return &vertex
}

func Part1(input string) string {
	maze := utils.SplitNewLines(input)

	path := findBestPath(maze)

	print(maze, linkedListToSlice(path))

	return strconv.Itoa(path.Cost)
}

func print(maze []string, trail []utils.Point) {
	for y, line := range maze {
		for x, r := range line {
			if slices.ContainsFunc(trail, func(a utils.Point) bool {
				return a.Y == y && a.X == x
			}) {
				fmt.Printf("O")
			} else {
				fmt.Printf("%s", string(r))
			}

		}
		fmt.Printf("\n")
	}
}

func queueAdjPart2(maze []string, vertex Vertex, queue *[]Vertex, visited *[]Vertex, bestPath map[utils.Point]*Vertex) {
	for i, dir := range utils.Directions {
		adj := vertex.Position.Add(dir)
		if maze[adj.Y][adj.X] == '#' {
			continue
		}

		if vertex.Position.Subtract(&adj) == vertex.Direction {
			continue
		}

		cost := vertex.Cost + 1
		if dir != vertex.Direction {
			cost += 1000
		}

		elem, ok := bestPath[adj]
		if ok && elem.Cost < cost && elem.Direction == dir {
			continue
		}

		if slices.ContainsFunc(*queue, func(v Vertex) bool {
			return v.Position == adj && v.Cost < cost && (v.Direction == dir || v.Direction == utils.Directions[(i+2)%4])
		}) {
			continue
		}

		if slices.ContainsFunc(*visited, func(v Vertex) bool {
			return v.Position == adj && v.Cost < cost && (v.Direction == dir || v.Direction == utils.Directions[(i+2)%4])
		}) {
			continue
		}

		*queue = append(*queue, newVertex(adj, dir, cost, vertex))
	}
}

func vertexToPoint(vertices []Vertex) []utils.Point {
	result := []utils.Point{}

	for _, v := range vertices {
		result = append(result, v.Position)
	}

	return result
}

func findAllBestPaths(maze []string, bestPath map[utils.Point]*Vertex) []Vertex {
	start := findRune(maze, 'S')
	end := findRune(maze, 'E')

	var queue []Vertex
	var visited []Vertex

	vertex := Vertex{Position: start, Direction: utils.Directions[1], Cost: 0}

	for {
		// print(maze, vertexToPoint(visited))
		visited = append(visited, vertex)
		queueAdjPart2(maze, vertex, &queue, &visited, bestPath)
		if len(queue) == 0 || vertex.Position == end {
			break
		}
		slices.SortFunc(queue, func(a, b Vertex) int {
			return a.Cost - b.Cost
		})

		vertex = queue[0]
		queue = queue[1:]
	}

	results := []Vertex{vertex}
	for _, q := range queue {
		if q.Cost == vertex.Cost && q.Position == vertex.Position {
			results = append(results, q)
		}
	}

	return results
}

func Part2(input string) string {
	maze := utils.SplitNewLines(input)
	path := findBestPath(maze)
	trail := make(map[utils.Point]*Vertex)
	for path != nil {
		trail[path.Position] = path
		path = path.Prev
	}

	paths := findAllBestPaths(maze, trail)

	result := make(map[utils.Point]bool)
	for _, p := range paths {
		print(maze, linkedListToSlice(&p))
		v := &p
		for v != nil {
			result[v.Position] = true
			v = v.Prev
		}
	}

	return fmt.Sprint(len(result))
}

func linkedListToSlice(vertex *Vertex) []utils.Point {
	result := []utils.Point{}

	for vertex != nil {
		result = append(result, vertex.Position)
		vertex = vertex.Prev
	}

	return result
}
