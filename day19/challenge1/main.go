package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var rules map[int]string = make(map[int]string)
var memoSet map[int][]string = make(map[int][]string)

func stringSliceContains(slice []string, m string) bool {
	for _, v := range slice {
		if v == m {
			return true
		}
	}
	return false
}

func evalRules(rule int) []string {
	if val, ok := memoSet[rule]; ok {
		return val
	}
	if strings.ContainsAny(rules[rule], "ab") {
		temp := make([]string, 0)
		if strings.Contains(rules[rule], "\"") {
			temp = append(temp, rules[rule][(1):(len(rules[rule])-1)])
		} else {
			temp = append(temp, rules[rule])
		}
		return temp
	} else if strings.Contains(rules[rule], "|") {

	} else {
		temp := make([]string, 0)
		ruleList := strings.Split(rules[rule], " ")
		for _, v := range ruleList {
			if len(temp) == 0 {
				temp3, _ := strconv.Atoi(v)
				temp2 := evalRules(temp3)
				for _, v := range temp2 {
					temp = append(temp, v)
				}
			} else {
				newRules := make([]string, 0)
				for _, v2 := range temp {
					temp3, _ := strconv.Atoi(v)
					temp2 := evalRules(temp3)
					for _, v3 := range temp2 {
						newRules = append(newRules, v2+v3)
					}
				}
				temp = newRules
			}
		}
	}
	fmt.Println(rules[rule])
	return make([]string, 0)
}

func main() {
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var blanklines int = 0
	var acc int = 0
	var messages []string = make([]string, 0)

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) < 2 {
			blanklines++
		}
		if blanklines == 0 {
			temp := strings.Split(line, ": ")
			temp2, _ := strconv.Atoi(temp[0])
			rules[temp2] = temp[1]
		} else {
			messages = append(messages, line)
		}
	}

	for k := range rules {
		memoSet[k] = evalRules(k)
	}
	rulesList := make([]string, 0)
	for _, v := range memoSet {
		for _, v2 := range v {
			rulesList = append(rulesList, v2)
		}
	}
	for _, v := range messages {
		if stringSliceContains(rulesList, v) {
			acc++
		}
	}
	fmt.Println(acc)
}
