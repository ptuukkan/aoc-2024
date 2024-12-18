package day18

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

var length = 71
var limit = 1024

type Vertex struct {
	Position  utils.Point
	Direction utils.Point
	Cost      int
	Trail     []utils.Point
}

func newVertex(p utils.Point, d utils.Point, cost int, trail []utils.Point) Vertex {
	return Vertex{Position: p, Direction: d, Cost: cost, Trail: trail}
}

func queueAdj(bytes []utils.Point, vertex Vertex, queue *[]Vertex, visited *[]Vertex) {
	for _, dir := range utils.Directions {
		adj := vertex.Position.Add(dir)
		if adj.OutOfBounds(length) {
			continue
		}
		if slices.Contains(bytes, adj) {
			continue
		}
		cost := vertex.Cost + 1

		if slices.ContainsFunc(*queue, func(v Vertex) bool {
			return v.Position == adj && v.Cost <= cost
		}) {
			continue
		}

		if slices.ContainsFunc(*visited, func(v Vertex) bool {
			return v.Position == adj
		}) {
			continue
		}

		*queue = append(*queue, newVertex(adj, dir, cost, append(vertex.Trail, vertex.Position)))
	}

}

func parseInput(input string) []utils.Point {
	lines := utils.SplitNewLines(input)
	bytes := make([]utils.Point, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		num_a, err_a := strconv.Atoi(split[0])
		num_b, err_b := strconv.Atoi(split[1])
		if err_a != nil || err_b != nil {
			fmt.Println("error convering numbers")
		}
		bytes[i] = utils.NewPoint(num_b, num_a)
	}
	return bytes
}

func dijkstra(bytes []utils.Point) (bool, Vertex) {
	start := utils.NewPoint(0, 0)
	end := utils.NewPoint(length-1, length-1)

	var queue []Vertex
	var visited []Vertex

	vertex := Vertex{Position: start, Direction: utils.Directions[1], Cost: 0}

	for {
		visited = append(visited, vertex)
		queueAdj(bytes, vertex, &queue, &visited)
		if len(queue) == 0 || vertex.Position == end {
			break
		}
		slices.SortFunc(queue, func(a, b Vertex) int {
			return a.Cost - b.Cost
		})

		vertex = queue[0]
		queue = queue[1:]
	}

	if vertex.Position == end {
		return true, vertex
	}
	return false, vertex
}

func Part1(input string) string {
	bytes := parseInput(input)
	bytes = bytes[:limit]

	_, vertex := dijkstra(bytes)

	return strconv.Itoa(vertex.Cost)
}

func Part2(input string) string {
	bytes := parseInput(input)

	currentLimit := limit
	for {
		ok, vertex := dijkstra(bytes[:currentLimit])
		if !ok {
			return fmt.Sprintf("%d,%d", bytes[currentLimit-1].X, bytes[currentLimit-1].Y)
		}

		for i, b := range bytes[currentLimit:] {
			if slices.Contains(vertex.Trail, b) {
				currentLimit += i + 1
				break
			}
		}

	}
}
