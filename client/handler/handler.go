package handler

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"tank_war/client/game"
	pb "tank_war/client/handler/pb/quic"
)

type Handler struct {
	client *Client
	screen tcell.Screen
}

func NewHandler(client *Client, screen tcell.Screen) *Handler {
	return &Handler{
		client: client,
		screen: screen,
	}
}

func (m *Handler) GetRockList(rockList *pb.Action_GetRockList) {
	for _, v := range rockList.GetRockList.Rock {
		rock := &game.Rock{
			X:      v.X,
			Y:      v.Y,
			Height: v.Height,
			Width:  v.Width,
		}
		game.RockBucket = append(game.RockBucket, rock)
	}
}

func (m *Handler) GetTankList(tankList *pb.Action_GetTankList) {
	game.TankBucket = make(map[int64]*game.Tank)
	for _, v := range tankList.GetTankList.Tank {

		tank := &game.Tank{
			X:         v.X,
			Y:         v.Y,
			Direction: v.Direction,
			Id:        v.Id,
			Color:     v.Color,
			Name:      v.Name,
			IsDead:    v.IsDead,

			Kill: v.Kill,
		}
		game.TankBucket[v.Id] = tank
		log.Println("get tank", v)
	}
}

func (m *Handler) GetBulletList(bulletList *pb.Action_GetBulletList) {
	game.BulletBucket = make(map[int32]*game.Bullet)
	for _, v := range bulletList.GetBulletList.Bullet {

		bullet := &game.Bullet{
			X: v.X,
			Y: v.Y,
		}
		game.BulletBucket[v.Id] = bullet
	}
}

func (m *Handler) GetExplosionList(explosionList *pb.Action_GetExplosionList) {
	game.ExplosionBucket = make([]*game.Explosion, 0)
	for _, v := range explosionList.GetExplosionList.Explosion {

		explosion := &game.Explosion{
			X: v.X,
			Y: v.Y,
		}
		game.ExplosionBucket = append(game.ExplosionBucket, explosion)
	}
}

func (m *Handler) TankMoveUp() {
	tankMove := &pb.Action_TankMove{
		TankMove: &pb.TankMove{
			Id:        m.client.id,
			Direction: '↑',
		},
	}
	action := &pb.Action{
		Type: tankMove,
	}

	m.client.send <- action

}

func (m *Handler) TankMoveDown() {
	tankMove := &pb.Action_TankMove{
		TankMove: &pb.TankMove{
			Id:        m.client.id,
			Direction: '↓',
		},
	}
	action := &pb.Action{
		Type: tankMove,
	}

	m.client.send <- action
}

func (m *Handler) TankMoveLeft() {
	tankMove := &pb.Action_TankMove{
		TankMove: &pb.TankMove{
			Id:        m.client.id,
			Direction: '←',
		},
	}
	action := &pb.Action{
		Type: tankMove,
	}

	m.client.send <- action
}

func (m *Handler) TankMoveRight() {
	tankMove := &pb.Action_TankMove{
		TankMove: &pb.TankMove{
			Id:        m.client.id,
			Direction: '→',
		},
	}
	action := &pb.Action{
		Type: tankMove,
	}

	m.client.send <- action
}

func (m *Handler) Fire() {
	t := game.TankBucket[m.client.id]
	var diretion int32
	switch t.Direction {
	case '↑':
		diretion = 0
	case '↓':
		diretion = 1
	case '←':
		diretion = 2
	case '→':
		diretion = 3
	}
	nb := &pb.Action_NewBullet{
		NewBullet: &pb.NewBullet{
			X:         t.X,
			Y:         t.Y,
			TankId:    t.Id,
			Direction: diretion,
		},
	}
	action := &pb.Action{
		Type: nb,
	}

	m.client.send <- action
}
