package robot

import (
	"fmt"
	"github.com/diist/toy-robot-golang/table"
)

type Robot struct {
	X, Y      int
	Direction string
}

func Place(x int, y int, direction string) (*Robot, error) {
	if !isValidPosition(x, y) {
		return nil, fmt.Errorf("invalid position")
	}
	if !isValidDirection(direction) {
		return nil, fmt.Errorf("invalid direction")
	}
	return &Robot{x, y, direction}, nil
}

func Move(robot *Robot) *Robot {
	x, y, direction := robot.X, robot.Y, robot.Direction

	switch direction {
	case "NORTH":
		y += 1
	case "EAST":
		x += 1
	case "SOUTH":
		y -= 1
	case "WEST":
		x -= 1
	}
	if !isValidPosition(x, y) {
		return robot
	}
	return &Robot{x, y, direction}
}

func isValidPosition(x int, y int) bool {
	table := table.NewTable()
	if x > table.MaxX || x < 0 || y > table.MaxY || y < 0 {
		return false
	}
	return true
}

func isValidDirection(direction string) bool {
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
