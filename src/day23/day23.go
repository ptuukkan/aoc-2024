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

func intersect(maps []map[string]bool) string {
	freq := make(map[string]int)

	for _, m := range maps {
		for key := range m {
			freq[key]++
		}
	}

	count := make(map[int]int)

	for _, value := range freq {
		count[value]++
	}

	scores := make(map[int]int)

	for key, value := range count {
		scores[key] = key * value
	}

	cutOff := 0
	cutOffScore := 0
	for key, value := range scores {
		if value > cutOffScore {
			cutOffScore = value
			cutOff = key
		}
	}

	chosen := ""
	for key, value := range freq {
		if value >= cutOff {
			chosen = key
		}
	}

	filteredMaps := slices.DeleteFunc(maps, func(a map[string]bool) bool {
		for key := range a {
			if key == chosen {
				return false
			}
		}
		return true
	})

	newFreq := make(map[string]int)

	for _, m := range filteredMaps {
		for key := range m {
			newFreq[key]++
		}
	}
	result := make(map[string]bool)
	for key, count := range newFreq {
		if count == len(filteredMaps) {
			result[key] = true
		}
	}

	keys := []string{}
	for key := range result {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	return strings.Join(keys, ",")
}

func Part2(input string) string {
	nodes := parseInput(input)

	groups := make(map[string]int)

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

		group := intersect(nextConnections)
		groups[group]++
	}

	result := ""
	maxCount := 0
	for key, value := range groups {
		if value > maxCount {
			result = key
			maxCount = value
		}

	}

	return result
}
