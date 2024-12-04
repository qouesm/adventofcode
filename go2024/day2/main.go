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
				if level < prevLevel {
					reportType = 1
				} else if level > prevLevel {
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
	return 0
}
