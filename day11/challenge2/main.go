package main

import (
	"bufio"
	"log"
	"os"
)

func seatEmpty2(m [][]int, y int, x int) int {
	if x == 0 && y == 0 {
		return 1
	} else if x == 0 && y == len(m)-1 {
		return 1
	} else if x == len(m[0])-1 && y == 0 {
		return 1
	} else if x == len(m[0])-1 && y == len(m)-1 {
		return 1
	} else if x == 0 {
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][0] == 1 {
				return 0
			} else if m[temp][0] == 0 {
				break
			}
		}
		for temp := 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}

		inc := 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		return 1
	} else if x == len(m[0])-1 {
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := len(m[0]) - 2; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		return 1
	} else if y == 0 {
		for temp := 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		return 1
	} else if y == len(m)-1 {
		for temp := len(m) - 2; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		return 1
	} else {
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				return 0
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				return 0
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				return 0
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		return 1
	}
}

func seatFilled2(m [][]int, y int, x int) int {
	if x == 0 && y == 0 {
		return 1
	} else if x == 0 && y == len(m)-1 {
		return 1
	} else if x == len(m[0])-1 && y == 0 {
		return 1
	} else if x == len(m[0])-1 && y == len(m)-1 {
		return 1
	} else if x == 0 {
		acc := 0
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][0] == 1 {
				acc++
				break
			} else if m[temp][0] == 0 {
				break
			}
		}
		for temp := 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}

		inc := 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		if acc >= 5 {
			return 0
		}
		return 1
	} else if x == len(m[0])-1 {
		acc := 0
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := len(m[0]) - 2; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		if acc >= 5 {
			return 0
		}
		return 1
	} else if y == 0 {
		acc := 0
		for temp := 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		if acc >= 5 {
			return 0
		}
		return 1
	} else if y == len(m)-1 {
		acc := 0
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		if acc >= 5 {
			return 0
		}
		return 1
	} else {
		acc := 0
		for temp := y - 1; temp >= 0; temp-- {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x - 1; temp >= 0; temp-- {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		for temp := y + 1; temp < len(m); temp++ {
			if m[temp][x] == 1 {
				acc++
				break
			} else if m[temp][x] == 0 {
				break
			}
		}
		for temp := x + 1; temp < len(m[0]); temp++ {
			if m[y][temp] == 1 {
				acc++
				break
			} else if m[y][temp] == 0 {
				break
			}
		}
		inc := 1
		for {
			tempy := y - inc
			tempx := x + inc
			if tempy < 0 {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y - inc
			tempx := x - inc
			if tempy < 0 {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x + inc
			if tempy >= len(m) {
				break
			}
			if tempx >= len(m[0]) {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		inc = 1
		for {
			tempy := y + inc
			tempx := x - inc
			if tempy >= len(m) {
				break
			}
			if tempx < 0 {
				break
			}
			if m[tempy][tempx] == 1 {
				acc++
				break
			} else if m[tempy][tempx] == 0 {
				break
			}
			inc++
		}
		if acc >= 5 {
			return 0
		}
		return 1
	}
}

func checker(m [][]int, y int, x int) int {
	if m[y][x] == -1 {
		return -1
	} else if m[y][x] == 0 {
		return seatEmpty2(m, y, x)
	}
	return seatFilled2(m, y, x)
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
	// for _, v := range tempMap {
	// 	for _, v2 := range v {
	// 		if v2 == 1 {
	// 			fmt.Print("#")
	// 		} else if v2 == 0 {
	// 			fmt.Print("L")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()
	return tempMap
}

func main() {
	file, err := os.Open("input.txt")

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
	}
	// for _, v := range m {
	// 	for _, v2 := range v {
	// 		if v2 == 1 {
	// 			fmt.Print("#")
	// 		} else if v2 == 0 {
	// 			fmt.Print("L")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

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
