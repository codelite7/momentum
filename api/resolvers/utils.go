package resolvers

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
)

func getSession(ctx context.Context) sessmodels.SessionContainer {
	return session.GetSessionFromRequestContext(ctx)
}

func getUserUuid(ctx context.Context) (uuid.UUID, error) {
	var userUuid uuid.UUID
	userId := getSession(ctx).GetUserID()
	if userId == "" {
		return userUuid, errors.New("unauthorized")
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return userUuid, errors.New("unauthorized")
	}

	return userUuid, nil
}
