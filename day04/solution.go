package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isBingo(pos int, board []int, selected map[int]bool) bool {
	winning_streak := make([]int, 5)
	winning_streak[0] = board[pos]

	bingo := true
	row_start := (pos / 5) * 5
	for w := 1; w < 5; w++ {
		idx := ((pos + w) % 5) + row_start
		val := board[idx]
		if !helpers.InIntBoolMap(selected, val) {
			bingo = false
		} else {
			winning_streak[w] = val
		}
	}
	if bingo {
		// fmt.Println("Winning Streak: " + helpers.IntArrayToString(winning_streak, " ", 5))
		return bingo
	}

	bingo = true
	for h := 1; h < 5; h++ {
		idx := (pos + (h * 5)) % 25
		val := board[idx]
		if !helpers.InIntBoolMap(selected, val) {
			bingo = false
		} else {
			winning_streak[h] = val
		}
	}
	// if bingo { fmt.Println("Winning Streak: " + helpers.IntArrayToString(winning_streak, " ", 5)) }
	return bingo
}

func getResults(board []int, selected map[int]bool, selection_order []int) {
	unselected_total := 0
	for _, v := range board {
		if !helpers.InIntBoolMap(selected, v) {
			unselected_total += v
		}
	}
	last_called_num := selection_order[len(selection_order)-1]
	// fmt.Println("Board: " + helpers.IntArrayToString(board, " ", 5))
	// fmt.Println("Selected In Order: " + helpers.IntArrayToString(selection_order, " ", 5))
	fmt.Println("Sum Of Unselected: " + helpers.IntToString(unselected_total))
	fmt.Println("Last Called Number: " + helpers.IntToString(last_called_num))
	fmt.Println("Product: " + helpers.IntToString(unselected_total*last_called_num))
}

func part1(order []int, chosen map[int][]location, boards [][]int) {
	fmt.Println("====== PART ONE ======")

	selected := make(map[int]bool)
	selection_order := make([]int, 0)
	winning_board := make([]int, 25)
	bingo := false
	for _, v := range order {
		selected[v] = true
		selection_order = helpers.AppendToIntArray(selection_order, v)
		selection := chosen[v]
		for _, loc := range selection {
			c_board := boards[loc.x]
			pos := loc.y
			bingo = isBingo(pos, c_board, selected)
			if bingo {
				winning_board = c_board
				break
			}
		}
		if bingo {
			break
		}
	}

	getResults(winning_board, selected, selection_order)
}

func part2(order []int, chosen map[int][]location, boards [][]int) {
	fmt.Println("====== PART TWO ======")

	selected := make(map[int]bool)
	selection_order := make([]int, 0)
	bingo_boards := make(map[int]bool)
	bingo_count := 0
	losing_board := make([]int, 25)
	for _, v := range order {
		selected[v] = true
		selection_order = helpers.AppendToIntArray(selection_order, v)
		selection := chosen[v]
		for _, loc := range selection {
			c_board := boards[loc.x]
			pos := loc.y
			if !helpers.InIntBoolMap(bingo_boards, loc.x) && isBingo(pos, c_board, selected) {
				bingo_boards[loc.x] = true
				bingo_count++
			}
			if bingo_count == len(boards) {
				losing_board = c_board
				break
			}
		}
		if bingo_count == len(boards) {
			break
		}
	}

	getResults(losing_board, selected, selection_order)
}

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
				boards[board_idx][i+val_inc] = val
				chosen[val] = append(chosen[val], []location{{x: board_idx, y: i + val_inc}}...)
			}
			val_inc += len(vals)
		} else {
			boards = helpers.Extend2DIntArray(boards, 1, 25)
			board_idx++
			val_inc = 0
		}
	}

	return order, chosen, boards
}

func main() {
	f, e := os.Open("puzzle_input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	order, chosen, boards := parseInput(*scanner)
	part1(order, chosen, boards)
	part2(order, chosen, boards)

	if e := scanner.Err(); e != nil {
		log.Fatal(e)
	}
}
