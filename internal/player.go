package internal

var (
	AlwaysSwitch = func() bool { return true}
	AlwaysStay = func() bool { return false}
	FiftyFifty = func() bool {return RandInt(0, 2) > 0}
)

type Strategy func () bool

type Player struct {
	Strategy Strategy
	Wins     int
	Games     int
}

func (p *Player) PickDoor(numDoors int) int {
	return RandInt(0, numDoors)
}

func NewPlayer(strategy Strategy) Player {
	return Player{Strategy: strategy}
}
