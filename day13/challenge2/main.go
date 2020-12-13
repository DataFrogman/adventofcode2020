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

	var temp []string
	var busses []int
	busses = make([]int, 0)

	flag := false
	for scanner.Scan() {
		var line string = scanner.Text()
		if !flag {
			flag = true
		}
		temp = strings.Split(line, ",")
	}

	for _, v := range temp {
		if v == "x" {
			busses = append(busses, -1)
		} else {
			temp2, _ := strconv.Atoi(v)
			busses = append(busses, temp2)
		}
	}

	var touples [][2]int
	for i, v := range busses {
		if v != -1 {
			touples = append(touples, [2]int{i, v})
		}
	}

	var acc int = 0
	var offset int = 1
	for _, v := range touples {
		for ((acc + v[0]) % v[1]) != 0 {
			acc += offset
		}
		offset = offset * v[1]
	}
	fmt.Println(acc)
}
