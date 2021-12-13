package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCaves(filename string) map[string][]string {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	caves := make(map[string][]string)
	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		for i := range path {
			curr_cave := path[i]
			adj_cave := path[(i + 1) % len(path)]
			if (helpers.InStringStringArrayMap(caves, curr_cave)) {
				caves[curr_cave] = helpers.AppendToStringArray(caves[curr_cave], adj_cave)
			} else {
				caves[curr_cave] = []string{adj_cave}
			}
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return caves
}

func part1(caves map[string][]string) {
	fmt.Println("====== PART ONE ======")
	queue, unique_paths, pop := make([][]string, 0), 0, make([]string, 0)

	queue = helpers.AppendTo2DStringArray(queue, []string{"start"})
	for {
		if len(queue) < 1 {
			break
		}
		queue,pop = helpers.Dequeue2DStringArray(queue)
		caves_ahead := caves[pop[len(pop) - 1]]
		for _,cave := range caves_ahead {
			if (cave == "end") {
				unique_paths++
			} else if (helpers.IsUpper(cave) || !helpers.InStringArray(pop, cave)) {
				path := helpers.AppendToStringArray(pop, cave)
				queue = helpers.AppendTo2DStringArray(queue, helpers.CopyStringArray(path))
			}
		}
	}

	fmt.Println("Unique Paths: " + helpers.IntToString(unique_paths))
}

func canRevisitSmCave(path []string) bool {
	cave_counts := make(map[string]bool)
	for _,cave := range path {
		if (helpers.InStringBoolMap(cave_counts, cave) && !helpers.IsUpper(cave)) {
			return false
		} else {
			cave_counts[cave] = true
		}
	}

	return true
}

func part2(caves map[string][]string) {
	fmt.Println("====== PART TWO ======")
	queue, unique_paths, pop := make([][]string, 0), 0, make([]string, 0)

	queue = helpers.AppendTo2DStringArray(queue, []string{"start"})
	for {
		if len(queue) < 1 {
			break
		}
		queue,pop = helpers.Dequeue2DStringArray(queue)
		caves_ahead := caves[pop[len(pop) - 1]]
		for _,cave := range caves_ahead {
			if (cave == "end") {
				unique_paths++
			} else if (helpers.IsUpper(cave) || !helpers.InStringArray(pop, cave) || (cave != "start" && canRevisitSmCave(pop))) {
				path := helpers.AppendToStringArray(pop, cave)
				queue = helpers.AppendTo2DStringArray(queue, helpers.CopyStringArray(path))
			}
		}
	}

	fmt.Println("Unique Paths: " + helpers.IntToString(unique_paths))
}

func main() {
	filename := "puzzle_input.txt"
	caves := getCaves(filename)
	part1(caves)
	part2(caves)
}