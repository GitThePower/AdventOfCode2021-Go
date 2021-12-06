package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getBitCounts(counts [][]int, num string) [][]int {
	for j := 0; j < len(num); j++ {
		if num[j] == 48 {
			counts[j][0]++
		}
		if num[j] == 49 {
			counts[j][1]++
		}
	}
	return counts
}

func getGammaAndEpsilon(counts [][]int) (int, int) {
	gamma, epsilon := 0, 0
	for j := 0; j < len(counts); j++ {
		pow_of_two := helpers.Power(2, len(counts)-1-j)
		if counts[j][0] > counts[j][1] {
			epsilon += pow_of_two
		}
		if counts[j][0] < counts[j][1] {
			gamma += pow_of_two
		}
		if counts[j][0] == counts[j][1] {
			log.Fatal("Counts of bit were equal")
		}
	}
	return gamma, epsilon
}

func printResultsP1(gamma, epsilon int) {
	fmt.Println("Gamma: " + helpers.IntToString(gamma))
	fmt.Println("Epsilon: " + helpers.IntToString(epsilon))
	fmt.Println("Product: " + helpers.IntToString(gamma*epsilon))
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	counts := make([][]int, 1)
	counts[0] = make([]int, 2)
	for scanner.Scan() {
		num := scanner.Text()
		if len(counts) < len(num) {
			counts = helpers.Extend2DIntArray(counts, len(num)-len(counts), 2)
		}
		counts = getBitCounts(counts, num)
	}

	gamma, epsilon := getGammaAndEpsilon(counts)
	printResultsP1(gamma, epsilon)

	if e := scanner.Err(); e != nil {
		log.Fatal(e)
	}
}

func getJthBitCount(counts []int, num string, j int) []int {
	if j < len(num) {
		if num[j] == 48 {
			counts[0]++
		}
		if num[j] == 49 {
			counts[1]++
		}
	}
	return counts
}

func getRating(num_arr []string, counts []int, match bool) int {
	rating := 0
	ratings := helpers.CopyStringArray(num_arr)
	for j := 0; j < len(ratings[0]); j++ {
		mode := byte(49)
		if counts[0] > counts[1] {
			mode = byte(48)
		}

		next_counts := make([]int, 2)
		next_ratings := make([]string, 0)
		for _, s := range ratings {
			if (match && s[j] == mode) || (!match && s[j] != mode) {
				next_counts = getJthBitCount(next_counts, s, j+1)
				next_ratings = helpers.AppendToStringArray(next_ratings, s)
			}
		}

		if len(next_ratings) == 1 {
			binary_string := next_ratings[0]
			for k := 0; k < len(binary_string); k++ {
				if binary_string[k] == 49 {
					rating += helpers.Power(2, len(binary_string)-1-k)
				}
			}
			return rating
		}
		counts = next_counts
		ratings = next_ratings
	}

	log.Fatal("Rating never converged")
	return rating - 1
}

func printResultsP2(O2_rating, CO2_rating int) {
	fmt.Println("O2 Generator Rating: " + helpers.IntToString(O2_rating))
	fmt.Println("CO2 Scrubber Rating: " + helpers.IntToString(CO2_rating))
	fmt.Println("Product: " + helpers.IntToString(O2_rating*CO2_rating))
}

func part2(filename string) {
	fmt.Println("====== PART TWO ======")

	f, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	num_arr := make([]string, 0)
	counts := make([]int, 2)
	for scanner.Scan() {
		num := scanner.Text()
		num_arr = helpers.AppendToStringArray(num_arr, num)
		counts = getJthBitCount(counts, num, 0)
	}

	O2_rating := getRating(num_arr, helpers.CopyIntArray(counts), true)
	CO2_rating := getRating(num_arr, helpers.CopyIntArray(counts), false)
	printResultsP2(O2_rating, CO2_rating)

	if e := scanner.Err(); e != nil {
		log.Fatal(e)
	}
}

func main() {
	filename := "puzzle_input.txt"
	part1(filename)
	part2(filename)
}
