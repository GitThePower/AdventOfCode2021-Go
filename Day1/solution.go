package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func part1() {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open("depth_readings.txt")
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inc_c := 0
	last := 9223372036854775807
	for scanner.Scan() {
		i, e := strconv.Atoi(scanner.Text())
		if e != nil { log.Fatal(e) }

		if i > last { inc_c++ }
		last = i
	}

	fmt.Println("Incs: " + strconv.Itoa(inc_c))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func part2() {
	fmt.Println("====== PART TWO ======")

	f, e := os.Open("depth_readings.txt")
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	win_inc_c := 0
	win := make([]int, 3)
	startCounter := 0
	for scanner.Scan() {
		i, e := strconv.Atoi(scanner.Text())
		if e != nil { log.Fatal(e) }

		last_win := make([]int, 3)
		copy(last_win, win)
		win = append(win[1:], []int{i}...)

		if startCounter >= 3 {
			win_diff := 0
			for j := 0; j < len(win); j++ {
				win_diff = win_diff + last_win[j] - win[j]
			}
			if win_diff < 0 { win_inc_c++ }
		}
		startCounter++
	}

	fmt.Println("Win Incs: " + strconv.Itoa(win_inc_c))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func main() {
	part1()
	part2()
}
