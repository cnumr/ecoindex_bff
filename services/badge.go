package services

func GetColor(grade string) string {
	colors := map[string]string{
		"A": "#349A47",
		"B": "#51B84B",
		"C": "#CADB2A",
		"D": "#F6EB15",
		"E": "#FECD06",
		"F": "#F99839",
		"G": "#ED2124",
	}

	if color, ok := colors[grade]; ok {
		return color
	}

	return "lightgrey"
}
