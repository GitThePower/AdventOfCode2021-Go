package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getSchool(filename string) []int {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s_school := strings.Split(scanner.Text(), ",")
	n_school := make([]int, len(s_school))
	for i,fish := range s_school {
		n_school[i] = helpers.StringToInt(fish)
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	consol_school := make([]int, 9)
	for _,val := range n_school { consol_school[val]++ }
	return consol_school
}

func proliferateFish(school []int, days int) {
	for j := 0; j < days; j++ {
		new_fish := school[0]
		school = helpers.LeftShiftIntArray(school, new_fish)
		school[6] += new_fish
	}

	total_fish := 0
	for _,fish := range school { total_fish += fish }
	fmt.Println("Number of Fish: " + helpers.IntToString(total_fish))
}

func main() {
	filename := "puzzle_input.txt"
	school := getSchool(filename)
	fmt.Println("====== PART ONE ======")
	proliferateFish(school, 80)
	fmt.Println("====== PART TWO ======")
	proliferateFish(school, 256)
}
