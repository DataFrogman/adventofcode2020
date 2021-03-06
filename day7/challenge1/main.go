package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//Bag ...
type Bag struct {
	name     string
	children map[string]int
	parents  map[string]int
}

func addParent(parent *Bag, child *Bag) {
	(*child).parents[(*parent).name]++
}

func addChild(parent *Bag, child *Bag) {
	(*parent).children[(*child).name]++
}

func newBag(name string) *Bag {
	b := Bag{name: name}
	b.children = make(map[string]int)
	b.parents = make(map[string]int)
	return &b
}

func recParRet(m map[string]*Bag, curr string, counted map[string]int) int {
	if len((*(m[curr])).parents) == 0 {
		return 0
	} else {
		temp := 0
		for key := range (*(m[curr])).parents {
			_, ok := counted[key]
			if !ok {
				temp++
				counted[key]++
				temp += recParRet(m, key, counted)
			}
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

	m := make(map[string][]string)

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 2 {
			temp := strings.Split(line, " bags contain")
			tempBagName := temp[0]
			tempChildrenList := make([]string, 0)
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
					tempChildrenList = append(tempChildrenList, temp2)
				}
				m[tempBagName] = tempChildrenList
			}
		}
	}

	m2 := make(map[string]*Bag)

	for key, value := range m {
		temp := newBag(key)
		for key2, value2 := range m {
			if contains(value2, key) {
				(*temp).parents[key2]++
			}
		}
		for _, value3 := range value {
			(*temp).children[value3]++
		}
		m2[key] = temp
	}
	m3 := make(map[string]int)
	println(recParRet(m2, "shiny gold", m3))

}
