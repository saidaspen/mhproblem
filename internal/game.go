package internal

import "math/rand"

const (
	Car = iota
	Goat
)

type Game struct{
	Doors []int
	WinningDoor int
}

func (g *Game) OpenDoors(pickedDoor int) int {
	if g.WinningDoor != pickedDoor {
		return g.WinningDoor
	}

	nDoorsTotal := len(g.Doors)
	nDoorsToOpen := nDoorsTotal - 2
	candidatesToKeep := make([]int, 0)

	i := 0
	for i < len(g.Doors) {
		if i != g.WinningDoor && i != pickedDoor {
			candidatesToKeep = append(candidatesToKeep, i)
		}
		i = i+1
	}

	// This is the door left after opening all other doors
	return candidatesToKeep[RandInt(0, nDoorsToOpen -1)]
}

func NewGame(numDoors int) Game{
	doors := make([]int, numDoors, numDoors)
	for i := range doors {
		doors[i] = Goat
	}
	winningDoor := RandInt(0, numDoors)
	doors[winningDoor] = Car
	g:=  Game{Doors : doors, WinningDoor : winningDoor}
	return g
}

func RandInt(min int, max int) int {
	if min == max{
		return min
	}
	return min + rand.Intn(max-min)
}

