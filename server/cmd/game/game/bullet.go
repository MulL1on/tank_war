package game

var d = [4][2]int32{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

type Bullet struct {
	Id        int32
	X         int32
	Y         int32
	Direction int32
	TankId    int64
}

func (g *Game) NewBullet(b *Bullet) {
	g.BulletBucket[b.Id] = b
}

func (b *Bullet) Move() {
	b.X += d[b.Direction][0]
	b.Y += d[b.Direction][1]
}

func (g *Game) RemoveBullet(b *Bullet) {
	delete(g.BulletBucket, b.Id)
}
