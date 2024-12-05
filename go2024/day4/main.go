package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
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
	fmt.Printf("4a: %v\n", part1(data))
	fmt.Printf("4b: %v\n", part2(data))
}

// average rotation enjoyer
func part1(data []byte) int {
	defer utils.Timer()()
	var answer int

	// horizontal
	answer += countMatches(data)

	// vertical
	// rotate string by 90deg
	data90 := rotate90(data)
	answer += countMatches(data90)

	// diagonal \
	data45 := rotate45(data)
	answer += countMatches(data45)

	// diagonal /
	/*
		  3
		 2 6
		1 5 9
		 4 8
		  7
	*/
	// rotate 90 by 45
	data135 := rotate45(data90)
	answer += countMatches(data135)

	return answer
}

func part2(data []byte) int {
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

	var answer int

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[0])-1; j++ {
			if matrix[i][j] != 'A' {
				continue
			}

			topLeft := matrix[i-1][j-1]
			topRight := matrix[i-1][j+1]
			bottomLeft := matrix[i+1][j-1]
			bottomRight := matrix[i+1][j+1]

			if topLeft == 'M' {
				if bottomRight != 'S' {
					continue
				}
			} else if topLeft == 'S' {
				if bottomRight != 'M' {
					continue
				}
			} else {
				continue
			}

			if topRight == 'M' {
				if bottomLeft != 'S' {
					continue
				}
			} else if topRight == 'S' {
				if bottomLeft != 'M' {
					continue
				}
			} else {
				continue
			}

			answer++
		}
	}

	return answer
}

func countMatches(data []byte) int {
	var count int
	re := regexp.MustCompile(`(XMAS|SAMX)`)
	curIndex := 0
	for {
		match := re.FindStringIndex(string(data[curIndex:]))
		if match == nil {
			break
		}
		curIndex += match[0] + 1
		count++
	}
	return count
}

func rotate90(data []byte) []byte {
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
	// matrix[row] = data[start:]

	matrix90 := make([][]byte, len(matrix[0]))
	for i := range matrix90 {
		matrix90[i] = make([]byte, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix90[j][len(matrix)-1-i] = matrix[i][j]
		}
	}

	var result []byte
	for i := range matrix90 {
		result = append(result, matrix90[i]...)
		if i < len(matrix90)-1 {
			result = append(result, '\n')
		}
	}

	result = append(result, byte('\n'))
	return result
}

func rotate45(data []byte) []byte {
	/*
			1
		 2 4
		3 5 7
		 6 8
			9
	*/

	nRows := len(strings.Split(string(data), "\n"))
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
	matrix45 := make([][]byte, len(matrix[0])*2-1)
	for i := range matrix45 {
		matrix45[i] = make([]byte, len(matrix))
	}

	cur := 0
	// left side
	for i := 0; i < len(matrix[0]); i++ {
		for j := i; j >= 0; j-- {
			a := j
			b := i - j
			matrix45[cur] = append(matrix45[i], matrix[a][b])
		}
		cur++
	}

	// bottom side
	for i := 1; i < len(matrix); i++ { // (j, i)
		for j := len(matrix[0]) - 1; j >= i; j-- {
			a := j
			b := len(matrix[0]) - 1 - j + i
			matrix45[cur] = append(matrix45[cur], matrix[a][b])
		}
		cur++
	}

	var data45 []byte
	for i := range matrix45 {
		data45 = append(data45, matrix45[i]...)
		if i < len(matrix45)-1 {
			data45 = append(data45, '\n')
		}
	}
	return data45
}
