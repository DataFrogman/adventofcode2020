package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func setBit(n uint64, pos uint) uint64 {
	n |= (1 << pos)
	return n
}

func clearBit(n uint64, pos uint) uint64 {
	var mask uint64 = ^(1 << pos)
	n &= mask
	return n
}

func eval(mask string, num int) uint64 {
	var temp uint64 = uint64(num)
	runes := []rune(mask)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	mask = string(runes)
	for i, v := range mask {
		if v == '1' {
			temp = setBit(temp, uint(i))
		} else if v == '0' {
			temp = clearBit(temp, uint(i))
		}
	}
	return temp
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var mask string
	var bank map[int]uint64 = make(map[int]uint64)

	for scanner.Scan() {
		var line string = scanner.Text()
		if line[0:4] == "mask" {
			mask = line[7:]
		} else {
			temp := strings.Split(line, "= ")
			temp2, _ := strconv.Atoi(temp[1])
			temp3 := strings.Split(line, "]")
			temp4, _ := strconv.Atoi(temp3[0][4:])
			bank[temp4] = eval(mask, temp2)
		}
	}
	var acc uint64 = 0
	for _, v := range bank {
		acc += v
	}
	fmt.Println(acc)
}
