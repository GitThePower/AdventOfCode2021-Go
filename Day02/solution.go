package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func printResults(h_pos, depth int) {
	fmt.Println("Horizontal Position: " + helpers.IntToString(h_pos))
	fmt.Println("Depth: " + helpers.IntToString(depth))
	fmt.Println("Product: " + helpers.IntToString(h_pos * depth))
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h_pos, depth := 0, 0
	for scanner.Scan() {
		cmd_prts := strings.Fields(scanner.Text())
		x := helpers.StringToInt(cmd_prts[1])

		if cmd_prts[0] == "forward" { h_pos += x }
		if cmd_prts[0] == "down" { depth += x }
		if cmd_prts[0] == "up" { depth -= x }
	}

	printResults(h_pos, depth)

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func part2(filename string) {
	fmt.Println("====== PART TWO ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h_pos, depth, aim := 0, 0, 0
	for scanner.Scan() {
		cmd_prts := strings.Fields(scanner.Text())
		x := helpers.StringToInt(cmd_prts[1])

		if cmd_prts[0] == "forward" {
			h_pos += x
			depth += aim * x
		}
		if cmd_prts[0] == "down" { aim += x }
		if cmd_prts[0] == "up" { aim -= x }
	}

	printResults(h_pos, depth)

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
	part2(filename)
}