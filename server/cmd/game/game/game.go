package game

import (
	"log"
	"time"
)

func NewGame() *Game {
	log.Println("create quic")

	g := &Game{
		RockBucket:      make([]*Rock, 0),
		BulletBucket:    make(map[int32]*Bullet),
		TankBucket:      make(map[int32]*Tank),
		ExplosionBucket: make([]*Explosion, 0),
	}
	g.GenerateRocks()
	go func() {
		for {
			if len(g.ExplosionBucket) > 0 {
				g.ExplosionBucket = g.ExplosionBucket[1:]
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return g
}

type Game struct {
	ExplosionBucket []*Explosion
	RockBucket      []*Rock
	BulletBucket    map[int32]*Bullet
	TankBucket      map[int32]*Tank
}
