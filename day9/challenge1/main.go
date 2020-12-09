package main

import (
	"bufio"
	"fmt"
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
				fmt.Printf("x: %d, y: %d, z: %d\n", nums[x], nums[y], nums[z])
				if z != y && nums[z]+nums[y] == nums[x] {
					flag = true
					fmt.Printf("VALID x: %d, y: %d, z: %d\n", nums[x], nums[y], nums[z])
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
	println(nums[validate(nums, limiter)])
}
