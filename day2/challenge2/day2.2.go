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
		var targetbyte = s[1][0]
		var value string = s[2]
		var temp = strings.Split(s[0], "-")

		var x, _ = strconv.Atoi(temp[0])
		var x2 = value[x-1]
		var y, _ = strconv.Atoi(temp[1])
		var y2 = value[y-1]

		if (x2 == targetbyte || y2 == targetbyte) && (x2 != y2) {
			acc++
		}
	}
	println(acc)
}
