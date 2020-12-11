package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func seatEmpty(m [][]int, y int, x int) int {
	if x == 0 && y == 0 {
		if (m[0][1] == 0 || m[0][1] == -1) && (m[1][0] == 0 || m[1][0] == -1) && (m[1][1] == 0 || m[1][1] == -1) {
			return 1
		}
		return 0
	} else if x == 0 && y == len(m)-1 {
		if (m[y][1] == 0 || m[y][1] == -1) && (m[y-1][0] == 0 || m[y-1][0] == -1) && (m[y-1][1] == 0 || m[y-1][1] == -1) {
			return 1
		}
		return 0
	} else if x == len(m[y])-1 && y == 0 {
		if (m[0][x-1] == 0 || m[0][x-1] == -1) && (m[1][0] == 0 || m[1][0] == -1) && (m[1][x-1] == 0 || m[1][x-1] == -1) {
			return 1
		}
		return 0
	} else if x == len(m[y])-1 && y == len(m)-1 {
		if (m[y][x-1] == 0 || m[y][x-1] == -1) && (m[y-1][x] == 0 || m[y-1][x] == -1) && (m[y-1][x-1] == 0 || m[y-1][x-1] == -1) {
			return 1
		}
		return 0
	} else if x == 0 {
		if (m[y-1][0] == 0 || m[y-1][0] == -1) && (m[y-1][x+1] == 0 || m[y-1][x+1] == -1) && (m[y][x+1] == 0 || m[y][x+1] == -1) && (m[y+1][x+1] == 0 || m[y+1][x+1] == -1) && (m[y+1][x] == 0 || m[y+1][x] == -1) {
			return 1
		}
		return 0
	} else if x == len(m[y])-1 {
		if (m[y-1][x] == 0 || m[y-1][x] == -1) && (m[y-1][x-1] == 0 || m[y-1][x-1] == -1) && (m[y][x-1] == 0 || m[y][x-1] == -1) && (m[y+1][x-1] == 0 || m[y+1][x-1] == -1) && (m[y+1][x] == 0 || m[y+1][x] == -1) {
			return 1
		}
		return 0
	} else if y == 0 {
		if (m[0][x-1] == 0 || m[0][x-1] == -1) && (m[y+1][x-1] == 0 || m[y+1][x-1] == -1) && (m[y+1][x] == 0 || m[y+1][x] == -1) && (m[y+1][x+1] == 0 || m[y+1][x+1] == -1) && (m[y][x+1] == 0 || m[y][x+1] == -1) {
			return 1
		}
		return 0
	} else if y == len(m)-1 {
		if (m[y][x-1] == 0 || m[y][x-1] == -1) && (m[y-1][x-1] == 0 || m[y-1][x-1] == -1) && (m[y-1][x] == 0 || m[y-1][x] == -1) && (m[y-1][x+1] == 0 || m[y-1][x+1] == -1) && (m[y][x+1] == 0 || m[y][x+1] == -1) {
			return 1
		}
		return 0
	} else {
		if (m[y-1][x-1] == 0 || m[y-1][x-1] == -1) && (m[y-1][x] == 0 || m[y-1][x] == -1) && (m[y-1][x+1] == 0 || m[y-1][x+1] == -1) && (m[y][x+1] == 0 || m[y][x+1] == -1) && (m[y+1][x+1] == 0 || m[y+1][x+1] == -1) && (m[y+1][x] == 0 || m[y+1][x] == -1) && (m[y+1][x-1] == 0 || m[y+1][x-1] == -1) && (m[y][x-1] == 0 || m[y][x-1] == -1) {
			return 1
		}
		return 0
	}
}

