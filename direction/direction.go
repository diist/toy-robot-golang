package direction

func Left(direction string) string {
	var new_direction string
	switch direction {
	case "NORTH":
		new_direction = "WEST"
	case "EAST":
		new_direction = "NORTH"
	case "SOUTH":
		new_direction = "EAST"
	case "WEST":
		new_direction = "SOUTH"
	}
	return new_direction
}

func Right(direction string) string {
	var new_direction string
	switch direction {
	case "NORTH":
		new_direction = "EAST"
	case "EAST":
		new_direction = "SOUTH"
	case "SOUTH":
		new_direction = "WEST"
	case "WEST":
		new_direction = "NORTH"
	}
	return new_direction
}

func IsValidDirection(direction string) bool {
	switch direction {
	case
		"NORTH",
		"EAST",
		"SOUTH",
		"WEST":
		return true
	}
	return false
}
