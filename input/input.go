package input

import (
	"fmt"
	"github.com/diist/toy-robot-golang/robot"
	"regexp"
	"strconv"
	"strings"
)

func PlayWithFile(filename string) {
	var my_robot *robot.Robot
	var err error
	commands := linesFromFile(filename)
	for _, command := range commands {
		my_robot, err = runCommand(command, my_robot)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func runCommand(command string, my_robot *robot.Robot) (*robot.Robot, error) {
	if regexp.MustCompile("PLACE [0-9]+,[0-9]+,[A-Z]+").MatchString(command) {
		return runPlaceCommand(command, my_robot)
	}
	if my_robot == nil {
		return nil, fmt.Errorf("You must first PLACE the robot")
	}

	switch {
	case command == "MOVE":
		return robot.Move(my_robot), nil
	case command == "LEFT":
		return robot.Left(my_robot), nil
	case command == "RIGHT":
		return robot.Right(my_robot), nil
	case command == "REPORT":
		fmt.Println(robot.Report(my_robot))
		return my_robot, nil
	case command == "" || command == "\n":
		return my_robot, nil
	}

	return my_robot, fmt.Errorf("Command not recognized: %s", command)
}

func runPlaceCommand(command string, my_robot *robot.Robot) (*robot.Robot, error) {
	args_string := strings.Fields(command)[1]
	args := strings.Split(args_string, ",")
	x, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	direction := args[2]
	new_robot, err := robot.Place(x, y, direction)
	if err != nil {
		return my_robot, fmt.Errorf("Invalid PLACE command")
	}
	return new_robot, nil
}
