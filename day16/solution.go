package main

import (
	"AdventOfCode2021/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
)

type packet struct {
	binaryString string
	version int
	typeId int
	lit_val int
}

func translateToBinary(filename string) string {
	f, e := os.Open(filename)
	if e != nil { log.Fatal(e) }
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	hex_string := scanner.Text()
	if e := scanner.Err(); e != nil { log.Fatal(e) }

	hex_to_bin := map[rune]string{48: "0000", 49: "0001", 50: "0010", 51: "0011", 52: "0100", 53: "0101", 54: "0110", 55: "0111",
															56: "1000", 57: "1001", 65: "1010", 66: "1011", 67: "1100", 68: "1101", 69: "1110", 70: "1111"}
	bin_string := ""
	for _,h := range hex_string {
		bin_string += hex_to_bin[h]
	}

	return bin_string
}

func decodePacket(bin_string string, version, typeId, iter int) packet {
	pack, incr := packet{version: version, typeId: typeId}, 6
	if (typeId == 4) {
		lit_val_bin_string := ""
		for {
			prefix := bin_string[iter + incr]
			lit_val_bin_string += bin_string[iter + incr + 1:iter + incr + 5]
			incr += 5
			if (prefix == 48) {
				remainder := (iter + incr) % 4
				if (remainder == 0) {
					pack.binaryString = bin_string[iter:iter + incr]
				} else {
					pack.binaryString = bin_string[iter:iter + incr - remainder + 4]
				}
				break
			}
		}
		pack.lit_val = helpers.BinaryStringToInt(lit_val_bin_string)
	} else {
		identifier := bin_string[iter + incr]
		incr++
		if (identifier == 48) {
			
		} else {
			
		}
	}

	return pack
}

func getPackets(bin_string string) []packet {
	packets, iter, version_sum := make([]packet, 0), 0, 0

	for {
		if (iter > len(bin_string) - 1) {
			break
		}
		version, typeId := helpers.BinaryStringToInt(bin_string[iter:iter + 3]), helpers.BinaryStringToInt(bin_string[iter + 3:iter + 6])
		packets = append(packets, []packet{decodePacket(bin_string, version, typeId, iter)}...)
		iter += len(packets[len(packets) - 1].binaryString)
		version_sum += version
	}

	return packets
}

func main() {
	// filename := "puzzle_input.txt"
	// bin_string := translateToBinary(filename)
	iter := 0
	pack := decodePacket("110100101111111000101000", 6, 4, iter)
	iter += len(pack.binaryString)
	fmt.Println(pack)
	fmt.Println(iter)
	// packets := getPackets(bin_string)
	// fmt.Println(packets)
}