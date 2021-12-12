package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func getSubsystem(filename string) []string {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	subsystem := make([]string, 0)
	for scanner.Scan() {
		subsystem = helpers.AppendToStringArray(subsystem, scanner.Text())
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return subsystem
}

func matchBracket(open, close string, close_chars map[string]int) int {
	if (open == "(") {
		if (close != ")") { return close_chars[close]}
	} else if (open == "[") {
		if (close != "]") { return close_chars[close]}
	} else if (open == "{") {
		if (close != "}") { return close_chars[close]}
	} else if (open == "<") {
		if (close != ">") { return close_chars[close]}
	}
	return 0
}

func part1(subsystem []string) {
	fmt.Println("====== PART ONE ======")
	close_chars := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	stack, syntax_err_score := make([]string, 0), 0

	for _,line := range subsystem {
		for j := 0; j < len(line); j++ {
			open := string(line[j])
			if !helpers.InStringIntMap(close_chars, open) {
				stack = helpers.AppendToStringArray(stack, open)
			} else {
				close := open
				stack, open = helpers.PopStringArray(stack)
				err_score := matchBracket(open, close, close_chars)
				if (err_score > 0) {
					syntax_err_score += err_score
					break
				}
			}
		}
		stack = make([]string, 0)
	}

	fmt.Println("Syntax Error Score: " + helpers.IntToString(syntax_err_score))
}

func scoreStack(stack []string, close_chars map[string]int) int {
	score, bracket := 0, ""
	for {
		if (len(stack) < 1) {
			break
		}
		score *= 5
		stack, bracket = helpers.PopStringArray(stack)
		if (bracket == "(") {
			score += close_chars[")"]
		} else if (bracket == "[") {
			score += close_chars["]"]
		} else if (bracket == "{") {
			score += close_chars["}"]
		} else if (bracket == "<") {
			score += close_chars[">"]
		}
	}
	return score
}

func part2(subsystem []string) {
	fmt.Println("====== PART TWO ======")
	close_chars := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	stack, stack_scores := make([]string, 0), make([]int, 0)

	for _,line := range subsystem {
		for j := 0; j < len(line); j++ {
			open := string(line[j])
			if !helpers.InStringIntMap(close_chars, open) {
				stack = helpers.AppendToStringArray(stack, open)
			} else {
				close := open
				stack, open = helpers.PopStringArray(stack)
				err_score := matchBracket(open, close, close_chars)
				if (err_score > 0) {
					stack = make([]string, 0)
					break
				}
			}
		}
		if (len(stack) > 0) {
			score := scoreStack(stack, close_chars)
			stack_scores = helpers.AppendToIntArray(stack_scores, score)
		}
		stack = make([]string, 0)
	}

	sort.Ints(stack_scores)
	fmt.Println("Median Autocomplete Score: " + helpers.IntToString(stack_scores[len(stack_scores) / 2]))
}

func main() {
	filename := "puzzle_input.txt"
	subsystem := getSubsystem(filename)
	part1(subsystem)
	part2(subsystem)
}