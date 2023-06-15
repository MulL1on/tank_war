// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"tank_war/server/cmd/api/config"
	"tank_war/server/shared/errno"
	"tank_war/server/shared/tools"

	"github.com/cloudwego/hertz/pkg/app"
	huser "tank_war/server/cmd/api/biz/model/user"
	kuser "tank_war/server/shared/kitex_gen/user"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.RegisterReq
	resp := new(kuser.RegisterResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if req.Username == "" || req.Password == "" {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := config.GlobalUserClient.Register(ctx, &kuser.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.BaseResp = tools.BuildBaseResp(errno.Success)
	c.JSON(http.StatusOK, res)
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req huser.LoginReq
	resp := new(kuser.LoginResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if req.Username == "" || req.Password == "" {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := config.GlobalUserClient.Login(ctx, &kuser.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, res)
}
