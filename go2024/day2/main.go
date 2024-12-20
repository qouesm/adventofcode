package main

import (
	"aoc2024/utils"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var data []byte

func init() {
	var err error
	data, err = utils.ReadInput()
	if err != nil {
		panic(err)
	}
}

func main() {
	list := makeList(data)
	fmt.Printf("2a: %v\n", part1(list))
	fmt.Printf("2b: %v\n", part2(list))
}

func makeList(data []byte) [][]int {
	listLength := bytes.Count(data, []byte{'\n'})
	list := make([][]int, listLength)
	for i, line := range strings.Split(string(data), "\n") {
		lineStrings := strings.Split(line, " ")
		for _, v := range lineStrings {
			if len(v) == 0 {
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			list[i] = append(list[i], num)
		}
	}
	return list
}

func part1(list [][]int) int {
	defer utils.Timer()()
	var answer int

	for _, report := range list {
		isSafeReport := true
		prevLevel := report[0]
		var reportType int // 0: ?, 1: inc, 2: dec
		for j, level := range report {
			if j == 0 || isSafeReport == false {
				continue
			}
			if j == 1 {
				if level < prevLevel && prevLevel-level <= 3 {
					reportType = 1
				} else if level > prevLevel && level-prevLevel <= 3 {
					reportType = 2
				} else {
					isSafeReport = false
					continue
				}
			}

			switch reportType {
			case 1:
				if level < prevLevel && prevLevel-level <= 3 {
					prevLevel = level
				} else {
					isSafeReport = false
				}
			case 2:
				if level > prevLevel && level-prevLevel <= 3 {
					prevLevel = level
				} else {
					isSafeReport = false
				}
			default:
				return -1
			}
		}
		if isSafeReport {
			answer++
		}
	}

	return answer
}

func part2(list [][]int) int {
	defer utils.Timer()()

	var answer int
	for _, report := range list {
		if isValidReport(report) {
			answer++
			continue
		}

		alreadyValid := false
		for i := range report {
			if alreadyValid {
				continue
			}
			r := make([]int, 0, len(report)-1)
			for j, v := range report {
				if j != i {
					r = append(r, v)
				}
			}
			if isValidReport(r) {
				answer++
				alreadyValid = true
				continue
			}
		}
	}
	return answer
}

func isValidReport(report []int) bool {
	prevLevel := report[0]
	var reportType int // 0: ?, 1: inc, 2: dec
	for i, level := range report {
		if i == 0 {
			continue
		} else if i == 1 {
			if level < prevLevel && prevLevel-level <= 3 {
				reportType = 1
				prevLevel = level
			} else if level > prevLevel && level-prevLevel <= 3 {
				reportType = 2
				prevLevel = level
			} else {
				return false
			}
		} else {
			switch reportType {
			case 1:
				if level < prevLevel && prevLevel-level <= 3 {
					prevLevel = level
				} else {
					return false
				}
			case 2:
				if level > prevLevel && level-prevLevel <= 3 {
					prevLevel = level
				} else {
					return false
				}
			default:
				fmt.Println("uh oh")
				return false
			}
		}
	}
	return true
}
