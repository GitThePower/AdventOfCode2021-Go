package helpers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Type Conversion Operations
func IntArrayToString(arr []int, delim string, width int) string {
	s := strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
	f_s := string(s[0])
	for i := 1; i < len(s); i++ {
		if (i % (width * (len(delim) + 1)) == 0) {
			f_s = f_s + "\n" + string(s[i])
		} else {
			f_s = f_s + string(s[i])
		}
	}
	return f_s
}

func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

func StringArrayToString(arr []string, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
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

func AppendTo2DStringArray(arr [][]string, val []string) [][]string {
	arr = append(arr, make([][]string, 1)...)
	arr[len(arr) - 1] = val
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

func DequeueIntArray(arr []int) ([]int, int) {
	pop := arr[0]
	arr = arr[1:]
	return arr, pop
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

func PopStringArray(arr []string) ([]string, string) {
	pop := arr[len(arr) - 1]
	arr = arr[:len(arr) - 1]
	return arr, pop
}

// Math Operations
func Abs(num int) int {
	n := float64(num)
	return int(math.Abs(n))
}

func Power(base, exp int) int {
	n := float64(exp)
	x := float64(base)
	return int(math.Pow(x, n))
}

// Map Operations
func InIntBoolMap(m map[int]bool, i int) bool {
	if _, in := m[i]; in { return true }
	return false
}

func InStringIntMap(m map[string]int, s string) bool {
	if _, in := m[s]; in { return true }
	return false
}