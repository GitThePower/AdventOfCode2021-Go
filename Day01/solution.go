package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func stringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil { log.Fatal(e) }
	return i
}

func copyArray(arr []int) []int {
	arr_copy := make([]int, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func leftShiftArray(arr []int, val int) []int {
	arr = append(arr[1:], []int{val}...)
	return arr
}

func part1(filename string) {
	fmt.Println("====== PART ONE ======")

	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inc_c, last := 0, 9223372036854775807
	for scanner.Scan() {
		x := stringToInt(scanner.Text())

		if x > last { inc_c++ }
		last = x
	}

	fmt.Println("Incs: " + strconv.Itoa(inc_c))

	if e := scanner.Err(); e != nil { log.Fatal(e) }
}

func part2(filename string) {
	fmt.Println("====== PART TWO ======")

	f, e := os.Open("puzzle_input.txt")
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	startCounter, win_inc_c, win := 0, 0, make([]int, 3)
	for scanner.Scan() {
		x := stringToInt(scanner.Text())

		last_win := copyArray(win)
		win = leftShiftArray(win, x)

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
	filename := "puzzle_input.txt"
	part1(filename)
	part2(filename)
}
