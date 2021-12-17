package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getCave(filename string) ([]int, int, int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cave, width, height := make([]int, 0), 0, 0
	for scanner.Scan() {
		caveRow := scanner.Text()
		width = len(caveRow)
		height++
		for i := range caveRow {
			cave = helpers.AppendToIntArray(cave, helpers.StringToInt(string(caveRow[i])))
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return cave, width, height
}

type path struct {
	idx int
	risk int
}

func getNextPaths(p path) []path {
	next_paths := make([]path, 2)
	return next_paths
}

func part1(cave []int, width, height int) {
	fmt.Println("====== PART ONE ======")
	p_queue := []path{{0, 0}}

	for {
		if (len(p_queue) < 1) {
			break
		}
		pop := p_queue[0]
		p_queue = p_queue[1:]
		next_paths := getNextPaths(pop)
		p_queue = append(p_queue, next_paths...)
	}

	fmt.Println(len(cave))
	fmt.Println(width)
	fmt.Println(height)
}

func main() {
	filename := "puzzle_input.txt"
	cave, width, height := getCave(filename)
	part1(cave, width, height)
}