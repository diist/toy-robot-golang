package main

import (
	"fmt"
	"github.com/diist/toy-robot-golang/input"
	"os"
)

func main() {
	fmt.Println("~~ Toy Robot ~~")
	filename := os.Getenv("ROBOT_FILE")
	if filename != "" {
		fmt.Printf("Playing with the file: %q \n", filename)
		input.PlayWithFile(filename)
	} else {
		fmt.Println("Playing with the CLI")
		input.PlayWithCLI()
	}
}
