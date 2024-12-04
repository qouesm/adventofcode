package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
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
	fmt.Printf("3a: %v\n", part1(data))
	fmt.Printf("3b: %v\n", part2(data))
}

func part1(data []byte) int {
	defer utils.Timer()()
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string(data), -1)
	re = regexp.MustCompile(`[a-z()]`)
	for i, v := range matches {
		matches[i] = re.ReplaceAllString(v, "")
	}
	var answer int
	for _, v := range matches {
		mulStr := strings.Split(v, ",")
		mulInt := make([]int, 2)
		mulInt[0], _ = strconv.Atoi(mulStr[0])
		mulInt[1], _ = strconv.Atoi(mulStr[1])
		answer += mulInt[0] * mulInt[1]
	}
	return answer
}

func part2(data []byte) int {
	defer utils.Timer()()
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(string(data), -1)
	re = regexp.MustCompile(`[a-z()]`)
	var answer int
	enabled := true
	for _, v := range matches {
		switch v[:3] {
		case "do(":
			enabled = true
		case "don":
			enabled = false
		case "mul":
			if !enabled {
				continue
			}
			w := re.ReplaceAllString(v, "")
			mulStr := strings.Split(w, ",")
			mulInt := make([]int, 2)
			mulInt[0], _ = strconv.Atoi(mulStr[0])
			mulInt[1], _ = strconv.Atoi(mulStr[1])
			answer += mulInt[0] * mulInt[1]
		default:
			fmt.Println("uhoh")
		}
	}
	return answer
}
