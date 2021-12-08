package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getSwarm(filename string) []int {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s_swarm := strings.Split(scanner.Text(), ",")
	n_swarm := make([]int, len(s_swarm))
	for i,crab := range s_swarm {
		n_swarm[i] = helpers.StringToInt(crab)
	}
	

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return n_swarm
}

func part1(swarm []int) {
	fmt.Println("====== PART ONE ======")
	sort.Ints(swarm)
	median_pos := swarm[len(swarm) / 2]

	fuel_total := 0
	for _,pos := range swarm {
		fuel_total += helpers.Abs(pos - median_pos)
	}
	
	fmt.Println("Median Position: " + helpers.IntToString(median_pos))
	fmt.Println("Total Fuel Spent: " + helpers.IntToString(fuel_total))
}


func part2(swarm []int) {
	fmt.Println("====== PART TWO ======")
	total := 0
	for _,pos := range swarm {
		total += pos
	}
	avg_pos := total / len(swarm)

	fuel_total := 0
	for _,pos := range swarm {
		dif := helpers.Abs(pos - avg_pos)
		fuel_total += (dif * (dif + 1)) / 2
	}
	
	fmt.Println("Average Position: " + helpers.IntToString(avg_pos))
	fmt.Println("Total Fuel Spent: " + helpers.IntToString(fuel_total))
}

func main() {
	filename := "puzzle_input.txt"
	swarm := getSwarm(filename)
	part1(swarm)
	part2(swarm)
}