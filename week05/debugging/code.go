package debugging

func letterGrade(score int) string {
	g := score / 10
	switch g {
	case 9, 10:
		return "A"
	case 8:
		return "B"
	case 7:
		return "C"
	case 6:
		return "D"
	}
	return "F"
}

func Sum(values []int) int {
	total := 0

	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total
}
