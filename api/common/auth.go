package common

import (
	"context"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	user2 "github.com/codelite7/momentum/api/ent/user"
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

var userIdContextKey = &contextKey{"userId"}

type contextKey struct {
	name string
}

func AddUserIdToContext(c echo.Context, workosUserId string) {
	user, err := EntClient.User.Query().Where(user2.WorkosUserID(workosUserId)).First(c.Request().Context())
	if err != nil {
		panic(err)
	}
	ctx := context.WithValue(c.Request().Context(), userIdContextKey, user.ID)
	c.SetRequest(c.Request().WithContext(ctx))
}

func GetUserIdFromContext(ctx context.Context) pulid.ID {
	val := ctx.Value(userIdContextKey)
	if val == nil {
		err := tracerr.New("user id not found in context")
		Logger.Error("user id not found in context", zap.Error(err))
		panic(err)
	}
	userId, ok := val.(pulid.ID)
	if !ok {
		err := tracerr.New("user id found in context but it is not a pulid")
		Logger.Error("user id not found in context", zap.Error(err))
		panic(err)
	}
	return userId
}
