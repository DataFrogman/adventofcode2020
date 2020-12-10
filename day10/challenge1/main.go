package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"sort"

	//"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	s := make([]int, 0)

	for scanner.Scan() {
		var line string = scanner.Text()

		temp, _ := strconv.Atoi(line)
		s = append(s, temp)
	}

	var num1 int = 0
	var num3 int = 0
	var prev int = 0

	sort.Ints(s)
	for _, value := range s {
		if value-prev == 1 {
			prev = value
			num1++
		} else if value-prev == 3 {
			prev = value
			num3++
		} else {
			prev = value
		}
	}
	num3++
	println(num1 * num3)

}
