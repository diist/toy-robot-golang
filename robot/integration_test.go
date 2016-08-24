package robot

import "testing"

func TestIntegrationA(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("There should be no error")
	}

	robot = Move(robot)

	if Report(robot) != "0,1,NORTH" {
		t.Errorf("Failed integration test A")
	}
}

func TestIntegrationB(t *testing.T) {
	robot, err := Place(0, 0, "NORTH")
	if err != nil {
		t.Errorf("Failed integration test B")
	}

	robot = Left(robot)

	if Report(robot) != "0,0,WEST" {
		t.Errorf("Report doesn't work")
	}
}

func TestIntegrationC(t *testing.T) {
	robot, err := Place(1, 2, "EAST")
	if err != nil {
		t.Errorf("There should be no error")
	}

	robot = Move(robot)
	robot = Move(robot)
	robot = Left(robot)
	robot = Move(robot)

	if Report(robot) != "3,3,NORTH" {
		t.Errorf("Failed integration test C")
	}
}
