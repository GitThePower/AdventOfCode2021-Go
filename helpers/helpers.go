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

func TwoDIntArrayToString(arr [][]int, delim string, width int) string {
	f_s_2d := ""
	for i := range arr {
		s := strings.Trim(strings.Replace(fmt.Sprint(arr[i]), " ", delim, -1), "[]")
		f_s := string(s[0])
		for i := 1; i < len(s); i++ {
			if (i % (width * (len(delim) + 1)) == 0) {
				f_s = f_s + "\n" + string(s[i])
			} else {
				f_s = f_s + string(s[i])
			}
		}
		f_s_2d = f_s_2d + f_s + "\n"
	}
	return f_s_2d
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

func Copy2DIntArray(arr [][]int) [][]int {
	arr_copy := make([][]int, len(arr))
	copy(arr_copy, arr)
	for i := range arr_copy {
		copy(arr_copy[i], arr[i])
	}
	return arr_copy
}

func DequeueIntArray(arr []int) ([]int, int) {
	pop := arr[0]
	arr = arr[1:]
	return arr, pop
}


func Dequeue2DStringArray(arr [][]string) ([][]string, []string) {
	pop := arr[0]
	arr = arr[1:]
	return arr, pop
}

func InStringArray(arr []string, s string) bool {
	for _,val := range arr {
		if val == s {
			return true
		}
	}
	return false
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

func Resize2DIntArray(arr [][]int, x, y int) [][]int {
	if (len(arr) < x) {
		arr = append(arr, make([][]int, x - len(arr))...)
	} else {
		arr = arr[:x]
	}
	for j := 0; j < len(arr); j++ {
		if (len(arr[j]) < y) {
			arr[j] = append(arr[j], make([]int, y - len(arr[j]))...)
		} else {
			arr[j] = arr[j][:y]
		}
	}
	return arr
}

// Math Operations
func Abs(num int) int {
	n := float64(num)
	return int(math.Abs(n))
}

func BinaryStringToInt(b string) int {
	i := 0
	for j := 0; j < len(b); j++ {
		if b[j] == 49 {
			i += Power(2, len(b) - 1 - j)
		}
	}
	return i
}

func Power(base, exp int) int {
	n := float64(exp)
	x := float64(base)
	return int(math.Pow(x, n))
}

// Map Operations
func CopyStringIntMap(old_m map[string]int) map[string]int {
	new_m := make(map[string]int)
	for k,v := range old_m {
		new_m[k] = v
	}
	return new_m
}

func InIntBoolMap(m map[int]bool, i int) bool {
	if _, in := m[i]; in { return true }
	return false
}

func InIntIntMap(m map[int]int, i int) bool {
	if _, in := m[i]; in { return true }
	return false
}

func InStringBoolMap(m map[string]bool, s string) bool {
	if _, in := m[s]; in { return true }
	return false
}

func InStringIntMap(m map[string]int, s string) bool {
	if _, in := m[s]; in { return true }
	return false
}

func InStringStringArrayMap(m map[string][]string, s string) bool {
	if _, in := m[s]; in { return true }
	return false
}

func InStringStringMap(m map[string]string, s string) bool {
	if _, in := m[s]; in { return true }
	return false
}

func SafeInsertStringIntMap(m map[string]int, s string, i int) map[string]int {
	if (InStringIntMap(m, s)) {
		m[s] += i
	} else {
		m[s] = i
	}

	return m
}

// String operations
func IsUpper(s string) bool {
	for _, r := range s {
		if r < 65 || r > 90 {
				return false
		}
	}
	return true
}