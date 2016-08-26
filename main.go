package main

import (
	"fmt"
	"github.com/diist/toy-robot-golang/ui"
	"os"
)

func main() {
	fmt.Println("~~ Toy Robot ~~")
	filename := os.Getenv("ROBOT_FILE")
	if filename != "" {
		fmt.Printf("Playing with the file: %q \n", filename)
		ui.PlayWithFile(filename)
	} else {
		fmt.Println("Playing with the CLI")
		ui.PlayWithCLI()
	}
}
