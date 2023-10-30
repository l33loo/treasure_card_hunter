package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
		points := 0
		tries := 0
		for points < 3 {
			printMap(gameMap)
			fmt.Println("Pick a place on the map.\nEnter the coordinates as x, y, like 2, 3")

			mapCoord, _ := reader.ReadString('\n')
			sp := strings.Split(mapCoord, ",")
			coordX, _ := strconv.Atoi(strings.TrimSpace(sp[0]))
			// if error || if out of range
			coordY, _ := strconv.Atoi(strings.TrimSpace(sp[1]))
			// if error || if out of range
			foundTreasure := updateMap(&gameMap, coordX-1, coordY-1)
			printMap(gameMap)
			if foundTreasure {
				points++
			}
			tries++
			break
		}
		str := fmt.Sprintf("Treasures all found with %d tries!", tries)
		fmt.Println(str)
		// multiplayer mode
	} else {
		cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
		player1Points := 0
		player2Points := 0

		// full game
		for player1Points+player2Points < 3 {
			// card round
			for {
				player1Card := cards[rand.Intn(len(cards))]
				player2Card := cards[rand.Intn(len(cards))]

				// 0 is a tie
				roundWinner := 0

				if player1Card > player2Card {
					roundWinner = 1
				}

				if player1Card < player2Card {
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
				str2 := fmt.Sprintf("Player %d, pick a place on the map.\nEnter the coordinates as x, y, like 2, 3", roundWinner)
				fmt.Println(str2)

				mapCoord, _ := reader.ReadString('\n')
				sp := strings.Split(mapCoord, ",")
				coordX, _ := strconv.Atoi(strings.TrimSpace(sp[0]))
				// if error || if out of range
				coordY, _ := strconv.Atoi(strings.TrimSpace(sp[1]))
				// if error || if out of range
				foundTreasure := updateMap(&gameMap, coordX-1, coordY-1)
				printMap(gameMap)
				if foundTreasure {
					if roundWinner == 1 {
						player1Points++
					}

					if roundWinner == 2 {
						player2Points++
					}
				}

				break
			}
			if player1Points > player2Points {
				str := fmt.Sprintf("Player 1 wins with {%d} points!", player1Points)
				fmt.Println(str)
			} else {
				str := fmt.Sprintf("Player 2 wins with {%d} points!", player2Points)
				fmt.Println(str)
			}
		}
	}
}

type Row [6]string

func createMap(countTreasures int) [6]Row {
	var gameMap [6]Row
	// fmt.Printf("%v", gameMap)

	for i := 0; i < countTreasures; i++ {
		x := rand.Intn(6)
		y := rand.Intn(6)

		for gameMap[x][y] == "T" {
			x = rand.Intn(6)
			y = rand.Intn(6)
		}
		gameMap[x][y] = "T"
	}

	return gameMap
}

func updateMap(gameMap *[6]Row, x int, y int) bool {
	cell := (*gameMap)[x][y]
	if cell == "T" {
		(*gameMap)[x][y] = "X"
		return true
	} else if cell == "" {
		(*gameMap)[x][y] = "O"
		return false
	} else {
		// already guessed
		fmt.Println("Already picked")
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
