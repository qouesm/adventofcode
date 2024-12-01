package main

import (
	"aoc2024/utils"
	"fmt"
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
	fmt.Printf("Xa: %v\n", part1())
	fmt.Printf("Xb: %v\n", part2())
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}
