# toy-roboy-golang

A solution to the [toy robot problem](PROBLEM.md), in golang.

## Usage
### CLI
By default, commands are read from the command line.
```
$ toy-robot-golang
~~ Toy Robot ~~
Playing with the CLI
PLACE 1,2,SOUTH
MOVE
LEFT
MOVE
REPORT
2,1,EAST
```

### File
If you set a `ROBOT_FILE` environment variable, the file will be used instead of the CLI.
```
$ ROBOT_FILE=${PWD}/test-files/c.txt toy-robot-golang
~~ Toy Robot ~~
Playing with the file: "/path/to/test-files/c.txt"
3,3,NORTH
```

## Installation
### Using `go get`
If your `golang` environment is set up, you can:
```
$ go get github.com/diist/toy-robot-golang
```
And then:
```
$ toy-robot-golang
```

### Download binaries
You can download precompiled binaries from the [releases](https://github.com/diist/toy-robot-golang/releases) page.

## Running the tests
To run the tests:
```
$ go test ./...
?       github.com/diist/toy-robot-golang    [no test files]
ok      github.com/diist/toy-robot-golang/input    0.037s
ok      github.com/diist/toy-robot-golang/robot    0.034s
?       github.com/diist/toy-robot-golang/table    [no test files]
```
