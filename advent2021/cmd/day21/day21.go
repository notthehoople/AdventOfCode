package main

import (
	"aoc/advent2021/utils"
	"fmt"
)

func rollTheDice(diceRoll int, rollTimes int) int {
	var totalDiceRoll, roll int

	for i := 0; i < rollTimes; i++ {
		roll = diceRoll + i
		if roll > 100 {
			roll = roll % 100
		}
	}
	totalDiceRoll = diceRoll
	roll = diceRoll + 1
	if roll > 100 {
		roll = roll % 100
	}

	totalDiceRoll += roll
	roll = diceRoll + 2
	if roll > 100 {
		roll = roll % 100
	}
	totalDiceRoll += roll

	return totalDiceRoll
}

func playGame(player1StartPos int, player2StartPos int, maxScore int, debug bool) int {
	var diceRoll int = 1
	var totalTimesDiceRolled int
	var player1Pos, player2Pos int = player1StartPos, player2StartPos
	var player1Score, player2Score int
	var movement int

	for player1Score < maxScore && player2Score < maxScore {
		movement = rollTheDice(diceRoll)
		//		fmt.Printf("player1 rolls %d, %d and %d\n", diceRoll, (diceRoll+1)%101, (diceRoll+2)%101)
		diceRoll = (diceRoll + 3) % 100
		if diceRoll == 0 {
			diceRoll = 100
		}
		totalTimesDiceRolled += 3

		fmt.Printf("player1 pos:%d movement:%d\n", player1Pos, movement)
		player1Pos = (player1Pos + movement) % 10
		if player1Pos == 0 {
			player1Pos = 10
		}
		player1Score += player1Pos

		movement = rollTheDice(diceRoll)
		//		fmt.Printf("player2 rolls %d, %d and %d\n", diceRoll, (diceRoll+1)%101, (diceRoll+2)%101)
		diceRoll = (diceRoll + 3) % 100
		if diceRoll == 0 {
			diceRoll = 100
		}
		totalTimesDiceRolled += 3

		fmt.Printf("player2 pos:%d movement:%d\n", player2Pos, movement)

		player2Pos = (player2Pos + movement) % 10
		if player2Pos == 0 {
			player2Pos = 10
		}
		player2Score += player2Pos

		fmt.Printf("diceRoll: %d player1Pos: %d player2Pos: %d\n", diceRoll, player1Pos, player2Pos)
	}

	fmt.Printf("player1score: %d player2score: %d totalTimes: %d\n",
		player1Score, player2Score, totalTimesDiceRolled)

	if player1Score >= maxScore {
		return totalTimesDiceRolled * player2Score
	}

	return totalTimesDiceRolled * player1Score
}

func solveDay(filename string, part byte, debug bool) int {
	puzzleInput, _ := utils.ReadFile(filename)

	var player1StartPos, player2StartPos int
	fmt.Sscanf(puzzleInput[0], "Player 1 starting position: %d\n", &player1StartPos)
	fmt.Sscanf(puzzleInput[1], "Player 2 starting position: %d\n", &player2StartPos)

	fmt.Printf("Player1: %d Player2: %d\n", player1StartPos, player2StartPos)

	finishScore := playGame(player1StartPos, player2StartPos, 1000, debug)

	return finishScore
}

// Main routine
func main() {
	filenamePtr, execPart, debug := utils.CatchUserInput()

	if execPart == 'z' {
		fmt.Println("Bad part choice. Available choices are 'a' and 'b'")
	} else {
		fmt.Printf("Result is: %d\n", solveDay(filenamePtr, execPart, debug))
	}
}
