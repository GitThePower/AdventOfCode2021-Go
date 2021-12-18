package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type path struct {
	idx int
	risk int
}

func getCave(filename string) ([]int, int, int) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cave, width, height := make([]int, 0), 0, 0
	for scanner.Scan() {
		caveRow := scanner.Text()
		width = len(caveRow)
		height++
		for i := range caveRow {
			cave = helpers.AppendToIntArray(cave, helpers.StringToInt(string(caveRow[i])))
		}
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return cave, width, height
}

func getNextPaths(p path, cave []int, width, height int, idx_seen map[int]int) []path {
	next_paths := make([]path, 0)
	if (p.idx >= width) {
		idx_u := p.idx - width
		risk_u := p.risk + cave[idx_u]
		if (!helpers.InIntIntMap(idx_seen, idx_u) || risk_u < idx_seen[idx_u]) {
			idx_seen[idx_u] = risk_u
			next_paths = append(next_paths, []path{{idx: idx_u, risk: risk_u}}...)
		}
	}
	if (p.idx % width > 0) {
		idx_l := p.idx - 1
		risk_l := p.risk + cave[idx_l]
		if (!helpers.InIntIntMap(idx_seen, idx_l) || risk_l < idx_seen[idx_l]) {
			idx_seen[idx_l] = risk_l
			next_paths = append(next_paths, []path{{idx: idx_l, risk: risk_l}}...)
		}
	}
	if (p.idx % width < width - 1) {
		idx_r := p.idx + 1
		risk_r := p.risk + cave[idx_r]
		if (!helpers.InIntIntMap(idx_seen, idx_r) || risk_r < idx_seen[idx_r]) {
			idx_seen[idx_r] = risk_r
			next_paths = append(next_paths, []path{{idx: idx_r, risk: risk_r}}...)
		}
	}
	if (p.idx < width * (height - 1)) {
		idx_d := p.idx + width
		risk_d := p.risk + cave[idx_d]
		if (!helpers.InIntIntMap(idx_seen, idx_d) || risk_d < idx_seen[idx_d]) {
			idx_seen[idx_d] = risk_d
			next_paths = append(next_paths, []path{{idx: idx_d, risk: risk_d}}...)
		}
	}
	return next_paths
}

func escapeTileBottomRightCorner(p_queue []path, cave []int, width, height int) int {
	tile_risk, idx_seen := -1, make(map[int]int)
	
	for {
		if (len(p_queue) < 1) {
			break
		}
		pop := p_queue[0]
		p_queue = p_queue[1:]
		next_paths := getNextPaths(pop, cave, width, height, idx_seen)
		for _,path := range next_paths {
			if (path.idx == len(cave) - 1 && (tile_risk < 0 || path.risk < tile_risk)) {
				tile_risk = path.risk
			}
		}
		if (tile_risk > 0) {
			break
		}
		p_queue = append(p_queue, next_paths...)
		sort.Slice(p_queue, func (i, j int) bool {
			return p_queue[i].risk < p_queue[j].risk })
	}
	
	return tile_risk
}

func part1(cave []int, width, height int) {
	fmt.Println("====== PART ONE ======")
  total_risk := escapeTileBottomRightCorner([]path{{0, 0}}, cave, width, height)

	fmt.Println("Total Risk: " + helpers.IntToString(total_risk))
}

func get5XCave(filename string) ([]int, int, int) {
	cave, width, height := make([]int, 0), 0, 0

	for y := 0; y < 5; y++ {
		f, e := os.Open(filename)
		if e != nil { log.Fatal(e) }
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			caveRow := scanner.Text()
			width = len(caveRow) * 5
			height++
			for x := 0; x < 5; x++ {
				for i := range caveRow {
					val := (helpers.StringToInt(string(caveRow[i])) + x + y) % 9
					if (val == 0) { val = 9 }
					cave = helpers.AppendToIntArray(cave, val)
				}
			}
		}

		if e := scanner.Err(); e != nil { log.Fatal(e) }
	}

	return cave, width, height
}

func part2(cave []int, width, height int) {
	fmt.Println("====== PART TWO ======")
  total_risk := escapeTileBottomRightCorner([]path{{0, 0}}, cave, width, height)

	fmt.Println("Total Risk: " + helpers.IntToString(total_risk))
}

func main() {
	filename := "puzzle_input.txt"
	cave, width, height := getCave(filename)
	part1(cave, width, height)
	cave, width, height = get5XCave(filename)
	part2(cave, width, height)
}