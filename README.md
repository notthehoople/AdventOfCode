[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/notthehoople/advent2020?color=blueviolet)](https://golang.org/doc/go1.15) [![Go Report Card](https://goreportcard.com/badge/github.com/notthehoople/advent2020)](https://goreportcard.com/report/github.com/notthehoople/advent2020)

# advent2020
Advent of Code 2020

## Progress
| Year  | Stars by 25th | Overall |
| :---: | :-----------: | :-----: |
| 2020 | 41 | 43 |
| 2019 | 17 | 21 |
| 2018 | 24 | 39 |
| 2017 | 0  | 2  |
| 2016 | 0  | 0  |
| 2015 | 0  | 0  |

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
./day06 -part b -file puzzleInput.txt [-debug]
```

### Day 7 - Handy Haversacks

+ [Day 7a](cmd/day07/day07.go)
+ [Day 7b](cmd/day07/day07b.go)

```
./day07 -part a -file puzzleInput.txt [-debug]
./day07 -part b -file puzzleInput.txt [-debug]
```

### Day 8 - Handheld Halting

+ [Day 8a](cmd/day08/day08.go)
+ [Day 8b](cmd/day08/day08b.go)

```
./day08 -part a -file puzzleInput.txt [-debug]
./day08 -part b -file puzzleInput.txt [-debug]
```

### Day 9 - Encoding Error

+ [Day 9a](cmd/day09/day09.go)
+ [Day 9b](cmd/day09/day09.go)

```
./day09 -part a -pre 5 -file testInput.txt [-debug]
./day09 -part a -pre 25 -file puzzleInput.txt [-debug]
./day09 -part b -pre 5 -file testInput.txt [-debug]
./day09 -part b -pre 25 -file puzzleInput.txt [-debug]
```

### Day 10 - Adapter Array

+ [Day 10a](cmd/day10/day10.go)
+ [Day 10b](cmd/day10/day10b.go)

```
./day10 -part a -file puzzleInput.txt [-debug]
./day10 -part b -file puzzleInput.txt [-debug]
```

### Day 11 - Seating System

+ [Day 11a](cmd/day11/day11.go)
+ [Day 11b](cmd/day11/day11b.go)

Unneccessarily verbose and plenty of duplicated code between part a and b that could be simplified. At least it works!

```
./day11 -part a -file puzzleInput.txt [-debug]
./day11 -part b -file puzzleInput.txt [-debug]
```

### Day 12 - Rain Risk

+ [Day 12a](cmd/day12/day12.go)
+ [Day 12b](cmd/day12/day12b.go)

Took ages to get my head around the rotations in part b!

```
./day12 -part a -file puzzleInput.txt [-debug]
./day12 -part b -file puzzleInput.txt [-debug]
```

### Day 13 - Shuttle Search

+ [Day 13a](cmd/day13/day13.go)
+ [Day 13b](cmd/day13/day13b.go)

That took ages to make part b work quickly enough. Got there in the end.
Note that part b has my puzzle input hard coded into main(). Change it here for your own puzzle input if you need

```
./day13 -part a -file puzzleInput.txt [-debug]
./day13 -part b
```

### Day 14 - Docking Data

+ [Day 14a](cmd/day14/day14.go)
+ [Day 14b](cmd/day14/day14b.go)

That was horrible. Got all twisted up with building lists in the wrong way. Ended up with a nasty hardcoded array of strings in the middle of a function as I was fed up with it!

```
./day14 -part a -file puzzleInput.txt [-debug]
./day14 -part b -file puzzleInput.txt [-debug]
```

### Day 15 - Rambunctious Recitation

+ [Day 15a](cmd/day15/day15.go)
+ [Day 15b](cmd/day15/day15.go)

Surprised that part a was efficient enough that part b was just "add the input and run it"!

```
./day15 -part a [-debug]
./day15 -part b [-debug]
```

### Day 16 - Ticket Translation

+ [Day 16a](cmd/day16/day16.go)
+ [Day 16b](cmd/day16/day16b.go)

Part b doesn't work. It *does* print out enough to be able to work out the answer by hand (yes, really), but needs to be finished to complete automatically. For now, I have my star so I don't care

```
./day16 -part a -file puzzleInput.txt [-debug]
./day16 -part b -file puzzleInput.txt [-debug]
```

### Day 17 - Conway Cubes

+ [Day 17a](cmd/day17/day17.go)
+ [Day 17b](cmd/day17/day17b.go)

```
./day17 -part a -file puzzleInput.txt [-debug]
./day17 -part b -file puzzleInput.txt [-debug]
```

### Day 18 - Operation Order

+ [Day 18a](cmd/day18/day18.go)
+ [Day 18b]

```
./day18 -part a -file puzzleInput.txt [-debug]
```

I haven't solved Day 18 part b with code yet; I'll need to come back and do this properly. I realised that it was trivial to solve using vi search and replace then using bc so that's what I did:

```
$ vi puzzleInput.txt
:%s;^;(;g
:%s;$;);g
:%s;*;)*(;g
:wq
bc -l puzzleInput.txt > output2.txt
quit
paste -sd+ output2.txt | bc
```

### Day 20 - Jurassic Jigsaw

+ [Day 20a](cmd/day20/day20.go)

Feels like I've gone wrong with the way I've approached Day 20 part a. It gives the correct answer but I'm matching double the number of sides that I should be doing

```
./day20 -part a -file puzzleInput.txt [-debug]
```

### Day 21 - Allergen Assessment

+ [Day 21a](cmd/day21/day21.go)
+ [Day 21b](cmd/day21/day21.go)

Very wordy code but it works quickly so it'll do

```
./day21 -part a -file puzzleInput.txt [-debug]
./day21 -part b -file puzzleInput.txt [-debug]
```

### Day 22 - Crab Combat

+ [Day 22a](cmd/day22/day22.go)

```
./day22 -part a -file puzzleInput.txt [-debug]
```

### Day 23 - Crab Cups

+ [Day 23a](cmd/day23/day23.go)
+ [Day 23b](cmd/day23/day23.go)

First time using container/ring. Got it working but had to rework the ring for part b. Went from looping through the ring looking for the destination, to using a map of rings as an index so I could jump around the ring easily

```
./day23 -part a [-debug]
./day23 -part b [-debug]
```

### Day 24 - Lobby Layout

+ [Day 24a](cmd/day24/day24.go)
+ [Day 24b](cmd/day24/day24b.go)

Getting a good coordinate system is a huge help here. I've used Axial Coordinates as described here https://www.redblobgames.com/grids/hexagons/

```
./day24 -part a -file puzzleInput.txt [-debug]
./day24 -part b -file puzzleInput.txt [-debug]
```

### Day 25 - Combo Breaker

+ [Day 25](cmd/day25/day25.go)

```
./day25 [-debug]
```
