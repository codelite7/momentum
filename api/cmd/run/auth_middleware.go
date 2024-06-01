package run

import (
	"context"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/config"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"github.com/ztrue/tracerr"
	"net/http"
	"strings"
	"time"
)

var keyset jwk.Set

func AuthMiddleware() echo.MiddlewareFunc {
	var err error
	keyset, err = NewJWKSet()
	if err != nil {
		panic(err)
	}
	return authMiddleware
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := getTokenFromHeader(c)
		if err != nil {
			return unauthenticated(c)
		}
		token, err := jwt.ParseString(tokenString, jwt.WithKeySet(keyset))
		if err != nil {
			return unauthenticated(c)
		}
		c.Set("token", token)
		return next(c)
	}
}

func getTokenFromHeader(c echo.Context) (string, error) {
	cookies := c.Cookies()
	for _, cookie := range cookies {
		common.Logger.Info(cookie.String())
	}
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", tracerr.Errorf("no authorization header")
	}
	split := strings.SplitN(authHeader, " ", 2)
	if len(split) != 2 {
		return "", tracerr.Errorf("invalid authorization header")
	}
	return split[1], nil
}

func unauthenticated(c echo.Context) error {
	return c.NoContent(http.StatusUnauthorized)
}

func NewJWKSet() (jwk.Set, error) {
	jwksUrl, err := usermanagement.GetJWKSURL(config.WorkosClientId)
	if err != nil {
		return nil, err
	}
	jwkCache := jwk.NewCache(context.Background())
	url := jwksUrl.String()
	// register a minimum refresh interval for this URL.
	// when not specified, defaults to Cache-Control and similar resp headers
	err = jwkCache.Register(url, jwk.WithMinRefreshInterval(10*time.Minute))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// fetch once on application startup
	_, err = jwkCache.Refresh(ctx, url)
	if err != nil {
		return nil, err
	}
	// create the cached key set
	return jwk.NewCachedSet(jwkCache, url), nil
}
