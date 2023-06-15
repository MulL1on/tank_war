package game

import "github.com/gdamore/tcell/v2"

type Tank struct {
	Id        int32
	X         int32
	Y         int32
	Direction rune
	IsLoading bool

	Kill int32
	Name int32
}

func DrawTank(screen tcell.Screen) {
	for _, v := range TankBucket {
		screen.SetContent(int(v.X), int(v.Y), v.Direction, nil, tcell.StyleDefault)
	}
}
