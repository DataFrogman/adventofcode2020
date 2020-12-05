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

	m := make(map[int]bool)

	for x := 0; x < (127*8 + 7); x++ {
		m[x] = false
	}

	var Row int = 0
	var Column int = 0
	for scanner.Scan() {
		var line string = scanner.Text()
		var bot int = 0
		var top int = 127
		Row = 0
		for x := 0; x < 7; x++ {
			if line[x] == 0x46 {
				top = top - (top-bot)/2 - 1
			} else {
				bot = bot + (top-bot)/2 + 1
			}
		}

		if line[7] == 0x46 {
			Row = bot
		} else {
			Row = bot
		}

		var l int = 0
		var r int = 7
		Column = 0
		for x := 7; x < 9; x++ {
			if line[x] == 0x4c {
				r = r - (r-l)/2 - 1
			} else {
				l = l + (r-l)/2 + 1
			}
		}

		if line[9] == 0x4c {
			Column = l
		} else {
			Column = r
		}
		m[(Row*8 + Column)] = true
	}

	var key int = 0
	var value bool = false

	for key, value = range m {
		//println(key)
		if !value && m[key-1] && m[key+1] {
			println(key)
		}
	}
}
