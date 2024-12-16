package day16

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Vertex struct {
	Key   utils.Point
	Edges map[*Vertex]int
}

func newVertex(key utils.Point) *Vertex {
	return &Vertex{Key: key, Edges: make(map[*Vertex]int)}
}

type Graph struct {
	Vertices map[utils.Point]*Vertex
}

func (g *Graph) dijkstra(startKey utils.Point) (map[utils.Point]int, error) {
	if _, ok := g.Vertices[startKey]; !ok {
		return nil, fmt.Errorf("start vertex %v not found", startKey)
	}

	distances := make(map[utils.Point]int)
	for key := range g.Vertices {
		distances[key] = math.MaxInt32
	}
	distances[startKey] = 0

	var vertices []*Vertex
	for _, vertex := range g.Vertices {
		vertices = append(vertices, vertex)
	}

	for len(vertices) != 0 {
		sort.SliceStable(vertices, func(i, j int) bool {
			return distances[vertices[i].Key] < distances[vertices[j].Key]
		})

		vertex := vertices[0]
		vertices = vertices[1:]

		for adjacent, cost := range vertex.Edges {
			alt := distances[vertex.Key] + cost
			if alt < distances[adjacent.Key] {
				distances[adjacent.Key] = alt
			}
		}
	}

	return distances, nil
}

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

func findEdges(maze []string, pos utils.Point, dir utils.Point) []utils.Point {
	var edges []utils.Point

	for _, d := range utils.Directions {
		p := pos.Add(d)
		if pos.Subtract(&p) == dir {
			continue
		}
		if maze[p.Y][p.X] != '#' {
			edges = append(edges, p)
		}
	}

	return edges
}

func createVertices(maze []string, graph *Graph, pos utils.Point, dir utils.Point) {
	fmt.Printf("Y: %d - X: %d\n", pos.Y, pos.X)
	// time.Sleep(1000 * time.Millisecond)
	var vertex *Vertex
	if elem, ok := graph.Vertices[pos]; ok {
		vertex = elem
	} else {
		vertex = newVertex(pos)
		graph.Vertices[pos] = vertex
	}

	edges := findEdges(maze, pos, dir)
	fmt.Printf("Possible Edges: %v\n", edges)
	for _, edge := range edges {
		var edgeVertex *Vertex
		cost := 1
		edgeDirection := edge.Subtract(&pos)
		if edgeDirection != dir {
			cost = 1001
		}

		if elem, ok := graph.Vertices[edge]; ok {
			fmt.Printf("Vertex %v already exists\n", edge)
			edgeVertex = elem
		} else {
			fmt.Printf("Vertex %v does not exist\n", edge)
			edgeVertex = newVertex(edge)
			graph.Vertices[edge] = edgeVertex
			createVertices(maze, graph, edge, edgeDirection)
		}

		if _, ok := vertex.Edges[edgeVertex]; !ok {
			fmt.Printf("New edge: %v\n", edge)
			fmt.Printf("Pos %v - Edge %v - Edge dir %v - dir %v - cost %d\n", pos, edge, edgeDirection, dir, cost)
			vertex.Edges[edgeVertex] = cost
		}
	}
}

func createGraph(input string) (*Graph, utils.Point, utils.Point) {
	maze := utils.SplitNewLines(input)
	graph := &Graph{Vertices: make(map[utils.Point]*Vertex)}
	startPos := findRune(maze, 'S')
	endPos := findRune(maze, 'E')

	createVertices(maze, graph, startPos, utils.Directions[1])

	return graph, startPos, endPos
}

func Part1(input string) string {
	graph, start, end := createGraph(input)

	distances, err := graph.dijkstra(start)
	if err != nil {
		return err.Error()
	}

	endCost := distances[end]

	var points []utils.Point
	for key := range distances {
		points = append(points, key)
		// fmt.Printf("y: %d - x: %d - cost: %d\n", key.Y, key.X, value)
	}
	slices.SortFunc(points, func(i, j utils.Point) int {
		return distances[i] - distances[j]
	})

	for _, p := range points {

		fmt.Printf("y: %d - x: %d - cost: %d\n", p.Y, p.X, distances[p])
	}

	return strconv.Itoa(endCost)
}


func Part2(input string) string {
	return ""
}
