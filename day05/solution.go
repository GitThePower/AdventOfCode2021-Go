package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type location struct {
	x int
	y int
}

func expandOceanFloor(start, end location, oceanFloor [][]int) [][]int {
	greatestX := start.x
	if end.x > start.x { greatestX = end.x }
	greatestY := start.y
	if end.y > start.y { greatestY = end.y }

	if greatestX > len(oceanFloor) { oceanFloor = make([][]int, greatestX) }
	if greatestY > len(oceanFloor[0]) {
		for i := range oceanFloor {
			oceanFloor[i] = make([]int, greatestY)
		}
	}

	return oceanFloor
}

func getVents(start, end location, ventList []location, diags bool) []location {
	if start.x == end.x {
		if start.y <= end.y {
			ventList = append(ventList, []location{start}...)
			for j := start.y + 1; j < end.y; j++ {
				ventList = append(ventList, []location{{x:start.x, y:j}}...)
			}
			ventList = append(ventList, []location{end}...)
		} else {
			ventList = append(ventList, []location{end}...)
			for j := end.y + 1; j < start.y; j++ {
				ventList = append(ventList, []location{{x:start.x, y:j}}...)
			}
			ventList = append(ventList, []location{start}...)
		}
	} else if start.y == end.y {
		if start.x <= end.x {
			ventList = append(ventList, []location{start}...)
			for j := start.x + 1; j < end.x; j++ {
				ventList = append(ventList, []location{{x:j, y:start.y}}...)
			}
			ventList = append(ventList, []location{end}...)
		} else {
			ventList = append(ventList, []location{end}...)
			for j := end.x + 1; j < start.x; j++ {
				ventList = append(ventList, []location{{x:j, y:start.y}}...)
			}
			ventList = append(ventList, []location{start}...)
		}
	} else if diags {
		if start.x <= end.x {
			if start.y <= end.y {
				ventList = append(ventList, []location{start}...)
				for j := 1; j < end.y - start.y; j++ {
					ventList = append(ventList, []location{{x:start.x + j, y:start.y + j}}...)
				}
				ventList = append(ventList, []location{end}...)
			} else {
				ventList = append(ventList, []location{start}...)
				for j := 1; j < start.y - end.y; j++ {
					ventList = append(ventList, []location{{x:start.x + j, y:start.y - j}}...)
				}
				ventList = append(ventList, []location{end}...)
			}
		} else {
			if start.y <= end.y {
				ventList = append(ventList, []location{end}...)
				for j := 1; j < end.y - start.y; j++ {
					ventList = append(ventList, []location{{x:end.x + j, y:end.y - j}}...)
				}
				ventList = append(ventList, []location{start}...)
			} else {
				ventList = append(ventList, []location{end}...)
				for j := 1; j < start.y - end.y; j++ {
					ventList = append(ventList, []location{{x:end.x + j, y:end.y + j}}...)
				}
				ventList = append(ventList, []location{start}...)
			}
		}
	}
	
	return ventList
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ventList, oceanFloor := make([]location, 0), make([][]int, 0)
	for scanner.Scan() {
		arrow := regexp.MustCompile(" -> ")
		rng := arrow.Split(scanner.Text(), -1)
		start := location{x: helpers.StringToInt(strings.Split(rng[0], ",")[0]), y: helpers.StringToInt(strings.Split(rng[0], ",")[1])}
		end := location{x: helpers.StringToInt(strings.Split(rng[1], ",")[0]), y: helpers.StringToInt(strings.Split(rng[1], ",")[1])}
		oceanFloor = expandOceanFloor(start, end, oceanFloor)
		ventList = getVents(start, end, ventList, false)
	}

	for _, vent := range ventList {
		oceanFloor[vent.x - 1][vent.y - 1]++
	}

	overLappingVentCount := 0
	for i := range oceanFloor {
		for _, vents := range oceanFloor[i] {
			if vents > 1 { overLappingVentCount++ }
		}
	}

	fmt.Println("Overlapping Vents: " + helpers.IntToString(overLappingVentCount))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func part2(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	ventList, oceanFloor := make([]location, 0), make([][]int, 0)
	for scanner.Scan() {
		arrow := regexp.MustCompile(" -> ")
		rng := arrow.Split(scanner.Text(), -1)
		start := location{x: helpers.StringToInt(strings.Split(rng[0], ",")[0]), y: helpers.StringToInt(strings.Split(rng[0], ",")[1])}
		end := location{x: helpers.StringToInt(strings.Split(rng[1], ",")[0]), y: helpers.StringToInt(strings.Split(rng[1], ",")[1])}
		oceanFloor = expandOceanFloor(start, end, oceanFloor)
		ventList = getVents(start, end, ventList, true)
	}

	for _, vent := range ventList {
		oceanFloor[vent.x - 1][vent.y - 1]++
	}

	overLappingVentCount := 0
	for i := range oceanFloor {
		for _, vents := range oceanFloor[i] {
			if vents > 1 { overLappingVentCount++ }
		}
	}

	fmt.Println("Overlapping Vents: " + helpers.IntToString(overLappingVentCount))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
	part2(filename)
}