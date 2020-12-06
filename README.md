[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/notthehoople/advent2020?color=blueviolet)](https://golang.org/doc/go1.15) [![Go Report Card](https://goreportcard.com/badge/github.com/notthehoople/advent2020)](https://goreportcard.com/report/github.com/notthehoople/advent2020)

# advent2020
Advent of Code 2020

## Instructions

### Compiling the code

```
cd ./cmd/<dayxx>
go build
```

## Days of Advent

### Day 1 - Report Repair

+ [Day 1a](cmd/day01/day01.go)
+ [Day 1b](cmd/day01/day01.go)

```
./day01 -part a -file puzzleInput.txt
./day01 -part b -file puzzleInput.txt
```

### Day 2 - Password Philosophy

+ [Day 2a](cmd/day02/day02.go)
+ [Day 2b](cmd/day02/day02.go)

```
./day02 -part a -file puzzleInput.txt
./day02 -part b -file puzzleInput.txt
```

### Day 3 - Toboggan Trajectory

+ [Day 3a](cmd/day03/day03.go)
+ [Day 3b](cmd/day03/day03.go)

```
./day03 -part a -file puzzleInput.txt [-debug] [-slopex X] [-slopey Y]
./day03 -part b -file puzzleInput.txt [-debug]
```

### Day 4 - Passport Processing

+ [Day 4](cmd/day04/day04.go)

This code is nasty. Hardcoded array of structs, long winded validations. Needs to be reworked!

```
./day04 -part a -file puzzleInput.txt [-debug]
./day04 -part b -file puzzleInput.txt [-debug]
```

### Day 5 - Binary Boarding

+ [Day 5](cmd/day05/day05.go)

```
./day05 -part a -file puzzleInput.txt [-debug]
./day05 -part b -file puzzleInput.txt [-debug]
```

### Day 6 - Custom Customs

+ [Day 6](cmd/day06/day06.go)

```
./day06 -part a -file puzzleInput.txt [-debug]
```
