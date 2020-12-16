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
var validTickets [][]int = make([][]int, 0)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func verify(val int) bool {
	var matched bool = false
	for _, v := range attributes {
		if (val >= v[0] && val <= v[1]) || (val >= v[2] && val <= v[3]) {
			matched = true
			continue
		}
	}
	if !matched {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var newline int = 0
	var myTicket []int = make([]int, 0)

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
			temp := strings.Split(line, ",")
			for _, v := range temp {
				temp2, _ := strconv.Atoi(v)
				myTicket = append(myTicket, temp2)
			}
		} else if newline == 2 {
			if line == "nearby tickets:" {
				continue
			}
			temp := strings.Split(line, ",")
			newTicket := make([]int, 0)
			if len(temp) != len(attributes) {
				// incorrect++
			} else {
				output := true
				for _, v := range temp {
					temp2, _ := strconv.Atoi(v)
					if !verify(temp2) {
						output = false
					}
					newTicket = append(newTicket, temp2)
				}
				if output {
					validTickets = append(validTickets, newTicket)
				}
			}
		}
	}
	var att map[int][]int = make(map[int][]int)
	for i := range attributes {
		att[i] = make([]int, 0)
	}

	for _, ticket := range validTickets {
		for i, v := range ticket {
			for i2, v2 := range attributes {
				if !((v >= v2[0] && v <= v2[1]) || (v >= v2[2] && v <= v2[3])) {
					att[i2] = append(att[i2], i)
				}
			}
		}
	}

	//var correct map[int]int = make(map[int]int)

	var result int = 1

	for x := 0; x < 20; x++ {
		for k, v := range att {
			if len(v) == 19 {
				for y := 0; y < 20; y++ {
					if !contains(v, y) {
						for a, b := range att {
							att[a] = append(b, y)
						}
						if k < 6 {
							result *= myTicket[y]
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}
