package main

import (
	"fmt"
)

func day4_1()  {
	boards, numbers := readFile("bingo.txt")

	for _, number := range numbers {
		for _, board := range boards {
			markNumber(board, number)
			if checkVertical(board) || checkHorizontal(board) {
				fmt.Printf("%d\n", sumRemaining(board) * number)
				return
			}
		}
	}
}

