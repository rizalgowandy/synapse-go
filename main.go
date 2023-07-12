package synapse

import (
	"errors"

	"github.com/rizalgowandy/synapse-go/pkg/api"
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
		ClientItf: api.NewClient(cfg),
	}, nil
}

// Client is the main client to interact with Synapse API.
type Client struct {
	api.ClientItf
}
