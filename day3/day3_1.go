package main

import (
	"fmt"
)

func day3_1()  {
	bytes := readFile("bytes.txt")

	mostCommonVals := make([]int64, len(bytes[0]))
	leastCommonVals := make([]int64, len(bytes[0]))

	for i := 0; i < len(bytes[0]); i++ {
		bitSlice := make([]int64, len(bytes))
		for j := 0; j < len(bytes); j++ {
			bitSlice[j] = bytes[j][i]
		}
		mc, lc := getMostAndLeastCommon(bitSlice)
		mostCommonVals[i] = mc
		leastCommonVals[i] = lc
	}

	gamma := binaryToDec(mostCommonVals)
	epsilon := binaryToDec(leastCommonVals)

	fmt.Printf("%d * %d = %d\n", gamma, epsilon, gamma * epsilon)
}

