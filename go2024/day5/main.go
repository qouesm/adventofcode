package main

import (
	"aoc2024/utils"
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
	fmt.Printf("5a: %v\n", part1(data))
	fmt.Printf("5b: %v\n", part2(data))
}

func part1(data []byte) int {
	defer utils.Timer()()

	rules, updates := parseData(data)
	var answer int
	for _, update := range updates {
		if isValid, _ := isValidUpdate(update, rules); isValid {
			answer += update[len(update)/2]
		}
	}

	return answer
}

func part2(data []byte) int {
	defer utils.Timer()()

	rules, updates := parseData(data)
	invalidUpdates := make([][]int, 0)
	for _, update := range updates {
		if isValid, _ := isValidUpdate(update, rules); !isValid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	answer := 0
	for _, update := range invalidUpdates {
		for {
			if isValid, pageIndex := isValidUpdate(update, rules); isValid {
				answer += update[len(update)/2]
				break
			} else {
				// move element at pageIndex to start
				update = append(append([]int{update[pageIndex]}, update[:pageIndex]...), update[pageIndex+1:]...)
			}
		}
	}

	return answer
}

func parseData(data []byte) (rules map[int][]int, updates [][]int) {
	dataParts := strings.Split(string(data), "\n\n")

	dataRules := strings.Split(dataParts[0], "\n")
	rules = make(map[int][]int)
	for _, v := range dataRules {
		pages := strings.Split(v, "|")
		before, _ := strconv.Atoi(pages[0])
		after, _ := strconv.Atoi(pages[1])

		r, exists := rules[before]
		if !exists {
			r = []int{}
		}
		r = append(r, after)
		rules[before] = r
	}

	dataUpdates := strings.Split(dataParts[1][:len(dataParts[1])-1], "\n")
	updates = make([][]int, len(dataUpdates))
	for i, v := range dataUpdates {
		pages := strings.Split(v, ",")
		for _, page := range pages {
			n, _ := strconv.Atoi(page)
			updates[i] = append(updates[i], n)
		}
	}
	return
}

func isValidUpdate(update []int, rules map[int][]int) (isValid bool, invalidPageIndex int) {
	for i, page := range update[1:] {
		for _, p := range update[:i+1] {
			for _, r := range rules[page] {
				if p == r {
					return false, i + 1
				}
			}
		}
	}
	return true, 0
}
