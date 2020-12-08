package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func inc(val int, acc int) int {
	return acc + val
}

func jump(val int, curr int) int {
	temp := curr + val
	return temp
}

func nop(curr int) int {
	curr++
	return curr
}

func test(program []string) int {
	var curr int = 0
	var acc int = 0
	var lastJump int = 0
	run := make(map[int]int)

	for {
		//
		if curr == len(program) {
			println(acc)
			os.Exit(1)
		}
		_, ok := run[curr]
		if ok {
			print("CORRUPTED FOUND: ")
			println(curr - lastJump)
			return 0
			/*
				if program[curr-lastJump][:3] == "jmp" {
					println("was jump")
					println(lastJump)
					curr = curr - lastJump
					run[curr]++
					curr = nop(curr)
					println(program[curr])
					lastJump = 1
				} else {
					curr = curr - lastJump
					run[curr]++
					temp := strings.ReplaceAll(program[curr][3:], " ", "")
					temp = strings.ReplaceAll(program[curr][3:], "+", "")
					temp2, _ := strconv.Atoi(temp[1:])
					lastJump = temp2
					curr = jump(temp2, curr)
				}
			*/
		}
		switch program[curr][:3] {
		case "nop":
			run[curr]++
			curr = nop(curr)
			lastJump = 1
		case "acc":
			temp := strings.ReplaceAll(program[curr][3:], " ", "")
			temp = strings.ReplaceAll(program[curr][3:], "+", "")
			temp2, _ := strconv.Atoi(temp[1:])
			acc = inc(temp2, acc)
			run[curr]++
			curr++
			lastJump = 1
		case "jmp":
			run[curr]++
			temp := strings.ReplaceAll(program[curr][3:], " ", "")
			temp = strings.ReplaceAll(program[curr][3:], "+", "")
			temp2, _ := strconv.Atoi(temp[1:])
			lastJump = temp2
			curr = jump(temp2, curr)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var program []string

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 2 {
			program = append(program, strings.ReplaceAll(line, "\n", ""))
		}
	}

	for i, command := range program {
		if command[:3] == "jmp" {
			temp := command
			temp = strings.ReplaceAll(temp, "jmp", "nop")
			temp2 := make([]string, len(program))
			copy(temp2, program)
			temp2[i] = temp
			println(i)
			test(temp2)

		} else if command[:3] == "nop" {
			temp := command
			temp = strings.ReplaceAll(temp, "nop", "jmp")
			temp2 := make([]string, len(program))
			copy(temp2, program)
			temp2[i] = temp
			println(i)
			test(temp2)
		}
	}
}
