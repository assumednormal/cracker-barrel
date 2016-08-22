package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/assumednormal/cracker-barrel"
)

func main() {
	nGames := flag.Int("n.games", 1, "number of games to simulate")

	flag.Parse()

	if *nGames <= 0 {
		panic("n.games must be greater than 0")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *nGames; i++ {
		fmt.Println(play(r))
	}
}

func play(r *rand.Rand) int {
	game := cb.NewGame()

	for {
		moves := game.AvailableMoves()

		nMoves := len(moves)
		if nMoves == 0 {
			break
		}

		move := moves[r.Intn(nMoves)]

		game.PlayMove(move)
	}

	return len(game.Filled)
}
