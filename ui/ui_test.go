package ui

import (
	"github.com/diist/toy-robot-golang/robot"
	"testing"
)

func TestRunCommand(t *testing.T) {
	command_string := "PLACE 2,2,WEST"
	var result_robot *robot.Robot
	var err error
	result_robot, err = runCommand(command_string, result_robot)

	if err != nil {
		t.Errorf("There should be no error here")
	}
	if robot.Report(result_robot) != "2,2,WEST" {
		t.Errorf("Failed to process PLACE command")
	}
}

func TestRunCommandNotPlaced(t *testing.T) {
	command_string := "MOVE"
	var result_robot *robot.Robot
	var err error
	result_robot, err = runCommand(command_string, result_robot)

	if err == nil {
		t.Errorf("There should be an error here, because the robot hasn't been placed yet")
	}
	if result_robot != nil {
		t.Errorf("The robot should be nil until it's placed")
	}
}

func TestRunCommandUnrecognized(t *testing.T) {
	command_string := "blah blah what is this"
	var result_robot *robot.Robot
	var err error
	result_robot, err = runCommand(command_string, result_robot)

	if err == nil {
		t.Errorf("There should be an error here")
	}
}
