package day13

import (
	"fmt"
	"math"
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
	clawMachines := parseInput(input)

	cost := int64(0)
	for m, clawMachine := range clawMachines {
		fmt.Printf("Clawmachine %d\n", m)
		price_x := int64(clawMachine.Price.X) + 10000000000000
		price_y := int64(clawMachine.Price.Y) + 10000000000000
		a_x := int64(clawMachine.ButtonA.X)
		a_y := int64(clawMachine.ButtonA.Y)
		b_x := int64(clawMachine.ButtonB.X)
		b_y := int64(clawMachine.ButtonB.Y)

		i := int64(0)
		curr_x := i * a_x
		curr_y := i * a_y
		diff_x := price_x - curr_x
		diff_y := price_y - curr_y
		count_x := diff_x / b_x
		count_y := diff_y / b_y

		delta_count_x := int64(-1)
		delta_count_y := int64(-1)
		delta_i := int64(-1)

		first_count_x := int64(-1)
		first_count_y := int64(-1)
		first_i := int64(-1)

		for {
			// if curr_x > price_x || curr_y > price_y || count_y < i {
			// 	fmt.Printf("count_x %d - count_y %d - i %d\n", count_x, count_y, i)
			// 	fmt.Printf("curr_x %d - curr_y %d\n", curr_x, curr_y)
			// 	fmt.Printf("curr_x > price_x: %v", curr_x > price_x)
			// 	fmt.Printf("curr_y > price_y: %v", curr_y > price_y)
			// 	fmt.Printf("count_y < i: %v", count_y < i)
			// 	fmt.Printf("count_x < i: %v", count_x < i)
			// 	break
			// }
			if i > 1000000 {
				break
			}

			curr_x += a_x
			curr_y += a_y
			diff_x = price_x - curr_x
			diff_y = price_y - curr_y
			mod_x := diff_x % b_x
			mod_y := diff_y % b_y
			count_x = diff_x / b_x
			count_y = diff_y / b_y

			if mod_x == 0 && mod_y == 0 {
				if first_i == -1 {
					first_i = i
					first_count_x = count_x
					first_count_y = count_y
					fmt.Printf("First i: %d - first x: %d - first y: %d\n", first_i, first_count_x, first_count_y)
				} else {
					delta_i = i - first_i
					delta_count_x = count_x - first_count_x
					delta_count_y = count_y - first_count_y

					fmt.Printf("delta_i: %d - delta_x: %d - delta_y: %d\n", delta_i, delta_count_x, delta_count_y)

					a1 := -delta_count_x + delta_count_y
					a2 := -first_count_x + first_count_y
					a3 := a2 / a1

					fmt.Printf("a1: %d - a2: %d - a3: %d\n", a1, a2, a3)

					button_a_count := first_i + int64(math.Abs(float64(a3*delta_i)))
					button_a_count++

					fmt.Printf("button_a_count: %d\n", button_a_count)

					curr_x = a_x * button_a_count
					curr_y = a_y * button_a_count

					fmt.Printf("curr_x: %d - curr_y: %d \n", curr_x, curr_y)

					diff_x = price_x - curr_x
					diff_y = price_y - curr_y
					count_x = diff_x / b_x
					count_y = diff_y / b_y

					fmt.Printf("count_x: %d count_y: %d\n", count_x, count_y)

					total_x := curr_x + count_x*b_x
					total_y := curr_y + count_y*b_y

					fmt.Printf("Price_x: %d - Price_y: %d\n", price_x, price_y)
					fmt.Printf("total_x: %d - total_y: %d\n", total_x, total_y)

					if total_x == price_x && total_y == price_y && count_x == count_y {

						fmt.Println("success")

						curr_cost := button_a_count*3 + count_x
						cost += curr_cost
					}

					break
				}
			}
			// if mod_x == 0 && mod_y == 0 && count_y == count_x {
			// 	fmt.Printf("price_x: %d price_y: %d\n", price_x, price_y)
			// 	x := b_x * count_x
			// 	y := b_y * count_y
			// 	fmt.Printf("curr_x: %d curr_y: %d\n", curr_x+x, curr_y+y)
			// 	break
			// }

			i++
		}
	}

	return strconv.FormatInt(cost, 10)
}
