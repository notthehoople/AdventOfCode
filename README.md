[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/notthehoople/advent2019?color=blueviolet)](https://golang.org/doc/go1.13) [![Go Report Card](https://goreportcard.com/badge/github.com/notthehoople/advent2019)](https://goreportcard.com/report/github.com/notthehoople/advent2019)

# advent2019
Advent of Code 2019

## Instructions

### Compiling the code

```
cd ./srv/<dayxx>
go build
```

## Days of Advent

### Day 1 - The Tyranny of the Rocket Equation

+ [Day 1a](src/day01/day01a.go)
+ [Day 1b](src/day01/day01b.go)

Nothing clever here. Hard coded the input data into a variable declaration. Would be nicer to have this read from the file
```
./day01a
./day01b
```

### Day 2 - 1202 Program Alarm

+ [Day 2](src/day02/day02.go)

```
./day02 -file input.txt -noun 12 -verb 2 -part a
./day02 -file input.txt -part b
```

### Day 3 - Crossed Wires
This is awful. It works for part a but is really slow and inefficient, and isn't going to work for part b. I'll come back to this later and rework it
+ [Day 3](src/day03/day03.go)

```
./day03 -file input.txt
```

### Day 4 - Secure Container

+ [Day 4](src/day04/day04.go)

```
./day04 -start 156218 -end 652527 -part a
./day04 -start 156218 -end 652527 -part b
```