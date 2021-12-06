package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	some_val := 0
	for scanner.Scan() {
		some_val = helpers.StringToInt(scanner.Text())

	}

	fmt.Println("Solution: " + helpers.IntToString(some_val))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
	// part2(filename)
}
