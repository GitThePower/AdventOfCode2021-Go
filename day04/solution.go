package main

import (
	// "AdventOfCode2021/helpers"
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// List of values chosen.
//	each value contains a list of (x,y) locations
//
// 2-D Array where x = board_idx, y: [0,24] = val_idx

type location struct {
	x int
	y int
}

func parseInput(scanner bufio.Scanner) ([]int, map[int][]location, [][]int) {
	order, chosen, boards := make([]int, 0), make(map[int][]location), make([][]int, 0)

	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	for _, v := range vals {
		val := helpers.StringToInt(v)
		order = helpers.AppendToIntArray(order, val)
		chosen[val] = make([]location, 0)
	}

	board_idx := -1
	val_inc := 0
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		if len(vals) == 5 {
			for i, v := range vals {
				val := helpers.StringToInt(v)
				boards[board_idx][i + val_inc] = val
				chosen[val] = append(chosen[val], []location{{x: board_idx, y: i + val_inc}}...)
			}
			val_inc += len(vals)
		} else {
			boards = helpers.ExtendTwoDIntArray(boards, 1, 25)
			board_idx++
			val_inc = 0
		}
	}

	return order, chosen, boards
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	order, chosen, boards := parseInput(*scanner)

	fmt.Println(order)
	fmt.Println(chosen)
	fmt.Println(boards)

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
}