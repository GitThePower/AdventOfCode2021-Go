package helpers

import "log"
import "strconv"

func StringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil { log.Fatal(e) }
	return i
}

func CopyArray(arr []int) []int {
	arr_copy := make([]int, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func LeftShiftArray(arr []int, val int) []int {
	arr = append(arr[1:], []int{val}...)
	return arr
}

