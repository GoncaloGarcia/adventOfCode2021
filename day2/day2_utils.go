package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	movementType string
	distance int64
}

func readFile(path string) []Movement {
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
	var lines []Movement
	for scanner.Scan() {
		lineInt := stringToMovement(scanner)
		lines = append(lines, lineInt)
	}
	return lines
}

func stringToMovement(scanner *bufio.Scanner) Movement {
	line := strings.Split(scanner.Text(), " ")
	movementType := line[0]
	distance, err := strconv.ParseInt(line[1], 0, 64)

	if err != nil {
		panic(err)
	}

	return Movement{movementType: movementType, distance: distance}
}

