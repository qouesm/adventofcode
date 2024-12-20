package main

import (
	"aoc2024/utils"
	"bytes"
	"fmt"
	"math"
	"slices"
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
	var list [][]int
	list = makeList(data)
	fmt.Printf("1a: %v\n", part1(list))
	fmt.Printf("1b: %v\n", part2(list))
	fmt.Printf("1b: %v\n", part2a(list))
}

func makeList(data []byte) [][]int {
	listLength := bytes.Count(data, []byte{'\n'})
	list := make([][]int, 2)
	for i := range list {
		list[i] = make([]int, listLength)
	}
	for i, line := range strings.Split(string(data), "\n") {
		v := strings.Split(line, "   ")
		if len(v) != 2 {
			continue
		}

		v0, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		list[0][i] = int(float64(v0))

		v1, err := strconv.Atoi(v[1])
		if err != nil {
			panic(err)
		}
		list[1][i] = int(float64(v1))
	}
	return list
}

func part1(list [][]int) int {
	defer utils.Timer()()
	var answer int
	for _, v := range list {
		slices.Sort(v)
	}

	for i := 0; i < len(list[0]); i++ {
		x := int(math.Abs(float64(list[0][i] - list[1][i])))
		answer += x
		// fmt.Printf("%v: abs(%v - %v) = %v\n", i, list[0][i], list[1][i], x)
	}

	return answer
}

func part2(list [][]int) int {
	defer utils.Timer()()
	var answer int
	for _, v := range list[0] {
		count := 0
		for _, w := range list[1] {
			if v == w {
				count++
			}
		}
		answer += v * count
		// fmt.Printf("%v appears %v times\n", v, count)
	}

	return answer
}

func part2a(list [][]int) int {
	defer utils.Timer()()
	var answer int

	for _, v := range list {
		slices.Sort(v)
	}
	counts := make(map[int]int)

	i := 0
	for _, v := range list[0] {
		if counts[v] != 0 {
			continue
		}
		c := 0
		for i < len(list[1]) && list[1][i] <= v {
			if list[1][i] == v {
				c++
			}
			i++
		}
		counts[v] = c
	}

	for k, v := range counts {
		answer += k * v
	}

	return answer
}