func seatFilled(m [][]int, y int, x int) int {
	var acc int = 0
	if x == 0 && y == 0 {
		if m[0][1] == 1 {
			acc++
		}
		if m[1][0] == 1 {
			acc++
		}
		if m[1][1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if x == 0 && y == len(m)-1 {
		if m[y][1] == 1 {
			acc++
		}
		if m[y-1][0] == 1 {
			acc++
		}
		if m[y-1][1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if x == len(m[y])-1 && y == 0 {
		if m[0][x-1] == 1 {
			acc++
		}
		if m[1][0] == 1 {
			acc++
		}
		if m[1][x-1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if x == len(m[y])-1 && y == len(m)-1 {
		if m[y][x-1] == 1 {
			acc++
		}
		if m[y-1][x] == 1 {
			acc++
		}
		if m[y-1][x-1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if x == 0 {
		if m[y-1][0] == 1 {
			acc++
		}
		if m[y-1][x+1] == 1 {
			acc++
		}
		if m[y][x+1] == 1 {
			acc++
		}
		if m[y+1][x+1] == 1 {
			acc++
		}
		if m[y+1][x] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if x == len(m[y])-1 {
		if m[y-1][x] == 1 {
			acc++
		}
		if m[y-1][x-1] == 1 {
			acc++
		}
		if m[y][x-1] == 1 {
			acc++
		}
		if m[y+1][x-1] == 1 {
			acc++
		}
		if m[y+1][x] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if y == 0 {
		if m[0][x-1] == 1 {
			acc++
		}
		if m[y+1][x-1] == 1 {
			acc++
		}
		if m[y+1][x] == 1 {
			acc++
		}
		if m[y+1][x+1] == 1 {
			acc++
		}
		if m[y][x+1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else if y == len(m)-1 {
		if m[y][x-1] == 1 {
			acc++
		}
		if m[y-1][x-1] == 1 {
			acc++
		}
		if m[y-1][x] == 1 {
			acc++
		}
		if m[y-1][x+1] == 1 {
			acc++
		}
		if m[y][x+1] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	} else {
		if m[y-1][x-1] == 1 {
			acc++
		}
		if m[y-1][x] == 1 {
			acc++
		}
		if m[y-1][x+1] == 1 {
			acc++
		}
		if m[y][x-1] == 1 {
			acc++
		}
		if m[y+1][x+1] == 1 {
			acc++
		}
		if m[y][x+1] == 1 {
			acc++
		}
		if m[y+1][x-1] == 1 {
			acc++
		}
		if m[y+1][x] == 1 {
			acc++
		}
		if acc >= 4 {
			return 0
		}
		return 1
	}
}

func checker(m [][]int, y int, x int) int {
	if m[y][x] == -1 {
		return -1
	} else if m[y][x] == 0 {
		return seatEmpty(m, y, x)
	}
	return seatFilled(m, y, x)
}

func gameOfLife(m [][]int) [][]int {
	var tempMap [][]int
	tempMap = make([][]int, 0)
	for y, val := range m {
		temp := make([]int, 0)
		for x := range val {
			temp = append(temp, checker(m, y, x))
		}
		tempMap = append(tempMap, temp)
	}
	return tempMap
}

func main() {
	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var m [][]int
	m = make([][]int, 0)

	for scanner.Scan() {
		var line string = scanner.Text()
		temp := make([]int, 0)
		for _, v := range line {
			if v == '.' {
				temp = append(temp, -1)
			} else if v == 'L' {
				temp = append(temp, 0)
			} else if v == '#' {
				temp = append(temp, 1)
			}
		}
		m = append(m, temp)
	}

	var changed bool = true
	for changed {
		var tempMap [][]int
		tempMap = gameOfLife(m)
		tempAcc := 0
		tempAcc2 := 0
		for i, v := range m {
			for i2, v2 := range v {
				tempAcc2++
				if v2 != tempMap[i][i2] {
					changed = true
				} else {
					tempAcc++
				}
			}
		}
		if tempAcc == tempAcc2 {
			changed = false
		}
		m = make([][]int, 0)
		for _, val := range tempMap {
			temp := make([]int, 0)
			for _, val2 := range val {
				temp = append(temp, val2)
			}
			m = append(m, temp)
		}
		for _, v := range m {
			for _, v2 := range v {
				if v2 == 1 {
					fmt.Print("#")
				} else if v2 == 0 {
					fmt.Print("L")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println(m)
	var acc int = 0
	for _, v := range m {
		for _, v2 := range v {
			if v2 == 1 {
				acc++
			}
		}
	}
	println(acc)
}
