package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func expandPaper(x, y int, paper [][]int) [][]int {
	if (x > len(paper) && (len(paper) == 0 || y > len(paper[0]))) {
		paper = helpers.Resize2DIntArray(paper, x, y)
	} else if (x > len(paper)) {
		paper = helpers.Resize2DIntArray(paper, x, len(paper[0]))
	} else if (len(paper) == 0 || y > len(paper[0])) {
		paper = helpers.Resize2DIntArray(paper, len(paper), y)
	}
	return paper
}

func getPaper(filename string) ([][]int, []int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	paper, folds, switch_parse := make([][]int, 0), make([]int, 0), false

	for scanner.Scan() {

		if (!switch_parse) {
			line := scanner.Text()
			if (len(line) == 0) {
				switch_parse = true
			} else {
				coords := strings.Split(line, ",")
				x, y := helpers.StringToInt(coords[0]), helpers.StringToInt(coords[1])
				paper = expandPaper(x + 1, y + 1, paper)
				paper[x][y] = 1
			}
		} else {
			op := strings.Split(strings.Fields(scanner.Text())[2], "=")
			fold := helpers.StringToInt(op[1])
			if (op[0] == "x") { fold *= -1 }
			folds = helpers.AppendToIntArray(folds, fold)
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return paper, folds
}

func foldPaper(paper [][]int, fold int) [][]int {
	folded_paper, fold_idx, count := helpers.Copy2DIntArray(paper), fold, 0
	if (fold_idx > 0) {
		if (fold_idx > len(paper[0]) / 2) {
			for x := 0; x < len(paper); x++ {
				folded_paper[x] = folded_paper[x][:fold_idx]
				for y := 1; y < len(paper[x]) - fold_idx; y++ {
					tof_fold_y, bof_fold_y := fold_idx - y, fold_idx + y
					if (folded_paper[x][tof_fold_y] == 0 && paper[x][bof_fold_y] == 1) {
						folded_paper[x][tof_fold_y] = 1
					} else if (paper[x][bof_fold_y] == 1) {
						count++
					}
				}
				count += paper[x][fold_idx]
			}
		} else {
			for x := 0; x < len(paper); x++ {
				folded_paper[x] = folded_paper[x][fold_idx + 1:]
				for y := 1; y < fold_idx + 1; y++ {
					tof_fold_y, bof_fold_y := fold_idx - y, y - 1
					if (folded_paper[x][bof_fold_y] == 0 && paper[x][tof_fold_y] == 1) {
						folded_paper[x][bof_fold_y] = 1
					} else if (paper[x][tof_fold_y] == 1) {
						count++
					}
				}
				count += paper[x][fold_idx]
			}
		}
	} else {
		fold_idx *= -1
		if (fold_idx > len(paper) / 2) {
			folded_paper = folded_paper[:fold_idx]
			for x := 1; x < len(paper) - fold_idx; x++ {
				lof_fold_x, rof_fold_x := fold_idx - x, fold_idx + x
				for y := 0; y < len(paper[lof_fold_x]); y++ {
					if (folded_paper[lof_fold_x][y] == 0 && paper[rof_fold_x][y] == 1) {
						folded_paper[lof_fold_x][y] = 1
					} else if (paper[rof_fold_x][y] == 1) {
						count++
					}
				}
			}
		} else {
			folded_paper = folded_paper[fold_idx + 1:]
			for x := 1; x < fold_idx + 1; x++ {
				lof_fold_x, rof_fold_x := fold_idx - x, x - 1
				for y := 0; y < len(paper[rof_fold_x]); y++ {
					if (folded_paper[rof_fold_x][y] == 0 && paper[lof_fold_x][y] == 1) {
						folded_paper[rof_fold_x][y] = 1
					} else if (paper[lof_fold_x][y] == 1) {
						count++
					}
				}
			}
		}
		for _,dot := range paper[fold_idx] {
			count += dot
		}
	}
	fmt.Println("Overlapping Dot Count: " + helpers.IntToString(count))
	return folded_paper
}

func part1(paper [][]int, folds []int) {
	fmt.Println("====== PART ONE ======")

	fold := folds[0]
	paper = foldPaper(paper, fold)

	sum_dots := 0
	for x := 0; x < len(paper); x++ {
		for y := 0; y < len(paper[x]); y++ {
			sum_dots += paper[x][y]
		}
	} 
	
	fmt.Println("Width of Paper: " + helpers.IntToString(len(paper)))
	fmt.Println("Height of Paper: " + helpers.IntToString(len(paper[0])))
	fmt.Println("Sum of Dots on Paper: " + helpers.IntToString(sum_dots))
}

func part2(paper [][]int, folds []int) {
	fmt.Println("====== PART TWO ======")

	for _,fold := range folds {
		paper = foldPaper(paper, fold)
	}

	for x := 0; x < len(paper); x++ {
		for y := 0; y < len(paper[x]); y++ {
			if (paper[x][y] == 0) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	filename := "puzzle_input.txt"
	paper, folds := getPaper(filename)
	fmt.Println("Starting Width of Paper: " + helpers.IntToString(len(paper)))
	fmt.Println("Starting Height of Paper: " + helpers.IntToString(len(paper[0])))
	part1(paper, folds)
	paper, folds = getPaper(filename)
	part2(paper, folds)
}