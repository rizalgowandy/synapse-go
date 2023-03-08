package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func (c *Client) CreateOAuthKey(ctx context.Context, req *entity.CreateOAuthKeyReq) (*entity.CreateOAuthKeyResp, error) {
	url := "/v3.1/oauth/{user_id}"

	var (
		content    entity.CreateOAuthKeyResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader("", req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Post(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) GenerateRefreshToken(ctx context.Context, req *entity.GenerateRefreshTokenReq) (*entity.GenerateRefreshTokenResp, error) {
	url := "/v3.1/oauth/{user_id}/login"

	var (
		content    entity.GenerateRefreshTokenResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader("", req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&content).
		SetError(&contentErr).
		Post(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}
