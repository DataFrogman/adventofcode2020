package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []int

const (
	add = iota - 1
	mul
	openParen
	closeParen
)

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(num int) {
	*s = append(*s, num)
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func eval2(stack *Stack, op int, num int) {
	if op == add {
		temp, _ := stack.pop()
		stack.push(num + temp)
	} else {
		stack.push(num)
	}
}

func eval(input string) int {
	var stack *Stack = &Stack{}

	var total int
	op := mul

	for x := 0; x < len(input); x++ {
		curr := input[x]

		if curr == '+' {
			eval2(stack, op, total)
			op = add
		} else if curr == '*' {
			eval2(stack, op, total)
			op = mul
		} else if curr == '(' {
			stack.push(op)
			op = mul
		} else if curr == ')' {
			eval2(stack, op, total)
			total = 1
			for !stack.isEmpty() && ((*stack)[len(*stack)-1] != add && (*stack)[len(*stack)-1] != mul) {
				temp, _ := stack.pop()
				total *= temp
			}
			temp, _ := stack.pop()
			op = temp
		} else {
			total, _ = strconv.Atoi(string(curr))
		}

		if x == len(input)-1 {
			eval2(stack, op, total)
		}
	}
	total = 1
	for !stack.isEmpty() {
		temp, _ := stack.pop()
		total *= temp
	}
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
		temp := eval(strings.ReplaceAll(line, " ", ""))
		acc += temp
	}
	//fmt.Println()
	fmt.Println(acc)
}
