package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getVents(filename string) ([]int, int, int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	vents := make([]int, 0)
	width := 0
	height := 0
	for scanner.Scan() {
		ventRow := scanner.Text()
		width = len(ventRow)
		height++
		for i := range ventRow {
			vents = helpers.AppendToIntArray(vents, helpers.StringToInt(string(ventRow[i])))
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return vents, width, height
}

func getRiskLevel(vents []int, width, height, idx int) int {
	risk_level := 0
	if ((idx < width || vents[idx] < vents[idx - width]) && 
		(idx % width == 0 || vents[idx] < vents[idx - 1]) && 
		(idx % width == width - 1 || vents[idx] < vents[idx + 1]) && 
		(idx >= width * (height -1) || vents[idx] < vents[idx + width])) {
			risk_level += vents[idx] + 1
	}

	return risk_level
}

func part1(vents []int, width, height int) {
	fmt.Println("====== PART ONE ======")

	sum_risk := 0
	for i := range vents {
		sum_risk += getRiskLevel(vents, width, height, i)
	}

	fmt.Println("Sum of Risk Levels: " + helpers.IntToString(sum_risk))
}

func getBasinEdge(vents []int, width, height, idx int) int {
	queue, marked_vents, basin_edges, pop := make([]int, 0), make(map[int]bool), 0, 0

	marked_vents[idx] = true
	queue = helpers.AppendToIntArray(queue, idx)
	for {
		if len(queue) < 1 {
			break;
		}
		queue,pop = helpers.DequeueIntArray(queue)
		basin_edges++
		if ((pop >= width && vents[pop - width] < 9) && !helpers.InIntBoolMap(marked_vents, pop - width)) {
			marked_vents[pop - width] = true
			queue = helpers.AppendToIntArray(queue, pop - width)
		}
		if ((pop % width > 0 && vents[pop - 1] < 9) && !helpers.InIntBoolMap(marked_vents, pop - 1)) {
			marked_vents[pop - 1] = true
			queue = helpers.AppendToIntArray(queue, pop - 1)
		}
		if ((pop % width < width - 1 && vents[pop + 1] < 9) && !helpers.InIntBoolMap(marked_vents, pop + 1)) {
			marked_vents[pop + 1] = true
			queue = helpers.AppendToIntArray(queue, pop + 1)
		}
		if ((pop < width * (height - 1) && vents[pop + width] < 9) && !helpers.InIntBoolMap(marked_vents, pop + width)) {
			marked_vents[pop + width] = true
			queue = helpers.AppendToIntArray(queue, pop + width)
		}
	}

	return basin_edges
}

func part2(vents []int, width, height int) {
	fmt.Println("====== PART TWO ======")

	largest_basins := make([]int, 3)
	for i := range vents {
		if (getRiskLevel(vents, width, height, i) > 0) {
			basin_size := getBasinEdge(vents, width, height, i)
			min_idx := 0
			for j := range largest_basins {
				if (largest_basins[min_idx] > largest_basins[j]) { min_idx = j }
			}
			if (basin_size > largest_basins[min_idx]) {
				largest_basins[min_idx] = basin_size
			}
		}
	}

	basin_product := 1
	for _,basin := range largest_basins {
		basin_product *= basin
	}

	fmt.Println("Product of Largest Basins: " + helpers.IntToString(basin_product))
}

func main() {
	filename := "puzzle_input.txt"
	vents, width, height := getVents(filename)
	part1(vents, width, height)
	part2(vents, width, height)
}