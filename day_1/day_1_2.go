package main

func day1_2() {
	distances := readFile("measurements.txt")
	var count int = -1
	var prevDist int64 = 0
	for i := 2; i < len(distances); i++ {
		currDist := distances[i] + distances[i - 1] + distances[i - 2]
		if currDist > prevDist {
			count++
		}
		prevDist = currDist
	}
	println(count)
}

