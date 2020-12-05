package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var max int = 0

	for scanner.Scan() {
		var line string = scanner.Text()
		var bot int = 0
		var top int = 127
		var row int = 0
		for x := 0; x < 7; x++ {
			if line[x] == 0x46 {
				top = top - (top-bot)/2 - 1
			} else {
				bot = bot + (top-bot)/2 + 1
			}
		}

		if line[7] == 0x46 {
			row = bot
		} else {
			row = bot
		}

		var l int = 0
		var r int = 7
		var column int = 0
		for x := 7; x < 9; x++ {
			if line[x] == 0x4c {
				r = r - (r-l)/2 - 1
			} else {
				l = l + (r-l)/2 + 1
			}
		}

		if line[9] == 0x4c {
			column = l
		} else {
			column = r
		}
		if max < (row*8 + column) {
			max = row*8 + column
		}
	}
	println(max)
}
