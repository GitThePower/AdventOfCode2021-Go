package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func getBitCounts(counts [][]int, num string) [][]int {
	for j := 0; j < len(num); j++ {
		if num[j] == 48 { counts[j][0]++ }
		if num[j] == 49 { counts[j][1]++ }
	}
	return counts
}

func getGammaAndEpsilon(counts [][]int) (int, int) {
	gamma := 0
	epsilon := 0
	for j := 0; j < len(counts); j++ {
		pow_of_two := float64(len(counts) - 1 - j)
		if counts[j][0] > counts[j][1] { epsilon += int(math.Pow(2, pow_of_two)) }
		if counts[j][0] < counts[j][1] { gamma += int(math.Pow(2, pow_of_two)) }
		if counts[j][0] == counts[j][1] { log.Fatal("Counts of bit were equal") }
	}
	return gamma, epsilon
}

func printResults(gamma, epsilon int) {
	fmt.Println("Gamma: " + helpers.IntToString(gamma))
	fmt.Println("Epsilon: " + helpers.IntToString(epsilon))
	fmt.Println("Product: " + helpers.IntToString(gamma * epsilon))
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	counts := make([][]int, 1)
	counts[0] = make([]int, 2)
	for scanner.Scan() {
		num := scanner.Text()
		if len(counts) < len(num) {
			counts = helpers.TwoDIntArrayExtend(counts, len(num) - len(counts), 2)
		}
		counts = getBitCounts(counts, num)
	}

	gamma, epsilon := getGammaAndEpsilon(counts)
	printResults(gamma, epsilon)

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
}
