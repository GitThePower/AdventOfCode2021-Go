package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)


func getPolymer(filename string) (string, map[string]string) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	polymer, rules, switch_parse := "", make(map[string]string), false
	arrow := regexp.MustCompile(" -> ")
	for scanner.Scan() {
		if (!switch_parse) {
			line := scanner.Text()
			if (len(line) == 0) {
				switch_parse = true
			} else {
				polymer = line
			}
		} else {
			rule := arrow.Split(scanner.Text(), -1)
			rules[rule[0]] = rule[1]
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return polymer, rules
}

func part1(polymer string, rules map[string]string) {
	fmt.Println("====== PART ONE ======")
	steps := 10

	for j := 0; j < steps; j++ {
		copy_polymer := "" + polymer
		for k := 1; k < len(polymer); k++ {
			pair := string(polymer[k - 1]) + string(polymer[k])
			if (helpers.InStringStringMap(rules, pair)) {
				insertion := rules[pair]
				copy_polymer = copy_polymer[:k - 1] + insertion + copy_polymer[k - 1:]
			}
		}
		polymer = "" + copy_polymer
	}

	letter_counts := make(map[string]int)
	for j := 0; j < len(polymer); j++ {
		letter := string(polymer[j])
		if (helpers.InStringIntMap(letter_counts, letter)) {
			letter_counts[letter]++
		} else {
			letter_counts[letter] = 1
		}
	}

	high_count, low_count, total := 0, len(polymer), 0
	for _,count := range letter_counts {
		total += count
		if (count > high_count) { high_count = count}
		if (count < low_count) { low_count = count}
	}

	fmt.Println(letter_counts)
	fmt.Println("Diff Between Most and Least Common Elements: " + helpers.IntToString(high_count - low_count))
}

func main() {
	filename := "puzzle_input.txt"
	polymer, rules := getPolymer(filename)
	part1(polymer, rules)
}