package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h_pos := 0
	depth := 0
	for scanner.Scan() {
		cmd_prts := strings.Fields(scanner.Text())
		i, e := strconv.Atoi(cmd_prts[1])
		if e != nil { log.Fatal(e) }

		if cmd_prts[0] == "forward" { h_pos += i }
		if cmd_prts[0] == "down" { depth += i }
		if cmd_prts[0] == "up" { depth -= i }
	}

	fmt.Println("Horizontal Position: " + strconv.Itoa(h_pos))
	fmt.Println("Depth: " + strconv.Itoa(depth))
	fmt.Println("Product: " + strconv.Itoa(h_pos * depth))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func part2(filename string) {
	fmt.Println("====== PART TWO ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	h_pos := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		cmd_prts := strings.Fields(scanner.Text())
		i, e := strconv.Atoi(cmd_prts[1])
		if e != nil { log.Fatal(e) }

		if cmd_prts[0] == "forward" {
			h_pos += i
			depth += aim * i
		}
		if cmd_prts[0] == "down" { aim += i }
		if cmd_prts[0] == "up" { aim -= i }
	}

	fmt.Println("Horizontal Position: " + strconv.Itoa(h_pos))
	fmt.Println("Depth: " + strconv.Itoa(depth))
	fmt.Println("Product: " + strconv.Itoa(h_pos * depth))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
	part2(filename)
}