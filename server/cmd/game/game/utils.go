package game

import (
	"math"
	"math/rand"
)

type Position struct {
	X int32
	Y int32
}

func (p *Position) Equal(p2 *Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (g *Game) IsHitBorder(x int32, y int32) bool {
	return x <= 0 || x >= GlobalConfig.ScreenWidth-1 || y <= 0 || y >= GlobalConfig.ScreenHeight-1
}

// NRand returns a random number between min and max
func NRand(min, max int32) int32 {
	return rand.Int31n(max-min) + min
}

func (g *Game) IsHitRock(x int32, y int32) bool {
	for _, v := range g.RockBucket {
		if x >= v.X && x < v.X+v.Width && y >= v.Y && y < v.Y+v.Height {
			return true
		}
	}
	return false
}

func (g *Game) IsTankHitTank(x int32, y int32) bool {
	for _, v := range g.TankBucket {
		if v.X == x && v.Y == y {
			return true
		}
	}
	return false
}

func (g *Game) IsBulletHitTank(b *Bullet) bool {
	for _, v := range g.TankBucket {
		if v.X == b.X && v.Y == b.Y {
			if v.Id == b.TankId || v.IsDead {
				return false
			}
			g.TankBucket[b.TankId].Kill++
			g.TankBucket[v.Id].IsDead = true
			return true
		}
	}
	return false
}

func (g *Game) IsTankAroundTank(x int32, y int32) bool {
	for _, v := range g.TankBucket {
		if math.Abs(float64(v.X-x)) < 2 || math.Abs(float64(v.Y-y)) < 2 {
			return true
		}
	}
	return false
}
