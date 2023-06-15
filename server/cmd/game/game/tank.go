package game

type Tank struct {
	X         int32
	Y         int32
	Id        int32
	Direction rune
	IsLoading bool

	Name  string
	Level int32
	Kill  int32
}

func (g *Game) NewTank(id int32) {
	t := &Tank{
		X:         1,
		Y:         1,
		Direction: '↑',
		Id:        id,
		Kill:      0,
	}
	g.TankBucket[id] = t
}

func (g *Game) TankMove(id, direction int32) {
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
