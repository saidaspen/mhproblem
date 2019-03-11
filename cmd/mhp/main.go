package main

import (
	"flag"
	"fmt"
	"github.com/saidaspen/mhproblem/internal"
	"math"
	"math/rand"
	"time"
)

// Defaults
const (
	defaultGames   = 100
	defaultDoors   = 3
	defaultVerbose = false
)

var (
	numGames int
	numDoors int
	verbose  bool
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.IntVar(&numGames, "n", defaultGames, "number of games")
	flag.IntVar(&numDoors, "d", defaultDoors, "number of doors")
	flag.BoolVar(&verbose, "v", defaultVerbose, "verbose")
	flag.Parse()
	games := make([]internal.Game, numGames, numGames)
	for i := 0; i < numGames; i++ {
		games[i] = internal.NewGame(numDoors)
	}
	fmt.Printf("Created %v games with %v doors each.\n", numGames, numDoors)

	pAlwaysSwitch := internal.NewPlayer(internal.AlwaysSwitch)
	pAlwaysStay := internal.NewPlayer(internal.AlwaysStay)
	pFiftyFifty := internal.NewPlayer(internal.FiftyFifty)
	for i := 1; i <= numGames; i++ {
		game := games[i-1]
		play(&pAlwaysSwitch, &game, i)
		play(&pAlwaysStay, &game, i)
		play(&pFiftyFifty, &game, i)
	}
	fmt.Printf("Player that always switch won %v%%\n", wPerc(&pAlwaysSwitch))
	fmt.Printf("Player that always stays won %v%%\n", wPerc(&pAlwaysStay))
	fmt.Printf("Player fifty-fifty won %v%%\n", wPerc(&pFiftyFifty))
}

func wPerc(p *internal.Player) float64 {
	return roundToTwo(float64(p.Wins) / float64(p.Games) * 100)
}

func roundToTwo(x float64) float64 {
	return math.Round(x*100) / 100
}

func play(player *internal.Player, game *internal.Game, gameID int) {
	player.Games++
	orgPickedDoor := player.PickDoor(len(game.Doors))
	if verbose {
		fmt.Printf("Game: %v\tPlayer picks door %v\n", gameID, orgPickedDoor)
	}
	otherDoor := game.OpenDoors(orgPickedDoor)
	if verbose {
		fmt.Printf("Game: %v\tMonty opens all doors except for %v\n", gameID, otherDoor)
	}
	var finalChoiceOfDoor int
	if player.Strategy() {
		if verbose {
			fmt.Printf("Game: %v\tPlayer switched from door %v to door %v\n", gameID, orgPickedDoor, otherDoor)
		}
		finalChoiceOfDoor = otherDoor
	} else {
		if verbose {
			fmt.Printf("Game: %v\tPlayer stays with door %v\n", gameID, orgPickedDoor)
		}
		finalChoiceOfDoor = orgPickedDoor
	}
	if finalChoiceOfDoor == game.WinningDoor {
		if verbose {
			fmt.Printf("Game: %v\tPlayer wins! The correct door was %v\n", gameID, game.WinningDoor)
		}
		player.Wins++
	} else if verbose {
		fmt.Printf("Game: %v\tPlayer looses! The correct door was %v\n", gameID, game.WinningDoor)
	}

	if verbose {
		fmt.Println("...........................................................")
	}
}
