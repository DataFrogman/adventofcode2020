package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func validate(nums []int, limiter int) int {
	for x := 0 + limiter; x < len(nums); x++ {
		flag := false
		for y := x - limiter; y < x; y++ {
			for z := x - limiter; z < x; z++ {
				if z != y && nums[z]+nums[y] == nums[x] {
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
		if !flag {
			return x
		}
	}
	return 0
}

func weakness(nums []int, target int) int {
	var beginning int = 0
	for x := 0; nums[x] < target; x++ {
		beginning = nums[x]
		var total int = beginning
		var curr int = x
		var biggest int = beginning
		for total < (target + 1) {
			curr++
			total += nums[curr]
			if nums[curr] > biggest {
				biggest = nums[curr]
			}
			if total == target {
				return (beginning + biggest)
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums []int

	for scanner.Scan() {
		var line string = scanner.Text()
		temp2 := strings.ReplaceAll(line, "\n", "")
		temp, _ := strconv.Atoi(temp2)
		nums = append(nums, temp)
	}

	var limiter int = 25
	erro := nums[validate(nums, limiter)]
	println(weakness(nums, erro))
}
