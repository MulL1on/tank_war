package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/paseto"
	"strconv"
	"tank_war/server/cmd/user/pkg/mysql"
	"tank_war/server/shared/consts"
	"tank_war/server/shared/errno"
	"tank_war/server/shared/kitex_gen/base"
	user "tank_war/server/shared/kitex_gen/user"
	"tank_war/server/shared/tools"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	TokenGenerator
	MysqlManager
	IDGenerator
	EncryptManager
}

type IDGenerator interface {
	CreateUUID() int64
}

type EncryptManager interface {
	EncryptPassword(password string) string
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

type MysqlManager interface {
	CreateUser(user *mysql.User) error
	GetUserById(id int64) (*mysql.User, error)
	GetUserByUsername(username string) (*mysql.User, error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp = new(user.RegisterResp)
	// 判断用户名是否已经存在
	_, err = s.MysqlManager.GetUserByUsername(req.Username)
	if err == nil {
		resp.BaseResp = tools.BuildBaseResp(errno.RecordAlreadyExist)
		return resp, nil
	}

	// 创建用户
	err = s.MysqlManager.CreateUser(&mysql.User{
		ID:       s.IDGenerator.CreateUUID(),
		Username: req.Username,
		Password: s.EncryptManager.EncryptPassword(req.Password),
	})
	if err != nil {
		klog.Error("create user error", err.Error())
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}
	resp.BaseResp = tools.BuildBaseResp(errno.Success)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = new(user.LoginResp)
	u, err := s.MysqlManager.GetUserByUsername(req.Username)
	if err != nil {
		if err == errno.RecordNotFound {
			resp.BaseResp = tools.BuildBaseResp(errno.RecordNotFound)
			return resp, nil
		}
		klog.Error("get user error", err.Error())
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}
	if u.Password != s.EncryptManager.EncryptPassword(req.Password) {
		resp.BaseResp = tools.BuildBaseResp(errno.BadRequest)
		return resp, nil
	}

	now := time.Now()
	resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
		ID:        strconv.FormatInt(u.ID, 10),
		Issuer:    consts.Issuer,
		Audience:  consts.User,
		IssuedAt:  now,
		NotBefore: now,
		ExpiredAt: now.Add(time.Hour * 24 * 7),
	})

	if err != nil {
		klog.Error("create token error", err.Error())
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}

	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp = new(user.GetUserInfoResp)
	u, err := s.MysqlManager.GetUserById(req.UserId)
	if err != nil {
		klog.Error("get user error", err.Error())
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
		return resp, nil
	}
	resp.User = &base.User{
		UserID:   u.ID,
		Username: u.Username,
	}

	return resp, nil
}
