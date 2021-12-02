package main

type Position struct {
	depth int64
	horizontal int64
}

func day2_1() {
	movements := readFile("positions.txt")
	pos := Position{}

	for _, m := range movements {
		switch m.movementType {
		case "forward":
			pos.horizontal += m.distance
		case "down":
			pos.depth += m.distance
		case "up":
			pos.depth -= m.distance
		}
	}
	println(pos.depth * pos.horizontal)
}
