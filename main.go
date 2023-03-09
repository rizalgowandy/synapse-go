package synapse

import (
	"context"
	"errors"

	"github.com/rizalgowandy/synapse-go/pkg/api"
	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

// NewClient creates a client to interact with Synapse API.
func NewClient(cfg *api.Config) (*Client, error) {
	if cfg == nil {
		return nil, errors.New("config: cannot be nil")
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		caller: api.NewClient(cfg),
	}, nil
}

// Client is the main client to interact with Synapse API.
type Client struct {
	caller api.ClientItf
}

func (c *Client) CreateOAuthKey(ctx context.Context, req *entity.CreateOAuthKeyReq) (*entity.CreateOAuthKeyResp, error) {
	return c.caller.CreateOAuthKey(ctx, req)
}

func (c *Client) GenerateRefreshToken(ctx context.Context, req *entity.GenerateRefreshTokenReq) (*entity.GenerateRefreshTokenResp, error) {
	return c.caller.GenerateRefreshToken(ctx, req)
}

func (c *Client) AllowedEntityTypes(ctx context.Context) (*entity.AllowedEntityTypesResp, error) {
	return c.caller.AllowedEntityTypes(ctx)
}
