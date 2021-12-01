package main

func day1_1() {
	distances := readFile("measurements.txt")
	var count int
	for i := 1; i < len(distances); i++ {
		if distances[i] > distances[i - 1] {
			count++
		}
	}
	println(count)
}
