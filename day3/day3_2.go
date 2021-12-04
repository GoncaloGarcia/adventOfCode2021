package main

import (
	"fmt"
	"reflect"
)

func day3_2() {
	bytes := readFile("bytes.txt")

	remainingBytesGamma := make([][]int64, len(bytes))
	copy(remainingBytesGamma, bytes)

	remainingBytesEpsilon := make([][]int64, len(bytes))
	copy(remainingBytesEpsilon, bytes)

	for i := 0; i < len(bytes[0]); i++ {
		bitSliceGamma := make([]int64, len(remainingBytesGamma))
		bitSliceEpsilon := make([]int64, len(remainingBytesEpsilon))

		for j := 0; j < len(remainingBytesGamma); j++ {
			bitSliceGamma[j] = remainingBytesGamma[j][i]
		}
		for j := 0; j < len(remainingBytesEpsilon); j++ {
			bitSliceEpsilon[j] = remainingBytesEpsilon[j][i]
		}
		mc, _ := getMostAndLeastCommon(bitSliceGamma)
		_, lc := getMostAndLeastCommon(bitSliceEpsilon)

		if(len(remainingBytesGamma) > 1) {
			remainingBytesGamma = findByPrefix(remainingBytesGamma, i, mc)
		}
		if(len(remainingBytesEpsilon) > 1) {
			remainingBytesEpsilon = findByPrefix(remainingBytesEpsilon, i, lc)
		}
	}

	gamma := binaryToDec(remainingBytesGamma[0])
	epsilon := binaryToDec(remainingBytesEpsilon[0])

	fmt.Printf("%d * %d = %d", gamma, epsilon, gamma*epsilon)
}

func findByPrefix(bytes [][]int64, index int, bit int64) [][]int64 {
	remainingBytes := make([][]int64, len(bytes))
	copy(remainingBytes, bytes)

	for _, bitSlice := range bytes {
		if bitSlice[index] != int64(bit) {
			remainingBytes = removeByte(remainingBytes, bitSlice)
		}
	}

	return remainingBytes
}

func removeByte(bytes [][]int64, bitSlice []int64) [][]int64 {
	for i, b := range bytes {
		if reflect.DeepEqual(b, bitSlice) {
			return append(bytes[:i], bytes[i+1:]...)
		}
	}
	return bytes
}
