package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) [][]int64 {
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
	var lines [][]int64
	for scanner.Scan() {
		lineInt := stringToBitArray(scanner)
		lines = append(lines, lineInt)
	}
	return lines
}

func stringToBitArray(scanner *bufio.Scanner) []int64 {
	var bitArray []int64
	for _, bit  := range strings.Split(scanner.Text(), "") {
		value, err := strconv.ParseInt(bit, 0, 64)

		if err != nil {
			panic(err)
		}

		bitArray = append(bitArray, value)
	}
	return bitArray
}

func binaryToDec(bitArray []int64) int64 {
	var dec int64
	for i := 0; i < len(bitArray); i++ {
		dec += int64(bitArray[i]) * int64(math.Pow(float64(2),  float64(len(bitArray) - 1 - i)))
	}
	return dec
}

func getMostAndLeastCommon(bitArray []int64) (mostCommon int64, leastCommon int64){
	oneCount, zeroCount := 0, 0

	for _, bit := range bitArray {
		switch bit {
		case 0:
			zeroCount++
		case 1:
			oneCount++
		}
	}

	if oneCount >= zeroCount {
		mostCommon = 1
		leastCommon = 0
	} else {
		mostCommon = 0
		leastCommon = 1
	}
	return mostCommon, leastCommon
}

