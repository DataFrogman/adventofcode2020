package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_3_1.txt")

	if err != nil {
		log.Fatalf("Failed to open file", err)
	}

	var mapArray [31][323]bool

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var acc = 0
	var period byte = 0x2e

	for scanner.Scan() {
		var line string = scanner.Text()
		for x := 0; x < 31; x++ {
			if line[x] == period {
				mapArray[x][acc] = false
			} else {
				mapArray[x][acc] = true
			}
		}
		acc++
	}
	var x = 0
	var y = 0
	acc = 0

	for y < 322 {
		if (x + 1) >= 31 {
			x = (x + 1) % 31
		} else {
			x = x + 1
		}
		y += 2

		if mapArray[x][y] {
			acc++
		}
	}
	println(acc)
}
