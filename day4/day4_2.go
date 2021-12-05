package main

import (
	"fmt"
)

func day4_2()  {
	boards, numbers := readFile("bingo.txt")

	var lastNumber = int64(-1)
	var lastToWin [][]int64
	var completed = map[int]bool{}

	for _, number := range numbers {
		for i, board := range boards {
			if !completed[i] {
				markNumber(board, number)
				if checkVertical(board) || checkHorizontal(board) {
					lastNumber = number
					lastToWin = board
					completed[i] = true
				}
			}
		}
	}
	fmt.Printf("%d\n", sumRemaining(lastToWin) * lastNumber)
}


