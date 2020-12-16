package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var attributes [][]int = make([][]int, 0)
var incorrect int = 0

func verify(val int) {
	var matched bool = false
	for _, v := range attributes {
		if (val >= v[0] && val <= v[1]) || (val >= v[2] && val <= v[3]) {
			matched = true
			continue
		}
	}
	if !matched {
		incorrect += val
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var newline int = 0

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) < 2 {
			newline++
		} else if newline == 0 {
			temp := strings.Split(line, ": ")
			temp = strings.Split(temp[1], " or ")
			leftvals := strings.Split(temp[0], "-")
			rightvals := strings.Split(temp[1], "-")
			tempAttributes := make([]int, 0)
			temp2, _ := strconv.Atoi(leftvals[0])
			tempAttributes = append(tempAttributes, temp2)
			temp2, _ = strconv.Atoi(leftvals[1])
			tempAttributes = append(tempAttributes, temp2)
			temp2, _ = strconv.Atoi(rightvals[0])
			tempAttributes = append(tempAttributes, temp2)
			temp2, _ = strconv.Atoi(rightvals[1])
			tempAttributes = append(tempAttributes, temp2)
			attributes = append(attributes, tempAttributes)
		} else if newline == 1 {
			if line == "your ticket:" {
				continue
			}
			// ignore until part 2
			continue
		} else if newline == 2 {
			if line == "nearby tickets:" {
				continue
			}
			temp := strings.Split(line, ",")
			if len(temp) != len(attributes) {
				// incorrect++
			} else {
				for _, v := range temp {
					temp2, _ := strconv.Atoi(v)
					verify(temp2)
				}
			}
		}
	}
	fmt.Println(incorrect)
}
