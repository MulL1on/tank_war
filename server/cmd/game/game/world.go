package game

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

// GenerateRocks randomly generates rocks
func (g *Game) GenerateRocks() {
	blockCount := 7

	for i := 0; i < blockCount; i++ {
		width := NRand(2, 6)
		height := NRand(2, 6)
		r := &Rock{
			Width:  width,
			Height: height,
			X:      NRand(0+width, GlobalConfig.ScreenWidth-width),
			Y:      NRand(0+height, GlobalConfig.ScreenHeight-height),
		}

		g.RockBucket = append(g.RockBucket, r)
	}
}

func (g *Game) NewExplosion(x, y int32) {
	e := &Explosion{
		X: x,
		Y: y,
	}
	g.ExplosionBucket = append(g.ExplosionBucket, e)
}
