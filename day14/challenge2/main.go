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

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func eval(mask string, num uint64, list *map[uint64]int) {
	var temp uint64 = uint64(num)
	var flag bool = false
	for i, v := range mask {
		if v == '1' {
			temp = setBit(temp, uint(i))
		} else if v == 'X' && !flag {
			flag = true
			temp2 := mask
			temp2 = replaceAtIndex(temp2, '0', i)
			temp3 := setBit(temp, uint(i))
			eval(temp2, temp3, list)
			temp4 := clearBit(temp, uint(i))
			temp5 := mask
			temp5 = replaceAtIndex(temp2, '0', i)
			eval(temp5, temp4, list)
		}
	}
	if !flag {
		(*list)[temp]++
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var mask string
	var bank map[uint64]uint64 = make(map[uint64]uint64)

	for scanner.Scan() {
		var line string = scanner.Text()
		if line[0:4] == "mask" {
			mask = line[7:]
		} else {
			temp := strings.Split(line, "= ")
			temp2, _ := strconv.Atoi(temp[1])
			temp3 := strings.Split(line, "]")
			temp4, _ := strconv.Atoi(temp3[0][4:])
			runes := []rune(mask)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			tempRunes := string(runes)
			var temp5 *map[uint64]int
			temp6 := make(map[uint64]int)
			temp5 = &temp6
			eval(tempRunes, uint64(temp4), temp5)
			for k := range *temp5 {
				bank[k] = uint64(temp2)
			}
		}
	}
	var acc uint64 = 0
	for _, v := range bank {
		acc += v
	}
	fmt.Println(acc)
}
