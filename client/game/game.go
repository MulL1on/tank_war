package game

var (
	ExplosionBucket []*Explosion
	RockBucket      []*Rock
	BulletBucket    map[int32]*Bullet
	TankBucket      map[int32]*Tank
	Me              int64 // 自己的 id
)

func NewGame(id int64) {
	//ExplosionBucket = make([]*Explosion, 0)
	RockBucket = make([]*Rock, 0)
	BulletBucket = make(map[int32]*Bullet)
	TankBucket = make(map[int32]*Tank)

	Me = id
}
