package game

import "sync"

var (
	ExplosionBucket []*Explosion
	RockBucket      []*Rock
	BulletBucket    map[int32]*Bullet
	TankBucket      map[int64]*Tank
	Me              int64 // 自己的 id
)

func NewGame(id int64) {
	//ExplosionBucket = make([]*Explosion, 0)
	RockBucket = make([]*Rock, 0)
	// TODO sync.map
	BulletBucket = make(map[int32]*Bullet)
	TankBucket = make(map[int64]*Tank)

	Me = id
}

var Mu = sync.RWMutex{}
