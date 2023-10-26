package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Failed to find the "go" binary in either GOROOT() or
// PATH(/home/lila/inst/bin:/home/lila/cmake-3.17.5-Linux-x86_64/bin:/home/lila/gcc-arm-none-eabi-9-2019-q4-major/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/home/lila/.fly/bin:/home/lila/Downloads/discord-0.0.30).
// Check PATH, or Install Go and reload the window.
// If PATH isn't what you expected, see https://github.com/golang/vscode-go/issues/971

func main() {
	fmt.Print("Bemvindo ao Treasure Card Hunter!")
	startOk := false

	reader := bufio.NewReader(os.Stdin)
	gameModeChoice := ""

	for !startOk {
		fmt.Println("Escolha o modo de jogo:\n(1) para single player\n(2) para multiplayer\n(x) para sair")

		gameModeChoice, _ = reader.ReadString('\n')

		switch strings.TrimSpace(gameModeChoice) {
		case "1":
			fmt.Println("Em modo single player")
			startOk = true
		case "2":
			fmt.Println("Em modo multiplayer")
			startOk = true
		case "x":
			fallthrough
		case "X":
			os.Exit(0)
		default:
			fmt.Println("Commande invalida.")
		}
	}

	cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffleCards(cards)

	isGameOver := false
	player1Points := 0
	player2Points := 0

	for !isGameOver {
		
		if gameModeChoice == "1" {

		} else /* if gameModeChoice == "2" */ {
			player1Card := giveCard(cards)
			player2Card := giveCard(cards)

			roundWinner := player1Card > player2Card ? 1 : 2

			fmt.Println("")

			fmt.Println()
		}
	}
}

func shuffleCards(array []int) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func giveCard(cards []int) int {
	card := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	return card
}
