package game

import "github.com/gdamore/tcell/v2"

type Bullet struct {
	X         int32
	Y         int32
	Direction int32
	Id        int32
}

func DrawBullet(screen tcell.Screen) {
	bulletChar := GlobalConfig.BulletChar
	for _, b := range BulletBucket {
		screen.SetContent(int(b.X), int(b.Y), bulletChar, nil, tcell.StyleDefault)
	}
}
