package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var acc = 0

	for scanner.Scan() {
		var line = scanner.Text()
		s := strings.Split(line, " ")
		target := s[1][:len(s[1])-1]
		count := strings.Count(s[2], target)

		var temp = strings.Split(s[0], "-")

		var x, _ = strconv.Atoi(temp[0])
		var y, _ = strconv.Atoi(temp[1])

		if x <= count {
			if count <= y {
				acc++
			}
		}
	}
	println(acc)
}
