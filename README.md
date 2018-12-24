# Advent of Code 2018 & Learning to Code again

My attempts at the [Advent of Code](https://adventofcode.com/2018) puzzles.
I'm using this as a "learning to code again" exercise since I haven't coded for 20+ years.
I've chosen [go](https://golang.org)

## Instructions

### Compiling the code

```
cd ./srv/<dayxx>
go build
```

###

## Days of Advent

+ [Day 1](src/day01/day01.go)
Day 1 - [How to Write Go Code](https://golang.org/doc/code.html). A few nasty hacks in part b that don't feel good, especially the "break" in the "for" loop.
```
./day01 -file original.txt -part a
./day01 -file original.txt -part b
```

+ [Day 2](src/day02/day02.go)
Day 2 - 
```
./day02 -file original.txt -part a
./day02 -file original.txt -part b
```

+ [Day 3](src/day03/day03.go)
Day 3 - Done. A silly bug in my thoughts caused an hour of head scratching but sorted now.
```
./day03 -file original.txt -part a
./day03 -file original.txt -part b
```

+ [Day 4](src/day04/day04.go)
Day 4 - Done. Finally. Struggled getting part "b" map to work
```
./day04 -file original.txt -part a
./day04 -file original.txt -part b
```

+ [Day 5](src/day05/day05.go)
Day 5 - Done. Part "b" runs really slow. Could do with some more work!
```
./day05 -file original.txt -part a
./day05 -file original.txt -part b
```

+ [Day 6](src/day06)
Day 6 - Skipped.
```
```

+ [Day 7](src/day07/day07.go)
Day 7 - Part A done.
Day 7 - Part B done. Output isn't cleaned up. Answer for original.txt is one less than that listed
```
./day07 -file original.txt -part a
./day07 -file original.txt -part a -const 60 -workers 5
```

+ [Day 8](src/day08/day08.go)
Day 8 - Part A done.
Day 8 - Part B done.
```
./day08 -file original.txt -part a
./day08 -file original.txt -part b
```

+ [Day 10](src/day10/day10.go)
Day 10 - Part A done. Need to automate
Day 10 - Part B done. Need to automate
```
./day10 -file original.txt -grid 100 -max 100000 -part a | grep -v 50
    # Look for the highest number - this is the most likely time of the message. Then:
./day10 -file original.txt -grid 100 -max 100000 -part a -printsecond <time> | grep -v 50
    # Prints out the message
./day10 -file original.txt -grid 100 -max 100000 -part a -printsecond <time> | grep -v 50
    # Remember we count time from 0, so add 1 to the "printsecond" time to get your result
```

+ [Day 11](src/day11/day11.go)
Day 11 - Part A done
Day 11 - Part B done. Slow, but it works
```
./day11 -part a -puzzle 7165
./day11 -part b -puzzle 7165
```

+ [Day 12](src/day12/day12.go)
Day 12 - Part A done
Day 12 - Part B done
```
./day12 -file original.txt -part a -generations 20
./day12 -file original.txt -part b
```

+ [Day 18](src/day18/day18.go)
Day 18 - Part A done.
Day 18 - Got answer. Need to automate.
```
./day18 -file original.txt -part a -minutes 10 -grid 50
./day18 -file original.txt -part b -minutes 1000 -grid 50
```