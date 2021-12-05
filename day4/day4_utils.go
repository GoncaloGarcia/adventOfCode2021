package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(path string) ([][][]int64, []int64) {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var boards [][][]int64

	scanner.Scan()
	numbers := parseNumbers(strings.Split(scanner.Text(), ","))

	var board [][]int64

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if len(board) > 0 {
				boards = append(boards, board)
				board = nil
			}
			continue
		}
		re := regexp.MustCompile("\\s")
		boardLine := parseNumbers(re.Split(line, 100))
		board = append(board, boardLine)
	}

	boards = append(boards, board)
	return boards, numbers
}

func parseNumbers(numbers []string) []int64 {
	var intNums []int64
	for _, n := range numbers {
		if len(n) > 0 {
			intNums = append(intNums, stringToInt(strings.TrimSpace(n)))
		}
	}
	return intNums
}

func stringToInt(string string) int64 {
	lineInt, err := strconv.ParseInt(string, 0, 64)

	if err != nil {
		panic(err)
	}

	return lineInt
}

func sumRemaining(board [][]int64) int64 {
	sum := int64(0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func markNumber(board [][]int64, number int64)  {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == number {
				board[i][j] = -1
			}
		}
	}
}

func checkHorizontal(board [][]int64) bool {
row:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != -1 {
				continue row
			}
		}
		return true
	}
	return false
}

func checkVertical(board [][]int64) bool {
column:
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[j][i] != -1 {
				continue column
			}
		}
		return true
	}
	return false
}


