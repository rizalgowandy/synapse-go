package synapse

import (
	"github.com/rizalgowandy/synapse-go/pkg/api"
)

// NewClient creates a client to interact with Synapse API.
func NewClient(cfg api.Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{}, nil
}

// Client is the main client to interact with Synapse API.
type Client struct {
}
