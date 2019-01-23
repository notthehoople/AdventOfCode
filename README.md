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

Day 2 - Done
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

Day 6 - Part A done.

Day 6 - Part B not started yet
```
./day06 -file original.txt -part a [-debug]
```

+ [Day 7](src/day07/day07.go)

Day 7 - Part A done.; 

Day 7 - Part B done. Output isn't cleaned up. Answer for original.txt is one less than that listed
```
./day07 -file original.txt -part a
./day07 -file original.txt -part a -const 60 -workers 5
```

+ [Day 8](src/day08/day08.go)

Day 8 - Part A done.; 

Day 8 - Part B done.
```
./day08 -file original.txt -part a
./day08 -file original.txt -part b
```

+ [Day 9](src/day09/day09.go)

Day 9 - Part A done;

Day 9 - Part B done but very, very slow. Needs to be rewritten to use ring package (https://golang.org/pkg/container/ring/) instead of array
```
./day09 -part a -marble 71223 -players 455
./day09 -part b -marble 7122300 -players 455
```

+ [Day 10](src/day10/day10.go)

Day 10 - Part A done. Need to automate; 

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

Day 11 - Part A done; 

Day 11 - Part B done. Slow, but it works
```
./day11 -part a -puzzle 7165
./day11 -part b -puzzle 7165
```

+ [Day 12](src/day12/day12.go)

Day 12 - Part A done; 

Day 12 - Part B done
```
./day12 -file original.txt -part a -generations 20
./day12 -file original.txt -part b
```

+ [Day 13](src/day13)

Day 13 - Skipped.
```
```

+ [Day 14](src/day14/day14.go)

Day 14 - Part A done; 

Day 14 - Part B done. Kinda. Produces the right answer, but has many bugs in output for other numbers. Needs work on the comparison
```
./day14 -part a -recipes 825401 -answers 10 [-print]
./day14 -part b -result "825401"
```

+ [Day 15](src/day15)

Day 15 - Skipped.
```
```

+ [Day 16a](src/day16a)
+ [Day 16b](src/day16b)

Day 16 - Part A done.

Day 16 - Part B done but included manual work. Will need to come back to it. Also hacked around with the original.txt input file
```
./day16a -part a -file original.txt
./day16b -part b -file originalb.txt
```

+ [Day 17](src/day17/day17.go)

Day 17 - Part A done. Tricky! Not an elegant solution; 

Day 17 - Part B done. Nothing to this, just an extra print statement.
```
./day17 -file original.txt -part a [> output]
./day17 -file original.txt -part b [> output]
```

+ [Day 18](src/day18/day18.go)

Day 18 - Part A done; 

Day 18 - Got answer. Need to automate.
```
./day18 -file original.txt -part a -minutes 10 -grid 50
./day18 -file original.txt -part b -minutes 1000 -grid 50
```

+ [Day 19](src/day19/day19.go)
+ [Day 19b](src/day19/day19b.go)

Day 19 - Part A done; 

Day 19 - Part B done. Doesn't use the input files. Works exactly for my input file data only
```
./day19 -file original.txt -part a
./day19b
```

+ [Day 21](src/day21/day21.go)
+ [Day 21a](src/day21a/day21a.go)
+ [Day 21b](src/day21b/day21b.go)

**Note for all parts of day 21 the answers are hardcoded to the specific input I've been given, and do not refer to the input files at all.**

Day 21 - Part A done. Brute force method;

Day 21a - Part A done after realising what's going on. Not brute force! 

Day 21b - Part B done. Not a clean answer, and needs manual work on the output file. Come back to this!
```
./day21
./day21a
./day21b > output      # Needs manual work done on the output file to find the answer
```

+ [Day 23](src/day23/day23.go)

Day 23 - Part A done; 

Day 23 - Part B Not complete. Not sure what to do with this yet
```
./day23 -file original.txt -part a [-debug]
./day23 -file original.txt -part b [-debug]
```

+ [Day 25](src/day25/day25.go)

Day 25 - Part A done; 

Day 25 - Part B.....there is no part B
```
./day25 -file original.txt -range 3
or
./runtests
```