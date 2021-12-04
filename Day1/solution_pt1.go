package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
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
