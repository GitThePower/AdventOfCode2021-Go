package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getSignalsAndOutputs(filename string) ([][]string, [][]string) {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	signal_patterns := make([][]string, 0)
	outputs := make([][]string, 0)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		signal_patterns = helpers.AppendTo2DStringArray(signal_patterns, strings.Fields(parts[0])) 
		outputs = helpers.AppendTo2DStringArray(outputs, strings.Fields(parts[1])) 
	}

	if e := scanner.Err(); e != nil { log.Fatal(e) }

	return signal_patterns, outputs
}

func part1(outputs [][]string) {
	fmt.Println("====== PART ONE ======")

	unique_count := 0
	for _,output := range outputs {
		for _,display := range output {
			if len(display) == 2 || len(display) == 3 || len(display) == 4 || len(display) == 7 {
				unique_count++
			}
		}
	}

	fmt.Println("Instances of 1, 4, 7, and 8: " + helpers.IntToString(unique_count))
}

func signalSimilarity(a, b string) int {
	similarity := 0
	for i := range a {
		if (strings.Contains(b, string(a[i]))) { similarity++ }
	}

	return similarity
}

func translateSignals(signals []string) map[int]string {
	translation := make(map[int]string)
	unknown := make([]string, 0)
	for _, signal := range signals {
		if (len(signal) == 2) { translation[1] = signal }
		if (len(signal) == 3) { translation[7] = signal }
		if (len(signal) == 4) { translation[4] = signal }
		if (len(signal) == 5 || len(signal) == 6) { unknown = helpers.AppendToStringArray(unknown, signal) }
		if (len(signal) == 7) { translation[8] = signal }
	}

	for _,signal := range unknown {
		if (len(signal) == 5 && signalSimilarity(translation[1], signal) == 1 && signalSimilarity(translation[4], signal) == 2) {
			translation[2] = signal
		} else if (len(signal) == 5 && signalSimilarity(translation[1], signal) == 2) {
			translation[3] = signal
		} else if (len(signal) == 5) {
			translation[5] = signal
		} else if (signalSimilarity(translation[1], signal) == 1) {
			translation[6] = signal
		} else if (signalSimilarity(translation[4], signal) == 4) {
			translation[9] = signal
		} else {
			translation[0] = signal
		}
	}

	return translation
}

func decodeOutput(output string, translation map[int]string, idx int) int {
	for j := 0; j < 10; j++ {
		if (len(output) == len(translation[j]) && signalSimilarity(output, translation[j]) == len(output)) {
			return j * helpers.Power(10, 3 - idx)
		}
	}
	return 0
}

func part2(signal_patterns, outputs [][]string) {
	fmt.Println("====== PART TWO ======")

	sum_display_outputs := 0
	for i := range signal_patterns {
		translation := translateSignals(signal_patterns[i])
		for j,output := range outputs[i] {
			sum_display_outputs += decodeOutput(output, translation, j)
		}
	}

	fmt.Println("Sum of Display Outputs: " + helpers.IntToString(sum_display_outputs))
}

func main() {
	filename := "puzzle_input.txt"
	signal_patterns, outputs := getSignalsAndOutputs(filename)
	part1(outputs)
	part2(signal_patterns, outputs)
}