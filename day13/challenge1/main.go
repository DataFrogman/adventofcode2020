package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	timestamp := 0
	busses := make(map[int]int)
	var temp []string

	flag := false
	for scanner.Scan() {
		var line string = scanner.Text()
		if !flag {
			timestamp, _ = strconv.Atoi(line)
			flag = true
		}
		temp = strings.Split(line, ",")
	}

	for _, v := range temp {
		if v != "x" {
			temp2, _ := strconv.Atoi(v)
			busses[temp2] = (timestamp/temp2+1)*temp2 - timestamp
		}
	}

	minDistance := timestamp
	minBus := 0
	for k, v := range busses {
		if v < minDistance {
			minDistance = v
			minBus = k
		}
	}
	fmt.Println(minBus * minDistance)
}
