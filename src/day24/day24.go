package day24

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Gate struct {
	InputA, InputB, Output string
	Type                   string
}

func addWire(wires map[string]int, wire string, value int) {
	if _, ok := wires[wire]; !ok {
		wires[wire] = value
	}
}

func parseInput(input string) (map[string]int, []Gate) {
	wires := make(map[string]int)
	gates := []Gate{}

	split := strings.Split(input, "\n\n")

	wireInputs := utils.SplitNewLines(split[0])
	for _, wi := range wireInputs {
		spl := strings.Split(wi, ": ")
		num, err := strconv.Atoi(spl[1])
		if err != nil {
			fmt.Println("error")
		}
		addWire(wires, spl[0], num)
	}

	gateInputs := utils.SplitNewLines(split[1])
	for _, gi := range gateInputs {
		fields := strings.Fields(gi)
		gate := Gate{InputA: fields[0], InputB: fields[2], Output: fields[4], Type: fields[1]}
		gates = append(gates, gate)
		addWire(wires, fields[0], -1)
		addWire(wires, fields[2], -1)
		addWire(wires, fields[4], -1)
	}

	return wires, gates
}

func calculateResult(wires map[string]int) int {
	keys := []string{}
	for key := range wires {
		if strings.HasPrefix(key, "z") {
			keys = append(keys, key)
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	result := 0
	for _, k := range keys {
        result = result << 1
		result = result | wires[k]
	}

	return result
}

func Part1(input string) string {
	wires, gates := parseInput(input)

	for {
		nextGateIndex := slices.IndexFunc(gates, func(gate Gate) bool {
			return wires[gate.InputA] != -1 && wires[gate.InputB] != -1 && wires[gate.Output] == -1
		})

		if nextGateIndex == -1 {
			break
		}
		gate := gates[nextGateIndex]

		switch gate.Type {
		case "AND":
			wires[gate.Output] = wires[gate.InputA] & wires[gate.InputB]
		case "OR":
			wires[gate.Output] = wires[gate.InputA] | wires[gate.InputB]
		case "XOR":
			wires[gate.Output] = wires[gate.InputA] ^ wires[gate.InputB]
		}
	}
	result := calculateResult(wires)

	return fmt.Sprint(result)
}

func Part2(input string) string {
	return ""
}
