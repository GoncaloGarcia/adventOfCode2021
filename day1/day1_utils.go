package main

import (
	"bufio"
	"os"
	"strconv"
)

func readFile(path string) []int64 {
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
	var lines []int64
	for scanner.Scan() {
		lineInt := stringToInt(scanner)
		lines = append(lines, lineInt)
	}
	return lines
}

func stringToInt(scanner *bufio.Scanner) int64 {
	lineInt, err := strconv.ParseInt(scanner.Text(), 0, 64)

	if err != nil {
		panic(err)
	}

	return lineInt
}

