package day12

import (
	"slices"
	"strconv"

	"github.com/ptuukkan/aoc-2024/utils"
)

type Region struct {
	Plant rune
	Plots map[utils.Point]*Plot
}

func newRegion(plant rune) *Region {
	plots := make(map[utils.Point]*Plot)

	return &Region{Plant: plant, Plots: plots}
}

func (r *Region) hasPlot(plot utils.Point) bool {
	for key := range r.Plots {
		if key == plot {
			return true
		}
	}
	return false
}

type Plot struct {
	Fences map[int]bool
}

func (p *Plot) fencesCount() int {
	count := 0
	for _, v := range p.Fences {
		if v {
			count++
		}
	}
	return count
}

func newPlot(fences map[int]bool) *Plot {
	return &Plot{Fences: fences}
}

func shoudFence(garden []string, plant rune, plot utils.Point) bool {
	if plot.OutOfBounds(len(garden)) {
		return true
	}
	if rune(garden[plot.Y][plot.X]) != plant {
		return true
	}
	return false
}

func countFences(garden []string, plant rune, plot utils.Point) map[int]bool {
	fences := make(map[int]bool)
	fences[0] = shoudFence(garden, plant, plot.Up())
	fences[1] = shoudFence(garden, plant, plot.Right())
	fences[2] = shoudFence(garden, plant, plot.Down())
	fences[3] = shoudFence(garden, plant, plot.Left())
	return fences
}

func isInRegion(regions []*Region, plot utils.Point) bool {
	for _, region := range regions {
		if region.hasPlot(plot) {
			return true
		}
	}
	return false
}

func getRegionPlots(garden []string, regionPlots map[utils.Point]bool, plant rune, plot utils.Point) {
	if plot.OutOfBounds(len(garden)) {
		return
	}
	if rune(garden[plot.Y][plot.X]) != plant {
		return
	}
	if regionPlots[plot] {
		return
	}
	regionPlots[plot] = true
	getRegionPlots(garden, regionPlots, plant, plot.Up())
	getRegionPlots(garden, regionPlots, plant, plot.Right())
	getRegionPlots(garden, regionPlots, plant, plot.Down())
	getRegionPlots(garden, regionPlots, plant, plot.Left())
}

func populateRegion(region *Region, garden []string, startingPlot utils.Point) {
	regionPlots := make(map[utils.Point]bool)
	getRegionPlots(garden, regionPlots, region.Plant, startingPlot)

	for regionPlot := range regionPlots {
		fences := countFences(garden, region.Plant, regionPlot)
		plot := newPlot(fences)
		region.Plots[regionPlot] = plot
	}
}

func Part1(input string) string {
	garden := utils.SplitNewLines(input)

	var regions []*Region

	for y, line := range garden {
		for x := range line {
			plot := utils.NewPoint(y, x)
			if isInRegion(regions, plot) {
				continue
			}
			region := newRegion(rune(garden[y][x]))
			regions = append(regions, region)
			populateRegion(region, garden, plot)
		}
	}

	fenceCost := 0
	for _, region := range regions {
		area := len(region.Plots)
		fences := 0
		for _, plot := range region.Plots {
			fences += plot.fencesCount()
		}
		fenceCost += area * fences

	}
	return strconv.Itoa(fenceCost)
}

func groupPlots(plots map[utils.Point]*Plot, length int) [][][]int {
	groups := make([][][]int, 4)

	for groupId := 0; groupId < 4; groupId++ {
		for i := 0; i < length; i++ {
			var plotsWithFences []int
			for plotPos, plot := range plots {
				mainAxis := plotPos.X
				crossAxis := plotPos.Y
				if groupId%2 == 0 {
					mainAxis = plotPos.Y
					crossAxis = plotPos.X
				}
				if mainAxis == i && plot.Fences[groupId] {
					plotsWithFences = append(plotsWithFences, crossAxis)
				}

			}
			if len(plotsWithFences) > 0 {
				groups[groupId] = append(groups[groupId], plotsWithFences)
			}
		}
	}

	return groups
}

func Part2(input string) string {
	garden := utils.SplitNewLines(input)

	var regions []*Region

	for y, line := range garden {
		for x := range line {
			plot := utils.NewPoint(y, x)
			if isInRegion(regions, plot) {
				continue
			}
			region := newRegion(rune(garden[y][x]))
			regions = append(regions, region)
			populateRegion(region, garden, plot)
		}
	}

	fenceCost := 0
	for _, region := range regions {
		area := len(region.Plots)
		fences := 0
		groups := groupPlots(region.Plots, len(garden))

		for _, group := range groups {
			asd := discount(group)
			fences += asd
		}

		fenceCost += area * fences

	}
	return strconv.Itoa(fenceCost)
}

func discount(groups [][]int) int {
	fences := 0

	for _, group := range groups {
		slices.Sort(group)
		fences++
		for i := 1; i < len(group); i++ {
			if group[i]-group[i-1] > 1 {
				fences++
			}
		}
	}

	return fences
}
