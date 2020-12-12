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
	file, err := os.Open("message.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var directions []string = make([]string, 0)
	var norSou int = 0
	var wesEas int = 0
	var dir string = "E"
	cards := []string{"N", "E", "S", "W"}
	cardsBack := []string{"W", "S", "E", "N"}

	for scanner.Scan() {
		var line string = scanner.Text()
		line = strings.ReplaceAll(line, "\n", "")
		directions = append(directions, line)
	}
	for _, v := range directions {
		switch v[0] {
		case 0x4e:
			temp, _ := strconv.Atoi(v[1:])
			norSou += temp
		case 0x53:
			temp, _ := strconv.Atoi(v[1:])
			norSou -= temp
		case 0x45:
			temp, _ := strconv.Atoi(v[1:])
			wesEas -= temp
		case 0x57:
			temp, _ := strconv.Atoi(v[1:])
			wesEas += temp
		case 0x46:
			switch dir {
			case "N":
				temp, _ := strconv.Atoi(v[1:])
				norSou += temp
			case "S":
				temp, _ := strconv.Atoi(v[1:])
				norSou -= temp
			case "W":
				temp, _ := strconv.Atoi(v[1:])
				wesEas += temp
			case "E":
				temp, _ := strconv.Atoi(v[1:])
				wesEas -= temp
			}
		case 0x52:
			temp, _ := strconv.Atoi(v[1:])
			temp2 := (temp / 90) % 4
			temp2 = (temp2 + locationFetcher(cards, dir)) % 4
			dir = cards[temp2]
		case 0x4c:
			temp, _ := strconv.Atoi(v[1:])
			temp2 := (temp / 90) % 4
			temp2 = (temp2 + locationFetcher(cardsBack, dir)) % 4
			dir = cardsBack[temp2]
		}
	}
	fmt.Println(abs(norSou) + abs(wesEas))
}
