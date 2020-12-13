package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
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

	var cartBusses []bigmod.CrtEntry
	for i, v := range busses {
		if v == -1 {
			continue
		}
		tempA := big.NewInt(int64(v - i))
		tempN := big.NewInt(int64(v))
		tempCRT := bigmod.CrtEntry{A: tempA, N: tempN}
		cartBusses = append(cartBusses, tempCRT)
	}
	fmt.Println(bigmod.SolveCrtMany(cartBusses).String())
}
