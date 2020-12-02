package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var m map[int]bool
	m = make(map[int]bool)

	for scanner.Scan() {
		var temp = scanner.Text()
		var temp2, _ = strconv.Atoi(temp)
		m[temp2] = true
	}

	for key := range m {
		if m[(2020-key)] == true {
			println(key * (2020 - key))
			break
		}
	}
}
