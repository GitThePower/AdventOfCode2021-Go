package helpers

import (
	"log"
	"strconv"
)

// Type Conversion Operations
func StringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

// Array Operations
func CopyIntArray(arr []int) []int {
	arr_copy := make([]int, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func LeftShiftIntArray(arr []int, val int) []int {
	arr = append(arr[1:], []int{val}...)
	return arr
}

func TwoDIntArrayExtend(arr [][]int, x, y int) [][]int {
	prev_len := len(arr)
	arr = append(arr, make([][]int, x)...)
	for j := prev_len; j < len(arr); j++ {
		arr[j] = make([]int, y)
	}
	return arr
}
