package game

type Tank struct {
	X         int32
	Y         int32
	Id        int64
	Direction rune
	IsLoading bool
	Color     uint64
	IsDead    bool

	Name  string
	Level int32
	Kill  int32
}

func (g *Game) NewTank(name string, id int64, color uint64) {
	//random tank position
	for {
		x := NRand(0, GlobalConfig.ScreenWidth)
		y := NRand(0, GlobalConfig.ScreenHeight)
		if !g.IsHitBorder(x, y) && !g.IsHitRock(x, y) && !g.IsTankAroundTank(x, y) {
			g.TankBucket[id] = &Tank{
				X:         x,
				Y:         y,
				Direction: '↑',
				Id:        id,
				Kill:      0,
				Color:     color,
				Name:      name,
				IsDead:    false,
			}
			break
		}
	}
}

func (g *Game) TankMove(id int64, direction int32) {
	_, ok := g.TankBucket[id]
	if !ok {
		return
	}
	g.TankBucket[id].Direction = direction
	switch direction {
	case '↑':
		if g.IsHitBorder(g.TankBucket[id].X, g.TankBucket[id].Y-1) {
			return
		} else if g.IsHitRock(g.TankBucket[id].X, g.TankBucket[id].Y-1) {
			return
		} else if g.IsTankHitTank(g.TankBucket[id].X, g.TankBucket[id].Y-1) {
			return
		}
		g.TankBucket[id].Y--
		g.TankBucket[id].Direction = '↑'
	case '↓':
		if g.IsHitBorder(g.TankBucket[id].X, g.TankBucket[id].Y+1) {
			return
		} else if g.IsHitRock(g.TankBucket[id].X, g.TankBucket[id].Y+1) {
			return
		} else if g.IsTankHitTank(g.TankBucket[id].X, g.TankBucket[id].Y+1) {
			return
		}
		g.TankBucket[id].Y++
		g.TankBucket[id].Direction = '↓'
	case '←':
		if g.IsHitBorder(g.TankBucket[id].X-1, g.TankBucket[id].Y) {
			return
		} else if g.IsHitRock(g.TankBucket[id].X-1, g.TankBucket[id].Y) {
			return
		} else if g.IsTankHitTank(g.TankBucket[id].X-1, g.TankBucket[id].Y) {
			return
		}
		g.TankBucket[id].X--
		g.TankBucket[id].Direction = '←'
	case '→':
		if g.IsHitBorder(g.TankBucket[id].X+1, g.TankBucket[id].Y) {
			return
		} else if g.IsHitRock(g.TankBucket[id].X+1, g.TankBucket[id].Y) {
			return
		} else if g.IsTankHitTank(g.TankBucket[id].X+1, g.TankBucket[id].Y) {
			return
		}

		g.TankBucket[id].X++
		g.TankBucket[id].Direction = '→'
	}
}
