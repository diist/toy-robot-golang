package main

import (
	"fmt"
	"github.com/diist/toy-robot-golang/ui"
	"os"
)

func main() {
	fmt.Println("~~ Toy Robot ~~")
	filename := os.Getenv("ROBOT_FILE")
	ui.Play(filename)
}
