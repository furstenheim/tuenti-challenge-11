package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/furstenheim/challenge_encoding"
	"strconv"
	"strings"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	challenge_encoding.Unmarshall(challenge, reader)
	log.Println(challenge.Games[1])
	for i, g := range(challenge.Games) {
		result := solveGame(g)
		logResult(i, result)
	}
}

func logResult(i int, result string) {
	fmt.Printf("Cases #%d: %s\n", i + 1, result)
}

func solveGame (game Game) string {
	rawDice := strings.Split(game.Raw, ":")
	left, leftErr := strconv.Atoi(rawDice[0])
	right, rightErr := strconv.Atoi(rawDice[1])

	if leftErr != nil {
		log.Fatalf("Unknown value %s", rawDice[0])
	}

	if rightErr != nil {
		log.Fatalf("Unknown value %s", rawDice[1])
	}

	if left + right == 12 {
		return "-"
	}

	return strconv.Itoa(left + right + 1)

}


type Challenge struct {
	NCases int   `index:"0"`
	Games []Game `index:"1" indexed:"NCases"`
}
type Game struct {
	Raw string `index:"0"`
}