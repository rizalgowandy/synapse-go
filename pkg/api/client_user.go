package api

import (
	"context"

	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func (c *Client) ViewAllUsers(ctx context.Context, req *entity.ViewAllUsersReq) (*entity.ViewAllUsersResp, error) {
	url := "/v3.1/users?show_refresh_tokens={show_refresh_tokens}&per_page={per_page}&page={page}"

	var (
		content    entity.ViewAllUsersResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"show_refresh_tokens": req.ShowRefreshTokens,
			"per_page":            req.PerPage,
			"page":                req.Page,
		}).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader("", "")).
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

func (c *Client) ViewUser(ctx context.Context, req *entity.ViewUserReq) (*entity.ViewUserResp, error) {
	url := "/v3.1/users/{user_id}?full_dehydrate={full_dehydrate}"

	var (
		content    entity.ViewUserResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetPathParam("full_dehydrate", req.FullDehydrate).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader("", req.UserID)).
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

func (c *Client) CreateUser(ctx context.Context, req *entity.CreateUserReq) (*entity.CreateUserResp, error) {
	url := "/v3.1/users"

	var (
		content    entity.CreateUserResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader("", req.UserID)).
		SetHeader("Content-Type", "application/json").
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

func (c *Client) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) (*entity.UpdateUserResp, error) {
	url := "/v3.1/users/{user_id}"

	var (
		content    entity.UpdateUserResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetResult(&content).
		SetError(&contentErr).
		Patch(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) GenerateUBODoc(ctx context.Context, req *entity.GenerateUBODocReq) (*entity.GenerateUBODocResp, error) {
	url := "/v3.1/users/{user_id}/ubo"

	var (
		content    entity.GenerateUBODocResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetResult(&content).
		SetError(&contentErr).
		Patch(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) GetDuplicates(ctx context.Context, req *entity.GetDuplicatesReq) (*entity.GetDuplicatesResp, error) {
	url := "/v3.1/users/{user_id}/get-duplicates"

	var (
		content    entity.GetDuplicatesResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
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

func (c *Client) SwapDuplicateUsers(ctx context.Context, req *entity.SwapDuplicateUsersReq) (*entity.SwapDuplicateUsersResp, error) {
	url := "/v3.1/users/{user_id}/swap-duplicate-users"

	var (
		content    entity.SwapDuplicateUsersResp
		contentErr entity.ErrResp
	)

	_, err := c.client.R().
		SetContext(ctx).
		SetPathParam("user_id", req.UserID).
		SetHeaders(c.UserIPHeader(req.UserIPAddress)).
		SetHeaders(c.UserHeader(req.UserOAuthKey, req.UserID)).
		SetHeader("Content-Type", "application/json").
		SetResult(&content).
		SetError(&contentErr).
		Patch(url)
	if err != nil {
		return nil, err
	}

	if contentErr.Valid() {
		return nil, contentErr
	}

	return &content, nil
}

func (c *Client) AllowedDocumentTypes(ctx context.Context) (*entity.AllowedDocumentTypesResp, error) {
	url := "/v3.1/users/document-types"

	var (
		content    entity.AllowedDocumentTypesResp
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
