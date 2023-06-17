package game

import (
	"time"
)

func NewGame() *Game {

	g := &Game{
		RockBucket:      make([]*Rock, 0),
		BulletBucket:    make(map[int32]*Bullet),
		TankBucket:      make(map[int64]*Tank),
		ExplosionBucket: make([]*Explosion, 0),
	}
	g.GenerateRocks()
	go func() {
		for {
			if len(g.ExplosionBucket) > 0 {
				g.ExplosionBucket = g.ExplosionBucket[1:]
			}
			time.Sleep(500 * time.Millisecond) // 设置爆炸效果的持续时间
		}
	}()
	return g
}

type Game struct {
	ExplosionBucket []*Explosion
	RockBucket      []*Rock
	BulletBucket    map[int32]*Bullet
	TankBucket      map[int64]*Tank
}
