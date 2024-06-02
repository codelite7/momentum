package common

import (
	"context"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	user2 "github.com/codelite7/momentum/api/ent/user"
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

var userInfoContextKey = &contextKey{"userInfo"}

type contextKey struct {
	name string
}

type UserInfo struct {
	UserId         pulid.ID
	ActiveTenantId pulid.ID
}

func AddUserInfoToContext(c echo.Context, workosUserId string) error {
	user, err := EntClient.User.Query().Where(user2.WorkosUserID(workosUserId)).First(c.Request().Context())
	if err != nil {
		return err
	}
	activeTenant, err := user.QueryActiveTenant().First(c.Request().Context())
	if err != nil {
		return err
	}
	userInfo := UserInfo{UserId: user.ID, ActiveTenantId: activeTenant.ID}
	ctx := context.WithValue(c.Request().Context(), userInfoContextKey, userInfo)
	c.SetRequest(c.Request().WithContext(ctx))
	return nil
}

func GetUserIdFromContext(ctx context.Context) UserInfo {
	val := ctx.Value(userInfoContextKey)
	if val == nil {
		err := tracerr.New("user id not found in context")
		Logger.Error("user id not found in context", zap.Error(err))
		panic(err)
	}
	userId, ok := val.(UserInfo)
	if !ok {
		err := tracerr.New("user id found in context but it is not a pulid")
		Logger.Error("user id not found in context", zap.Error(err))
		panic(err)
	}
	return userId
}
