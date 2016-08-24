package robot

import "testing"

func TestPlaceGood(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should not have been an error here")
	}
	if robot.X != 0 || robot.Y != 0 || robot.Direction != "NORTH" {
		t.Errorf("Failed placing the robot")
	}
}

func TestPlaceBadPosition(t *testing.T) {
	robot, err := Place(-1, 0, "NORTH")
	if err == nil {
		t.Errorf("There should be an error here")
	}
	if robot != nil {
		t.Errorf("Robot should be nil")
	}
}

func TestPlaceBadDirection(t *testing.T) {
	robot, err := Place(2, 2, "WARBLE")
	if err == nil {
		t.Errorf("There should be an error here")
	}
	if robot != nil {
		t.Errorf("Robot should be nil")
	}
}

func TestMove(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = Move(robot)
	if robot.Y != 1 {
		t.Errorf("Robot should have moved one unit north")
	}
}

func TestMoveEdge(t *testing.T) {
	robot, err := Place(4, 0, "EAST")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = Move(robot)
	if robot.X != 4 {
		t.Errorf("Robot should not move")
	}
}

func TestLeft(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = Left(robot)
	if robot.Direction != "WEST" {
		t.Errorf("Robot has not turned left")
	}
}

func TestRight(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	robot = Right(robot)
	if robot.Direction != "EAST" {
		t.Errorf("Robot has not turned right")
	}
}

func TestReport(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}
	if Report(robot) != "0,0,NORTH" {
		t.Errorf("Report doesn't work")
	}
}
