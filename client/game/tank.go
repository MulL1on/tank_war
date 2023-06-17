package game

import (
	"github.com/gdamore/tcell/v2"
	"sort"
)

type Tank struct {
	Id        int64
	X         int32
	Y         int32
	Direction rune
	IsLoading bool
	Color     uint64
	IsDead    bool

	Kill int32
	Name string
}

func DrawTank(screen tcell.Screen) {
	for _, v := range TankBucket {
		if v.IsDead {
			continue
		}
		screen.SetContent(int(v.X), int(v.Y), v.Direction, nil, tcell.StyleDefault.Foreground(tcell.Color(v.Color)))
	}
}

func DrawPlayerList(screen tcell.Screen) {
	x := int32(GlobalConfig.ScreenWidth) + 1

	y := int32(1)

	//转化为slice
	var tankSlice []*Tank
	for _, v := range TankBucket {
		tankSlice = append(tankSlice, v)
	}
	//排序
	sort.Slice(tankSlice, func(i, j int) bool {
		return tankSlice[i].Id > tankSlice[j].Id
	})

	for _, v := range tankSlice {
		if v.IsDead {
			screen.SetContent(int(x), int(y), 'X', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
		}
		writeStringToScreen(screen, x+2, y, v.Name, tcell.Color(v.Color))
		y++
	}
}

func writeStringToScreen(screen tcell.Screen, x, y int32, str string, color tcell.Color) {
	for _, v := range str {
		screen.SetContent(int(x), int(y), v, nil, tcell.StyleDefault.Foreground(color))
		x++
	}
}
