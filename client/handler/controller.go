package handler

import (
	"github.com/gdamore/tcell/v2"
	"os"
)

func (m *Handler) Listen() {

	for {
		ev := m.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'w':
					m.TankMoveUp()
				case 's':
					m.TankMoveDown()
				case 'a':
					m.TankMoveLeft()
				case 'd':
					m.TankMoveRight()
				case ' ':
					m.Fire()
				}
			case tcell.KeyCtrlC:
				os.Exit(1)
			}
		}
	}
}
