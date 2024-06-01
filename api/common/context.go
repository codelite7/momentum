package common

import (
	"context"
	"github.com/codelite7/momentum/api/ent/schema/pulid"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

func GetUserIdFromContext(ctx context.Context) pulid.ID {
	token := ctx.Value("token")
	if token != nil {
		Logger.Info("subject", zap.String("subject", token.(jwt.Token).Subject()))

	}
	return ctx.Value("userId").(pulid.ID)
}
