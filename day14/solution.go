package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)


func getPolymer(filename string) (map[string]int, map[string][]string, map[string]int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	polymer, rules, letter_counts, switch_parse := make(map[string]int), make(map[string][]string), make(map[string]int), false
	arrow := regexp.MustCompile(" -> ")
	for scanner.Scan() {
		if (!switch_parse) {
			line := scanner.Text()
			if (len(line) == 0) {
				switch_parse = true
			} else {
				letter_counts[string(line[0])] = 1
				for k := 1; k < len(line); k++ {
					pair := string(line[k - 1]) + string(line[k])
					letter_counts = helpers.SafeInsertStringIntMap(letter_counts, string(line[k]), 1)
					polymer = helpers.SafeInsertStringIntMap(polymer, pair, 1)
				}
			}
		} else {
			rule := arrow.Split(scanner.Text(), -1)
			rules[rule[0]] = []string{rule[1], string(rule[0][0]) + rule[1], rule[1] + string(rule[0][1])}
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return polymer, rules, letter_counts
}

func diffMostLeastCommonElements(polymer, letter_counts map[string]int, rules map[string][]string, steps int) {
	high_count, low_count := 0, 19 * helpers.Power(2, steps)

	for j := 0; j < steps; j++ {
		copy_polymer := helpers.CopyStringIntMap(polymer)
		for pair := range polymer {
			pair_instances := polymer[pair]
			insertions := rules[pair]
			letter_counts = helpers.SafeInsertStringIntMap(letter_counts, insertions[0], pair_instances)
			copy_polymer = helpers.SafeInsertStringIntMap(copy_polymer, insertions[1], pair_instances)
			copy_polymer = helpers.SafeInsertStringIntMap(copy_polymer, insertions[2], pair_instances)
			copy_polymer[pair] -= pair_instances
		}
		polymer = copy_polymer
	}

	for _,val := range letter_counts {
		if (val > high_count) { high_count = val }
		if (val < low_count) { low_count = val }
	}

	fmt.Println("Diff Between Most and Least Common Elements after: " + helpers.IntToString(high_count - low_count))
}

func main() {
	filename := "puzzle_input.txt"
	polymer, rules, letter_counts := getPolymer(filename)
	fmt.Println("====== PART ONE ======")
	steps := 10
	diffMostLeastCommonElements(helpers.CopyStringIntMap(polymer), helpers.CopyStringIntMap(letter_counts), rules, steps)
	fmt.Println("====== PART TWO ======")
	steps = 40
	diffMostLeastCommonElements(helpers.CopyStringIntMap(polymer), helpers.CopyStringIntMap(letter_counts), rules, steps)
}