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
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input map[int64]int64 = make(map[int64]int64)
	var stringInput []string = make([]string, 0)
	var speaking int64
	var nextNum int64

	for scanner.Scan() {
		var line string = scanner.Text()
		stringInput = strings.Split(line, ",")
	}
	for i, v := range stringInput {
		temp, _ := strconv.Atoi(v)
		speaking = int64(temp)
		input[int64(temp)] = int64(i + 1)
	}
	speaking = 0
	if val, ok := input[speaking]; ok {
		nextNum = int64(len(stringInput)+1) - val
		input[speaking] = int64(len(stringInput) + 1)
	} else {
		nextNum = 0
		input[speaking] = int64(len(stringInput) + 1)
	}
	for x := int64(len(stringInput) + 2); x < 30000001; x++ {
		if val, ok := input[nextNum]; ok {
			speaking = nextNum
			input[speaking] = x
			nextNum = x - val
		} else {
			speaking = nextNum
			input[nextNum] = x
			nextNum = 0
		}
	}
	fmt.Println(speaking)
}
