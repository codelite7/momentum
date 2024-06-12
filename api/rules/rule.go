package rules

import (
	"context"
	"github.com/codelite7/momentum/api/ent/privacy"
)

func AllowIfAuthenticated() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		return privacy.Allow
	})
}
