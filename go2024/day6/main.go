package main

import (
	"aoc2024/utils"
	"fmt"
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
	fmt.Printf("Xa: %v\n", part1(data))
	fmt.Printf("Xb: %v\n", part2(data))
}

func part1(data []byte) (answer int) {
	defer utils.Timer()()

	nRows := len(strings.Split(string(data), "\n")) - 1
	matrix := make([][]byte, nRows)
	row := 0
	start := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			matrix[row] = data[start:i]
			row++
			start = i + 1
		}
	}

	var y, x int
	for row, line := range matrix {
		for col, v := range line {
			if v == '^' {
				y, x = row, col
				matrix[y][x] = 'X'
			}
		}
	}

	const (
		UP    = 0
		RIGHT = 1
		DOWN  = 2
		LEFT  = 3
	)

	curDir := UP
	inMap := true
	for inMap {
		switch curDir {
		case UP:
			if y <= 0 {
				matrix[y][x] = 'X'
				inMap = false
				break
			}
			switch matrix[y-1][x] {
			case 'X':
				fallthrough
			case '.':
				y--
				matrix[y][x] = 'X'
			case '#':
				curDir = RIGHT
				x++
				matrix[y][x] = 'X'
			default:
				panic("uhoh")
			}

		case RIGHT:
			if x >= len(matrix[0])-1 {
				matrix[y][x] = 'X'
				inMap = false
				break
			}
			switch matrix[y][x+1] {
			case 'X':
				fallthrough
			case '.':
				x++
				matrix[y][x] = 'X'
			case '#':
				curDir = DOWN
				y++
				matrix[y][x] = 'X'
			default:
				panic("uhoh")
			}

		case DOWN:
			if y >= len(matrix)-1 {
				matrix[y][x] = 'X'
				inMap = false
				break
			}
			switch matrix[y+1][x] {
			case 'X':
				fallthrough
			case '.':
				y++
				matrix[y][x] = 'X'
			case '#':
				curDir = LEFT
				x--
				matrix[y][x] = 'X'
			default:
				panic("uhoh")
			}

		case LEFT:
			if x <= 0 {
				matrix[y][x] = 'X'
				inMap = false
				break
			}
			switch matrix[y][x-1] {
			case 'X':
				fallthrough
			case '.':
				x--
				matrix[y][x] = 'X'
			case '#':
				curDir = UP
				y--
				matrix[y][x] = 'X'
			default:
				panic("uhoh")
			}
		default:
			panic("uhoh")
		}
	}

	for _, line := range matrix {
		for _, v := range line {
			if v == 'X' {
				answer++
			}
		}
	}

	return
}

func part2(data []byte) (answer int) {
	defer utils.Timer()()
	return
}
