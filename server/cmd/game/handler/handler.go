package handler

import (
	"sync"
	"tank_war/server/cmd/game/game"
	pb "tank_war/server/cmd/game/handler/pb/quic"
	"tank_war/server/shared/consts"
	"time"
)

type Handler struct {
	//TODO: way to update status
	mu     sync.RWMutex
	status int
	game   *game.Game
}

func NewHandler() *Handler {
	return &Handler{
		mu:     sync.RWMutex{},
		game:   game.NewGame(),
		status: consts.GameNone,
	}
}

func (h *Handler) NewTank(name string, id int64, color uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.game.NewTank(name, id, color)
}

func (h *Handler) GetRockList() *pb.Action {
	// 不用加锁，这是只读的
	rockList := &pb.GetRockList{}

	for _, v := range h.game.RockBucket {
		rock := &pb.Rock{
			X:      v.X,
			Y:      v.Y,
			Height: v.Height,
			Width:  v.Width,
		}
		rockList.Rock = append(rockList.Rock, rock)
	}
	action := &pb.Action{
		Type: &pb.Action_GetRockList{
			GetRockList: rockList,
		},
	}
	return action
}

func (h *Handler) TankMove(move *pb.Action_TankMove) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.game.TankMove(move.TankMove.Id, move.TankMove.Direction)
}

func (h *Handler) NewBullet(nb *pb.Action_NewBullet) {
	h.mu.Lock()
	defer h.mu.Unlock()
	t, ok := h.game.TankBucket[nb.NewBullet.TankId]
	if !ok {
		return
	}
	if t.IsDead || t.IsLoading {
		return
	}
	id := len(h.game.BulletBucket) + 1
	b := &game.Bullet{
		Id:        int32(id),
		X:         nb.NewBullet.X,
		Y:         nb.NewBullet.Y,
		Direction: nb.NewBullet.Direction,
		TankId:    nb.NewBullet.TankId,
	}
	t.IsLoading = true
	go func() {
		time.Sleep(3 * time.Second)
		t.IsLoading = false
	}()
	//log.Println("new bullet", b)
	h.game.NewBullet(b)
}

func (h *Handler) UpdateStatus() []*pb.Action {
	actions := make([]*pb.Action, 0)
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, b := range h.game.BulletBucket {
		b.Move()
		if h.game.IsBulletHitTank(b) {
			h.game.RemoveBullet(b)
			h.game.NewExplosion(b.X, b.Y)
			count := 0
			for _, t := range h.game.TankBucket {
				if !t.IsDead {
					count++
				}
			}
			if count == 1 {
				h.status = consts.GameOver
				actions = append(actions, &pb.Action{
					Type: &pb.Action_GameOver{
						GameOver: &pb.GameOver{},
					},
				})
				return actions
			}

		} else if h.game.IsHitRock(b.X, b.Y) {
			h.game.RemoveBullet(b)
			h.game.NewExplosion(b.X, b.Y)
		} else if h.game.IsHitBorder(b.X, b.Y) {
			h.game.RemoveBullet(b)
			h.game.NewExplosion(b.X, b.Y)
		}
	}

	//wg := sync.WaitGroup{}
	//wg.Add(3)
	//var tl, bl, el *pb.Action
	//go func() {
	//	tl = h.UpdateTankList()
	//	wg.Done()
	//}()
	//go func() {
	//	bl = h.UpdateBulletList()
	//	wg.Done()
	//}()
	//go func() {
	//	el = h.UpdateExplosion()
	//	wg.Done()
	//}()
	//wg.Wait()

	// 三个 Update 都不用加读锁，因为写锁没有释放
	actions = append(actions, h.UpdateTankList(), h.UpdateBulletList(), h.UpdateExplosion())
	return actions
}

func (h *Handler) UpdateTankList() *pb.Action {

	tl := &pb.Action_GetTankList{
		GetTankList: &pb.GetTankList{},
	}
	for _, t := range h.game.TankBucket {
		tank := &pb.Tank{
			Id:        t.Id,
			X:         t.X,
			Y:         t.Y,
			Direction: t.Direction,
			Kill:      t.Kill,
			Color:     t.Color,
			Name:      t.Name,
			IsDead:    t.IsDead,
		}
		tl.GetTankList.Tank = append(tl.GetTankList.Tank, tank)
	}
	action := &pb.Action{
		Type: tl,
	}
	//log.Println("tank list", tl)
	return action
}

func (h *Handler) UpdateBulletList() *pb.Action {

	bl := &pb.Action_GetBulletList{
		GetBulletList: &pb.GetBulletList{},
	}
	for _, b := range h.game.BulletBucket {
		bullet := &pb.Bullet{
			Id: b.Id,
			X:  b.X,
			Y:  b.Y,
		}
		bl.GetBulletList.Bullet = append(bl.GetBulletList.Bullet, bullet)
	}
	action := &pb.Action{
		Type: bl,
	}
	return action
}

func (h *Handler) UpdateExplosion() *pb.Action {

	el := &pb.Action_GetExplosionList{
		GetExplosionList: &pb.GetExplosionList{},
	}
	for _, e := range h.game.ExplosionBucket {
		explosion := &pb.Explosion{
			X: e.X,
			Y: e.Y,
		}
		el.GetExplosionList.Explosion = append(el.GetExplosionList.Explosion, explosion)
	}
	action := &pb.Action{
		Type: el,
	}
	return action
}
