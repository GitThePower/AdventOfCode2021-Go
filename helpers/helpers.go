package helpers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Type Conversion Operations
func IntArrayToString(arr []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}

func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

func StringToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

// Array Operations
func AppendToIntArray(arr []int, val int) []int {
	arr = append(arr, []int{val}...)
	return arr
}

func AppendToStringArray(arr []string, val string) []string {
	arr = append(arr, []string{val}...)
	return arr
}

func CopyIntArray(arr []int) []int {
	arr_copy := make([]int, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func CopyStringArray(arr []string) []string {
	arr_copy := make([]string, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func Extend2DIntArray(arr [][]int, x, y int) [][]int {
	prev_len := len(arr)
	arr = append(arr, make([][]int, x)...)
	for j := prev_len; j < len(arr); j++ {
		arr[j] = make([]int, y)
	}
	return arr
}

func LeftShiftIntArray(arr []int, val int) []int {
	arr = append(arr[1:], []int{val}...)
	return arr
}

// Math Operations
func Power(base, exp int) int {
	n := float64(exp)
	return int(math.Pow(2, n))
}
