package common

import (
	"context"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/agent"
)

func GetDefaultAgent(ctx context.Context, client *ent.Client) (*ent.Agent, error) {
	return client.Agent.Query().Where(agent.Provider("perplexity"), agent.Model("llama-3-sonar-large-32k-online")).First(ctx)
}
