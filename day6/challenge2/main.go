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

	m := make(map[byte]int)
	var acc int = 0
	var tempacc int = 0

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) < 1 {
			var key byte
			var value int
			for key, value = range m {
				if value == tempacc && key > 0x60 {
					acc++
				}
				m[key] = 0
			}
			tempacc = 0
		} else {
			tempacc++
			for x := 0; x < len(line); x++ {
				if line[x] > 0x60 {
					m[line[x]]++
				}
			}
		}
	}
	println(acc)
}
