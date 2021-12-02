package main

type Position3 struct {
	depth int64
	horizontal int64
	aim int64
}

func day2_2() {
	movements := readFile("positions.txt")
	pos := Position3{}

	for _, m := range movements {
		switch m.movementType {
		case "forward":
			pos.horizontal += m.distance
			pos.depth += pos.aim * m.distance
		case "down":
			pos.aim += m.distance
		case "up":
			pos.aim -= m.distance
		}
	}
	println(pos.depth * pos.horizontal)
}
