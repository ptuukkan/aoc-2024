package day13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ptuukkan/aoc-2024/utils"
)

type ClawMachine struct {
	ButtonA utils.Point
	ButtonB utils.Point
	Price   utils.Point
}

func parsePoints(input []string) []utils.Point {
	points := make([]utils.Point, 3)
	for i, line := range input {
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(line, -1)
		x, err_x := strconv.Atoi(numbers[0])
		y, err_y := strconv.Atoi(numbers[1])
		if err_x != nil || err_y != nil {
			fmt.Println("error converting numbers")
		}
		points[i] = utils.NewPoint(y, x)
	}

	return points
}

func parseInput(input string) []*ClawMachine {
	var clawMachines []*ClawMachine

	splits := strings.Split(input, "\n\n")
	for _, split := range splits {
		doubleSplit := utils.SplitNewLines(split)
		points := parsePoints(doubleSplit)
		clawMachines = append(clawMachines, &ClawMachine{ButtonA: points[0], ButtonB: points[1], Price: points[2]})
	}

	return clawMachines
}

func Part1(input string) string {
	clawMachines := parseInput(input)

	cost := 0
	for m, clawMachine := range clawMachines {
		fmt.Printf("Machine: %d\n", m)
		curr := utils.NewPoint(0, 0)
		curr_cost := -1
		for i := 1; i < 100; i++ {
			if curr.X > clawMachine.Price.X || curr.Y > clawMachine.Price.Y {
				break
			}
			curr = curr.Add(clawMachine.ButtonA)
			diff_x := clawMachine.Price.X - curr.X
			diff_y := clawMachine.Price.Y - curr.Y
			mod_x := diff_x % clawMachine.ButtonB.X
			mod_y := diff_y % clawMachine.ButtonB.Y
			count_x := diff_x / clawMachine.ButtonB.X
			count_y := diff_y / clawMachine.ButtonB.Y
			if mod_x == 0 && mod_y == 0 && count_x == count_y {
				fmt.Printf("diff_x: %d - diff_y: %d\n", diff_x, diff_y)
				this_cost := i * 3
				b_count := (clawMachine.Price.X - curr.X) / clawMachine.ButtonB.X
				fmt.Printf("a count: %d - b count %d\n", i, b_count)
				this_cost += b_count

				fmt.Printf("cost %d - x: %d - y: %d\n", this_cost, clawMachine.ButtonA.X*i+clawMachine.ButtonB.X*b_count, clawMachine.ButtonA.Y*i+clawMachine.ButtonB.Y*b_count)
				if curr_cost == -1 || this_cost < curr_cost {
					curr_cost = this_cost
				}
			}

		}

		if curr_cost != -1 {
			cost += curr_cost
		}
	}

	return strconv.Itoa(cost)
}

func Part2(input string) string {
	return ""
}
