package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//Bag ...
type Bag struct {
	name     string
	children map[string]int
	parents  map[string]int
}

func newBag(name string) *Bag {
	b := Bag{name: name}
	b.children = make(map[string]int)
	b.parents = make(map[string]int)
	return &b
}

func recParRet(m map[string]*Bag, curr string) int {
	if len(((*(m[curr])).children)) == 0 {
		return 0
	} else {
		temp := 0
		for key, value := range (*(m[curr])).children {
			temp += value*recParRet(m, key) + value
		}
		return temp
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make(map[string]map[string]int)

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 2 {
			temp := strings.Split(line, " bags contain")
			tempBagName := temp[0]
			tempChildrenList := make(map[string]int)
			if temp[1] == " no other bags." {
				m[tempBagName] = nil
			} else {
				temp := temp[1]
				temp = strings.ReplaceAll(temp, " bags", "")
				temp = strings.ReplaceAll(temp, " bag", "")
				temp = strings.ReplaceAll(temp, " bags.", "")
				temp = strings.ReplaceAll(temp, " bag.", "")
				tempList := strings.Split(temp, ",")
				for _, value := range tempList {
					temp2 := strings.ReplaceAll(value[3:], ".", "")
					temp3 := value[:3]
					temp3 = strings.ReplaceAll(temp3, " ", "")
					i, _ := strconv.Atoi(temp3)
					tempChildrenList[temp2] = i
				}
				m[tempBagName] = tempChildrenList
			}
		}
	}

	m2 := make(map[string]*Bag)

	for key, value := range m {
		temp := newBag(key)
		for key2, value3 := range value {
			(*temp).children[key2] = value3
		}
		m2[key] = temp
	}

	temp := recParRet(m2, "shiny gold")
	println(temp)
}
