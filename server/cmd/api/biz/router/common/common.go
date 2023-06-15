package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/paseto"
	"net/http"
	"tank_war/server/shared/consts"

	pt "aidanwoods.dev/go-paseto"
	"tank_war/server/shared/errno"
	"tank_war/server/shared/tools"
)

func PasetoAuth(audience string) app.HandlerFunc {
	pf, err := paseto.NewV4PublicParseFunc(paseto.DefaultPublicKey, []byte(paseto.DefaultImplicit), paseto.WithAudience(audience), paseto.WithNotBefore())
	if err != nil {
		hlog.Fatal(err)
	}
	sh := func(ctx context.Context, c *app.RequestContext, token *pt.Token) {
		uid, err := token.GetString("id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, tools.BuildBaseResp(errno.BadRequest.WithMessage("missing user id  in token")))
			c.Abort()
			return
		}
		c.Set(consts.UserID, uid)
	}

	eh := func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusUnauthorized, tools.BuildBaseResp(errno.BadRequest.WithMessage("invalid token")))
		c.Abort()
	}
	return paseto.New(paseto.WithParseFunc(pf), paseto.WithSuccessHandler(sh), paseto.WithErrorFunc(eh))
}
