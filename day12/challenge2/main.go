package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func locationFetcher(l []string, s string) int {
	for i, v := range l {
		if v == s {
			return i
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var directions []string = make([]string, 0)
	var norSou int = 1
	var wesEas int = 10
	var actualNorSou int = 0
	var actualwesEas int = 0
	// var dir string = "E"
	// cards := []string{"N", "E", "S", "W"}
	// cardsBack := []string{"W", "S", "E", "N"}

	for scanner.Scan() {
		var line string = scanner.Text()
		line = strings.ReplaceAll(line, "\n", "")
		directions = append(directions, line)
	}
	for _, v := range directions {
		temp, _ := strconv.Atoi(v[1:])
		switch v[0] {
		case 0x4e:
			norSou += temp
		case 0x53:
			norSou -= temp
		case 0x45:
			wesEas += temp
		case 0x57:
			wesEas -= temp
		case 0x46:
			actualNorSou += temp * norSou
			actualwesEas += temp * wesEas
		case 0x52:
			temp2 := (temp / 90) % 4
			if temp2 == 0 {
				continue
			} else if temp2 == 1 {
				temp = wesEas
				wesEas = norSou
				norSou = -temp
			} else if temp2 == 2 {
				temp = wesEas
				wesEas = -wesEas
				norSou = -norSou
			} else if temp2 == 3 {
				temp = wesEas
				wesEas = -norSou
				norSou = temp
			}
		case 0x4c:
			temp2 := (temp / 90) % 4
			if temp2 == 0 {
				continue
			} else if temp2 == 1 {
				temp = wesEas
				wesEas = -norSou
				norSou = temp
			} else if temp2 == 2 {
				wesEas = -wesEas
				norSou = -norSou
			} else if temp2 == 3 {
				temp = wesEas
				wesEas = norSou
				norSou = -temp
			}
		}
	}
	fmt.Println(abs(actualNorSou) + abs(actualwesEas))
}
