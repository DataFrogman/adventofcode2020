package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var mastermap map[int]int64

func solve(curr int, master []int) int64 {
	if _, ok := mastermap[curr]; ok {
		return mastermap[curr]
	}
	if len(master) == 0 {
		mastermap[curr] = 1
	} else {
		var temp int64 = 0
		for i, v := range master {
			if v-curr <= 3 {
				temp += solve(v, master[i+1:])
			} else {
				break
			}
		}
		mastermap[curr] = temp
	}
	return mastermap[curr]
}

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

	s = append(s, 0)
	sort.Ints(s)
	//m := make(map[int][]int)
	//var acc int64 = 0

	// for i, v := range s {
	// 	var temp []int
	// 	for x := i + 1; x < i+4; x++ {
	// 		if x < len(s) && s[x] <= v+3 {
	// 			temp = append(temp, s[x])
	// 		}
	// 	}
	// 	if len(temp) != 0 {
	// 		m[v] = temp
	// 	}
	// }
	// fmt.Println(m)
	mastermap = make(map[int]int64)
	fmt.Println(solve(0, s) / 2)
	//println(acc)
}
