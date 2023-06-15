package game

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

type Explosion struct {
	X int32
	Y int32
}

type Rock struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func DrawBorder(screen tcell.Screen) {
	// 绘制上边框
	borderChar := GlobalConfig.BorderChar
	screenWidth := int(GlobalConfig.ScreenWidth)
	screenHeight := int(GlobalConfig.ScreenHeight)

	for x := 0; x < screenWidth; x++ {
		screen.SetContent(x, 0, borderChar, nil, tcell.StyleDefault)

	}

	// 绘制下边框
	for x := 0; x < screenWidth; x++ {
		screen.SetContent(x, screenHeight-1, borderChar, nil, tcell.StyleDefault)
	}

	// 绘制左边框
	for y := 0; y < screenHeight; y++ {
		screen.SetContent(0, y, borderChar, nil, tcell.StyleDefault)
	}

	// 绘制右边框
	for y := 0; y < screenHeight; y++ {
		screen.SetContent(screenWidth-1, y, borderChar, nil, tcell.StyleDefault)
	}
}

func DrawRocks(screen tcell.Screen) {
	rockChar := GlobalConfig.RockChar
	for _, r := range RockBucket {
		// 遍历石头块的行和列
		for row := int32(0); row < r.Height; row++ {
			for col := int32(0); col < r.Width; col++ {
				// 设置石头块的字符和样式
				screen.SetContent(int(r.X+col), int(r.Y+row), rockChar, nil, tcell.StyleDefault)
			}
		}
	}
}

func DrawExplosion(screen tcell.Screen) {
	height := GlobalConfig.ExplosionHeight
	width := GlobalConfig.ExplosionWidth
	for _, e := range ExplosionBucket {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				screen.SetContent(int(e.X)+col, int(e.Y)+row, GlobalConfig.ExplosionChar, nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
			}
		}
	}
}

func DrawBoard(screen tcell.Screen) {
	t, ok := TankBucket[Me]
	if !ok {
		return
	}
	//log.Println("Me:", Me)
	board := []byte(fmt.Sprintf("kill:%d", t.Kill))
	//write board to the bottom
	for i := 0; i < len(board); i++ {
		screen.SetContent(i, int(GlobalConfig.ScreenHeight+1), rune(board[i]), nil, tcell.StyleDefault)
	}
}
