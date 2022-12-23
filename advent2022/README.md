[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/notthehoople/advent2017?color=blueviolet)](https://golang.org/doc/go1.17) [![Go Report Card](https://goreportcard.com/badge/github.com/notthehoople/advent2017)](https://goreportcard.com/report/github.com/notthehoople/advent2017)

# advent2022
Advent of Code 2022

## Instructions

### Compiling the code

```
cd ./cmd/<dayxx>
go build
```

## Days of Advent

### Day 1 - Calorie Counting

Edge case fun with part b on the test input.

+ [Day 1](cmd/day01/day01.go)

```
./day01 -part a
./day01 -part b
```

### Day 2 - Rock Paper Scissors

+ [Day 2](cmd/day02/day02.go)

```
./day02 -part a [-file <filename>] [-debug]
./day02 -part b [-file <filename>] [-debug]
```

### Day 3 - Rucksack Reorganization

+ [Day 3](cmd/day03/day03.go)

```
./day03 -part a [-file <filename>] [-debug]
./day03 -part b [-file <filename>] [-debug]
```

### Day 4 - Camp Cleanup

+ [Day 4](cmd/day04/day04.go)

```
./day04 -part a [-file <filename>] [-debug]
./day04 -part b [-file <filename>] [-debug]
```

### Day 5 - Supply Stacks

+ [Day 5](cmd/day05/day05.go)

```
./day05 -part a [-file <filename>] [-debug]
./day05 -part b [-file <filename>] [-debug]
```

### Day 6 - Tuning Trouble

+ [Day 6](cmd/day06/day06.go)

```
./day06 -part a [-file <filename>] [-debug]
./day06 -part b [-file <filename>] [-debug]
```

### Day 7 - No Space Left On Device

+ [Day 7](cmd/day07/day07.go)

Needs tidied up. Part b doesn't give the answer directly. You need to pick from a list and work out what's the lowest value.

```
./day07 -part a [-file <filename>] [-debug]
./day07 -part b [-file <filename>] [-debug]
```

### Day 8 - Treetop Tree House

+ [Day 8](cmd/day08/day08.go)

```
./day08 -part a [-file <filename>] [-debug]
./day08 -part b [-file <filename>] [-debug]
```

### Day 9 - Rope Bridge

+ [Day 9](cmd/day09/day09.go)

```
./day09 -part a [-file <filename>] [-debug]
```

### Day 10 - Cathode-Ray Tube

+ [Day 10](cmd/day10/day10.go)

```
./day10 [-file <filename>] [-debug]
```

### Day 11 - Monkey in the Middle

+ [Day 11](cmd/day11/day11.go)

```
./day11 -part a [-file <filename>] [-debug]
./day11 -part b [-file <filename>] [-debug]
```

### Day 12 - 
### Day 13 - 
### Day 14 - Regolith Reservoir

+ [Day 14](cmd/day14/day14.go)

Note for part B: it will print out the complete finished cave system which is fun to look at!

```
./day14 -part a [-file <filename>] [-debug]
./day14 -part b [-file <filename>] [-debug]
```

### Day 15 - 
### Day 16 - 
### Day 17 - 
### Day 18 - 
### Day 19 - 
### Day 20 - 
### Day 21 - Monkey Math

+ [Day 21](cmd/day21/day21.go)

Need to revisit Part B. I've got the answer so I'm carrying on with other puzzles, but this code is utter crap and needs manual help to get the answer. To make it work, change humnShouts in calcMonkeySpeachPartB(). Set it to 1 to start with, and set changeSteps to a big number (say 1000000000). When the code stops, take the "humn shouts" number output 3 STEPS ABOVE where it stops. Use this number to replace humnShouts in the code, reduce the size of changeSteps, then run again. Repeat until you get the result "humn needs to shout: <x>" and <x> is the result you want.

```
./day21 -part a [-file <filename>] [-debug]
./day21 -part b [-file <filename>] [-debug]
```

### Day 22 - 
### Day 23 - 
### Day 24 - 
### Day 25 - 
