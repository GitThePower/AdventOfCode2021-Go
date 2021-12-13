package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)


func getDumbos(filename string) ([]int, int, int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dumbos := make([]int, 0)
	width := 0
	height := 0
	for scanner.Scan() {
		dumboRow := scanner.Text()
		width = len(dumboRow)
		height++
		for i := range dumboRow {
			dumbos = helpers.AppendToIntArray(dumbos, helpers.StringToInt(string(dumboRow[i])))
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return dumbos, width, height
}

func flashNeighbors(idx, width, height int, dumbos []int) []int {
	dumbos[idx] = 0

	if (idx >= width + 1 && idx % width > 0 && dumbos[idx - width - 1] > 0) {
		dumbos[idx - width - 1]++
	}
	if (idx >= width && dumbos[idx - width] > 0) {
		dumbos[idx - width]++
	}
	if (idx >= width - 1 && idx % width < width - 1 && dumbos[idx - width + 1] > 0) {
		dumbos[idx - width + 1]++
	}
	if (idx % width > 0 && dumbos[idx - 1] > 0) {
		dumbos[idx - 1]++
	}
	if (idx % width < width - 1 && dumbos[idx + 1] > 0) {
		dumbos[idx + 1]++
	}
	if (idx < (width * (height - 1)) + 1 && idx % width > 0 && dumbos[idx + width - 1] > 0) {
		dumbos[idx + width - 1]++
	}
	if (idx < width * (height - 1) && dumbos[idx + width] > 0) {
		dumbos[idx + width]++
	}
	if (idx < (width * (height - 1)) - 1 && idx % width < width - 1 && dumbos[idx + width + 1] > 0) {
		dumbos[idx + width + 1]++
	}
	return dumbos
}

func flashStep(dumbos []int, width, height int) ([]int, int) {
	flashes := 0
	for i := range dumbos {
		dumbos[i]++
	}

	for {
		res_dumbos, dumbo_change := helpers.CopyIntArray(dumbos), false
		for i,dumbo := range dumbos {
			if (dumbo > 9) {
				dumbo_change = true
				flashes += 1
				res_dumbos = flashNeighbors(i, width, height, res_dumbos)
			}
		}
		if (!dumbo_change) { break }
		dumbos = res_dumbos
	}

	return dumbos, flashes
}

func part1(dumbos []int, width, height int) {
	fmt.Println("====== PART ONE ======")
	steps, sum_flash := 100, 0

	for j := 0; j < steps; j++ {
		flashes := 0
		dumbos, flashes = flashStep(dumbos, width, height)
		sum_flash += flashes
		// fmt.Println("After Step " + helpers.IntToString(j + 1) + ":\n" + helpers.IntArrayToString(dumbos, "", width))
	}

	fmt.Println("Total Flashes: " + helpers.IntToString(sum_flash))
}

func part2(dumbos []int, width, height int) {
	fmt.Println("====== PART TWO ======")
	flash_step := 0

	for {
		flash_step++
		flashes := 0
		dumbos, flashes = flashStep(dumbos, width, height)
		// fmt.Println("After Step " + helpers.IntToString(flash_step) + ":\n" + helpers.IntArrayToString(dumbos, "", width))
		if (flashes == len(dumbos)) { break }
	}

	fmt.Println("Simultaneous Flash Occurs on Step " + helpers.IntToString(flash_step))
}

func main() {
	filename := "puzzle_input.txt"
	dumbos, width, height := getDumbos(filename)
	part1(helpers.CopyIntArray(dumbos), width, height)
	part2(helpers.CopyIntArray(dumbos), width, height)
}