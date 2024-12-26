package day23

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Node struct {
	Name  string
	Edges map[string]*Node
}

func getOrCreate(nodes map[string]*Node, name string) *Node {
	if elem, ok := nodes[name]; ok {
		return elem
	}

	node := &Node{Name: name, Edges: make(map[string]*Node)}
	nodes[name] = node
	return node
}

func addConnections(node_a, node_b *Node) {
	node_a.Edges[node_b.Name] = node_b
	node_b.Edges[node_a.Name] = node_a
}

func parseInput(input string) map[string]*Node {
	lines := utils.SplitNewLines(input)
	nodes := make(map[string]*Node)
	for _, line := range lines {
		split := strings.Split(line, "-")
		node_a := getOrCreate(nodes, split[0])
		node_b := getOrCreate(nodes, split[1])
		addConnections(node_a, node_b)
	}

	return nodes
}

func getConnections(node *Node) []string {
	var connections []string
	for key := range node.Edges {
		connections = append(connections, key)
	}

	return connections
}

func getGroups(nodes map[string]*Node) map[string]bool {
	groups := make(map[string]bool)
	for _, node := range nodes {
		connections := getConnections(node)
		for i := 0; i < len(connections); i++ {
			for j := i + 1; j < len(connections); j++ {
				a := node.Edges[connections[i]]
				b := node.Edges[connections[j]]
				_, ok_a := a.Edges[b.Name]
				_, ok_b := b.Edges[a.Name]
				if ok_a && ok_b {
					names := []string{node.Name, connections[i], connections[j]}
					sort.Strings(names)
					key := strings.Join(names, ",")
					groups[key] = true
				}
			}
		}
	}
	return groups
}

func Part1(input string) string {
	nodes := parseInput(input)
	groups := getGroups(nodes)

	result := 0
	for key := range groups {
		if strings.HasPrefix(key, "t") {
			result++
		} else if strings.Contains(key, ",t") {
			result++
		}
	}

	return fmt.Sprint(result)
}

var asd map[string]int

func intersect(maps []map[string]bool) int {
	freq := make(map[string]int)

	for _, m := range maps {
		for key := range m {
			freq[key]++
			asd[key]++
		}
	}

	keys := []string{}
	for key := range freq {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return freq[b] - freq[a]
	})

	for _, s := range keys {
		fmt.Printf("%s - %d\n", s, freq[s])
	}

	// results := make([]int, len(maps))
	//
	// for i := 0; i < len(maps); i++ {
	// 	for _, count := range freq {
	// 		results[count-1]++
	// 	}
	// }

	// result := make(map[string]bool)
	// for key, count := range freq {
	// 	if count == len(maps) {
	// 		result[key] = true
	// 	}
	// }

	return 0
}

func Part2(input string) string {
	nodes := parseInput(input)
    asd = make(map[string]int)
	for _, node := range nodes {
		connections := getConnections(node)
		nextConnections := make([]map[string]bool, len(connections))
		for i, conn := range connections {
			nextConnections[i] = make(map[string]bool)
			conns := getConnections(node.Edges[conn])
			for _, c := range conns {
				nextConnections[i][c] = true
			}
			nextConnections[i][conn] = true
		}

		for _, n := range nextConnections {
			fmt.Println(n)
		}

		common := intersect(nextConnections)

		fmt.Printf("Node: %s - len: %d\n", node.Name, common)

	}

	keys := []string{}
	for key := range asd {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return asd[b] - asd[a]
	})

	for _, s := range keys {
		fmt.Printf("%s - %d\n", s, asd[s])
	}

	return ""
}
