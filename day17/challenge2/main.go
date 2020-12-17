package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//                 w z y x
func evaluate(cube [][][][]int, x int, y int, z int, w int) int {
	acc := 0
	for tempw := -1; tempw < 2; tempw++ {
		for tempz := -1; tempz < 2; tempz++ {
			for tempy := -1; tempy < 2; tempy++ {
				for tempx := -1; tempx < 2; tempx++ {
					if x+tempx < len(cube) && x+tempx >= 0 && y+tempy < len(cube) && y+tempy >= 0 && z+tempz < len(cube) && z+tempz >= 0 && w+tempw < len(cube) && w+tempw >= 0 && !(tempy == 0 && tempz == 0 && tempx == 0 && tempw == 0) {
						acc += cube[w+tempw][z+tempz][y+tempy][x+tempx]
					}
				}
			}
		}
	}
	if (acc == 2 || acc == 3) && (cube)[w][z][y][x] == 1 {
		return 1
	} else if acc == 3 && (cube)[w][z][y][x] == 0 {
		return 1
	}
	return 0
}

func printRow(row []int) {
	for _, v := range row {
		if v == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("#")
		}
	}
}

func printPlane(plane [][]int) {
	for _, vx := range plane {
		for _, v := range vx {
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printCube(cube [][][]int) {
	acc := -len(cube) / 2
	for _, vy := range cube {
		fmt.Printf("z=%d\n", acc)
		for _, vx := range vy {
			for _, v := range vx {
				if v == 0 {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		acc++
	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var baseSize int = 8 + 12 // change to 12 for real run
	var numGens int = 6
	var baseArray [][][][]int = make([][][][]int, baseSize)

	for z := 0; z < baseSize; z++ {
		tempz := make([][][]int, baseSize)
		for y := 0; y < baseSize; y++ {
			tempy := make([][]int, baseSize)
			for x := 0; x < baseSize; x++ {
				tempx := make([]int, baseSize)
				tempy[x] = tempx
			}
			tempz[y] = tempy
		}
		baseArray[z] = tempz
	}
	yacc := 0
	for scanner.Scan() {
		var line string = scanner.Text()
		for i, letter := range line {
			if letter == '.' {
				baseArray[baseSize/2][baseSize/2][yacc+6][i+6] = 0
			} else {
				baseArray[baseSize/2][baseSize/2][yacc+6][i+6] = 1
			}
		}
		yacc++
	}

	for currGen := 1; currGen < numGens+1; currGen++ {
		var tempArray [][][][]int = make([][][][]int, baseSize)
		for z := 0; z < baseSize; z++ {
			tempz := make([][][]int, baseSize)
			for y := 0; y < baseSize; y++ {
				tempy := make([][]int, baseSize)
				for x := 0; x < baseSize; x++ {
					tempx := make([]int, baseSize)
					tempy[x] = tempx
				}
				tempz[y] = tempy
			}
			tempArray[z] = tempz
		}
		for w := 0; w < baseSize; w++ {
			for z := 0; z < baseSize; z++ {
				for y := 0; y < baseSize; y++ {
					for x := 0; x < baseSize; x++ {
						tempArray[w][z][y][x] = evaluate(baseArray, x, y, z, w)
					}
				}
			}
		}
		for w := 0; w < baseSize; w++ {
			for z := 0; z < baseSize; z++ {
				for y := 0; y < baseSize; y++ {
					for x := 0; x < baseSize; x++ {
						baseArray[w][z][y][x] = tempArray[w][z][y][x]
					}
				}
			}
		}
	}

	numActive := 0
	for w := 0; w < baseSize; w++ {
		for z := 0; z < baseSize; z++ {
			for y := 0; y < baseSize; y++ {
				for x := 0; x < baseSize; x++ {
					numActive += baseArray[w][z][y][x]
				}
			}
		}
	}
	fmt.Println(numActive)
}
