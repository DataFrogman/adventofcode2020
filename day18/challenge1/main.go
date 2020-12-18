package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func eval(input string) int {
	//fmt.Println(input)
	var acc int = 0
	var total int = 0
	var flag int = 0
	var tempHold bool = false

	var leftCount int
	//fmt.Println(input)
	for i, v := range input {
		if v == '*' && !tempHold {
			//fmt.Print(" *")
			flag = 1
		} else if v == '+' && !tempHold {
			//fmt.Print(" +")
			flag = 2
		} else if v == '(' {
			tempHold = true
			//fmt.Print(" (")
			//fmt.Printf("{%d}", acc)
			if acc == 0 {
				leftCount = i
			}
			acc++
		} else if v == ')' {
			//fmt.Print(" )")
			acc--
			//fmt.Printf("{%d}", acc)
			if acc == 0 {
				//fmt.Println("recursing")

				//fmt.Printf("[%d]", flag)
				if flag == 1 {
					temp, _ := strconv.Atoi(fmt.Sprint(eval(input[(leftCount + 1):i])))
					total *= temp
					// fmt.Println("multiplying")
					// fmt.Println(total)
				} else if flag == 2 {
					temp, _ := strconv.Atoi(fmt.Sprint(eval(input[(leftCount + 1):i])))
					total += temp
					// fmt.Println("adding")
					// fmt.Println(total)
				} else {
					total, _ = strconv.Atoi(fmt.Sprint(eval(input[(leftCount + 1):i])))
				}
				tempHold = false
				flag = 0
			}
		} else if !tempHold {
			//fmt.Println(tempHold)
			if total == 0 {
				total, _ = strconv.Atoi(string(v))
				//fmt.Println("engaging total")
			}
			if flag == 1 {
				temp, _ := strconv.Atoi(string(v))
				//fmt.Printf(" %d\n", temp)
				total *= temp
			} else if flag == 2 {
				temp, _ := strconv.Atoi(string(v))
				//fmt.Printf(" %d\n", temp)
				total += temp
			}

			flag = 0
		}
	}
	//fmt.Println(total)
	return total
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	acc := 0

	for scanner.Scan() {
		var line string = scanner.Text()
		fmt.Print(line + " = ")
		temp := eval(strings.ReplaceAll(line, " ", ""))
		fmt.Println(temp)
		acc += temp
	}
	//fmt.Println()
	fmt.Println(acc)
}
