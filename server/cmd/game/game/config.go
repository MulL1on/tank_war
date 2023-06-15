package game

type Config struct {
	ScreenWidth  int32
	ScreenHeight int32

	BorderChar    rune
	BulletChar    rune
	RockChar      rune
	ExplosionChar rune
}

var GlobalConfig = Config{
	ScreenWidth:  80,
	ScreenHeight: 20,

	BorderChar:    '#',
	BulletChar:    '*',
	RockChar:      '@',
	ExplosionChar: 'X',
}
