package input

import (
	"io/ioutil"
	"strings"
)

func linesFromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
