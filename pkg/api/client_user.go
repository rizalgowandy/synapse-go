package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func (c *Client) AllowedEntityScopes(ctx context.Context) (*entity.AllowedEntityScopesResp, error) {
	url := "/v3.1/users/entity-scopes"

	var (
		content    entity.AllowedEntityScopesResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) AllowedEntityTypes(ctx context.Context) (*entity.AllowedEntityTypesResp, error) {
	url := "/v3.1/users/entity-types"

	var (
		content    entity.AllowedEntityTypesResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetResult(&content).
		SetError(&contentErr).
		Get(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}
