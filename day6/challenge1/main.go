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

	for scanner.Scan() {
		var line string = scanner.Text()
	}
}
