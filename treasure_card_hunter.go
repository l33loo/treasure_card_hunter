package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Bemvindo ao Treasure Card Hunter!")

	reader := bufio.NewReader(os.Stdin)
	gameModeChoice := ""

intro:
	for {
		fmt.Println("Escolha o modo de jogo:\n(1) para single player\n(2) para multiplayer\n(x) para sair")

		gameModeChoice, _ = reader.ReadString('\n')

		switch strings.TrimSpace(gameModeChoice) {
		case "1":
			fmt.Println("Em modo single player")
			break intro
		case "2":
			fmt.Println("Em modo multiplayer")
			break intro
		case "x":
			fallthrough
		case "X":
			os.Exit(0)
		default:
			fmt.Println("Comande invalido.")
		}
	}

	gameMap := createMap(3)
	fmt.Printf("%v\n", gameMap)

	print()

	// single player mode
	if gameModeChoice == "1" {
		for {
			//

			break
		}
		// multiplayer mode
	} else {
		cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

		// full fame
		for {
			// card round
			for {
				player1Card := cards[rand.Intn(len(cards))]
				player2Card := cards[rand.Intn(len(cards))]

				// 0 is a tie
				roundWinner := 0

				if player1Card > player2Card {
					roundWinner = 1
				} else if player1Card < player2Card {
					roundWinner = 2
				}

				if roundWinner == 0 {
					str := fmt.Sprintf("Player 1 played with card {%d}, Player 2 with card {%d}.\nIt's a tie - let's try again!", player1Card, player2Card)
					fmt.Println(str)
					continue
				}

				str := fmt.Sprintf("Player 1 played with card {%d}, Player 2 with card {%d}.\nPlayer %d wins this round!", player1Card, player2Card, roundWinner)
				fmt.Println(str)

				printMap(gameMap)

				break
			}
			//
			break
		}
	}
}

type Row [6]string

func createMap(countTreasures int) [6]Row {
	var gameMap [6]Row
	// fmt.Printf("%v", gameMap)

	for i := 0; i < countTreasures; i++ {
		x := rand.Intn(countTreasures)
		y := rand.Intn(countTreasures)

		gameMap[x][y] = "T"
	}

	return gameMap
}

func updateMap(gameMap [6]Row, x int, y int) bool {
	cell := gameMap[x][y]
	if cell == "T" {
		gameMap[x][y] = "X"
		return true
	} else if cell == "" {
		gameMap[x][y] = "O"
		return false
	} else {
		// already guessed
		return false
	}
}

func printMap(gameMap [6]Row) {
	fmt.Println("    1   2   3   4   5   6")
	for i := 0; i < len(gameMap); i++ {
		str := fmt.Sprintf("%d |", (i + 1))
		for j := 0; j < len(gameMap[i]); j++ {
			val := gameMap[i][j]
			if val == "T" {
				val = ""
			}
			if val == "" {
				val = " "
			}
			str += fmt.Sprintf(" %s |", val)
		}
		str += "\n---------------------------"
		fmt.Println(str)
	}
}
