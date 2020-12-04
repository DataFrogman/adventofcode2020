package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var acc = 0

	var byr bool = false
	var iyr bool = false
	var eyr bool = false
	var hgt bool = false
	var hcl bool = false
	var ecl bool = false
	var pid bool = false
	var broke bool = false
	//var cid bool = false

	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) < 2 {
			if byr && iyr && eyr && hgt && hcl && ecl && pid && !broke {
				acc++
				//println("valid")
			} else {
				//println("invalid")
			}

			//println()

			byr = false
			iyr = false
			eyr = false
			hgt = false
			hcl = false
			ecl = false
			pid = false
			broke = false

		} else {
			s := strings.Split(line, " ")
			for _, v := range s {
				if len(v) > 0 {
					q := v[0:3]
					//println(v)
					switch q {
					case "byr":
						temp, _ := strconv.Atoi(v[4:])
						if 1920 <= temp && temp <= 2002 && len(v[4:]) == 4 {
							byr = true
							//println(true)
						} else {
							//println(false)
						}
					case "iyr":
						temp, _ := strconv.Atoi(v[4:])
						if 2010 <= temp && temp <= 2020 && len(v[4:]) == 4 {
							iyr = true
							//println(true)
						} else {
							//println(false)
						}
					case "eyr":
						temp, _ := strconv.Atoi(v[4:])
						if 2020 <= temp && temp <= 2030 && len(v[4:]) == 4 {
							eyr = true
							//println(true)
						} else {
							//println(false)
						}
					case "hgt":
						//println(v)
						if v[(len(v)-2):] == "cm" {
							temp, _ := strconv.Atoi(v[4:(len(v) - 2)])
							//println(temp)
							if 150 <= temp && temp <= 193 {
								hgt = true
								//println(true)
							} else {
								//println(false)
							}
						} else if v[(len(v)-2):] == "in" {
							temp, _ := strconv.Atoi(v[4:(len(v) - 2)])
							//println(temp)
							if 59 <= temp && temp <= 76 {
								hgt = true
								//println(true)
							} else {
								//println(false)
							}
						}
					case "hcl":
						temp := v[4:]
						//println(v)
						//println(temp)
						matched, _ := regexp.MatchString(`#[0-9a-f]{6}`, temp)
						if matched {
							hcl = true
							//println(true)
						} else {
							//println(false)
						}
					case "ecl":
						temp := v[4:]
						//println(v)
						//println(temp)
						if temp == "amb" || temp == "blu" || temp == "brn" || temp == "gry" || temp == "grn" || temp == "hzl" || temp == "oth" {
							ecl = true
							//println(true)
						} else {
							//println(false)
						}
					case "pid":
						temp := v[4:]
						matched, _ := regexp.MatchString(`[0-9]{9}`, temp)
						if matched {
							pid = true
							//println(true)
						} else {
							//println(false)
						}
					case "cid":
						//	cid = true
					default:
						broke = true
					}
				}
			}
		}
	}
	println(acc)
}
